package farcaster

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	backoff "github.com/cenkalti/backoff/v4"
)

const host = "https://guardian.farcaster.xyz"
const pathProfile = "/v1/profiles"
const pathNotifications = "/v1/notifications"

var minBackoff = 5 * time.Second
var maxBackoff = 1 * time.Minute

type Client interface {
	GetProfile(ctx context.Context, address string) (*User, error)
	GetFollowers(ctx context.Context, address string) ([]*UserAddress, error)
}

type client struct {
}

var ErrRateLimit = errors.New("rate limited")

type RetryOptions struct {
	MaxElapsedTime *time.Duration
	MinBackoff     *time.Duration
	MaxBackoff     *time.Duration
}

func withBackoff(
	ctx context.Context, operation func(ctx context.Context) error, options RetryOptions,
) error {
	bo := backoff.NewExponentialBackOff()
	err := backoff.Retry(func() error {
		if ctx.Err() != nil {
			return backoff.Permanent(ctx.Err())
		}

		tCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
		err := operation(tCtx)
		cancel()

		if err == nil {
			return nil
		} else if ctx.Err() != nil {
			return backoff.Permanent(ctx.Err())
		} else if err == ErrRateLimit {
			log.Warnf("%s (retryable)", err.Error())
			return err
		}
		return backoff.Permanent(err)
	}, bo)
	if err != nil {
		log.Errorf("failed with error: %s", err.Error())
	}
	return err
}

func (c *client) get(ctx context.Context, path string, target interface{}) error {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}

	return withBackoff(ctx, func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		req = req.WithContext(ctx)

		dc := http.DefaultClient
		resp, err := dc.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 {
			return ErrRateLimit
		}
		if resp.StatusCode >= 400 {
			return fmt.Errorf("error code %d, path=%s", resp.StatusCode, path)
		}
		return json.NewDecoder(resp.Body).Decode(target)
	}, RetryOptions{
		MinBackoff: &minBackoff,
		MaxBackoff: &maxBackoff,
	})
}

func (c *client) GetProfile(ctx context.Context, address string) (*User, error) {
	var p Profile
	err := c.get(ctx, fmt.Sprintf("%s%s/%s", host, pathProfile, address), &p)
	if err != nil {
		return nil, err
	}
	return p.Result.User, nil
}

func (c *client) GetFollowers(ctx context.Context, address string) ([]*UserAddress, error) {
	//TODO implement me
	followers := make(map[string]bool)
	var n Notifications
	err := c.get(ctx, fmt.Sprintf("%s%s?address=%s", host, pathNotifications, address), &n)
	if err != nil {
		return nil, err
	}
	result := n.ExtractFollowers()
	for _, f := range result {
		followers[f.Address] = true
	}
	nextCursor := n.Meta.Next
	for nextCursor != "" {
		var n Notifications
		err := c.get(ctx, nextCursor, &n)
		if err != nil {
			return nil, err
		}
		if len(n.Result.Notifications) == 0 {
			break
		}
		for _, f := range n.ExtractFollowers() {
			if _, ok := followers[f.Address]; !ok {
				followers[f.Address] = true
				result = append(result, f)
			}
		}
		nextCursor = n.Meta.Next
	}

	return result, nil
}

func NewClient() (Client, error) {
	return &client{}, nil
}
