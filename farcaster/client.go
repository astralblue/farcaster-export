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
const pathFollowers = "/indexer/followers"

var minBackoff = 5 * time.Second
var maxBackoff = 30 * time.Second

type Client interface {
	GetProfile(ctx context.Context, address string) (*User, error)
	GetFollowers(ctx context.Context, address string) ([]*UserAddress, error)
}

type client struct {
}

var ErrRateLimit = errors.New("rate limited")

type RetryOptions struct {
	MaxElapsedTime time.Duration
	MinBackoff     time.Duration
	MaxBackoff     time.Duration
}

func withBackoff(
	ctx context.Context, operation func(ctx context.Context) error, options RetryOptions,
) error {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = options.MaxElapsedTime
	bo.InitialInterval = options.MinBackoff
	bo.MaxInterval = options.MaxBackoff

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
		req = req.WithContext(ctx)

		dc := http.DefaultClient
		resp, err := dc.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 || resp.StatusCode == 502 {
			return ErrRateLimit
		}
		if resp.StatusCode >= 400 {
			return fmt.Errorf("error code %d, path=%s", resp.StatusCode, path)
		}
		return json.NewDecoder(resp.Body).Decode(target)
	}, RetryOptions{
		MinBackoff: minBackoff,
		MaxBackoff: maxBackoff,
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

func (c *client) GetFollowers(ctx context.Context, address string) (result []*UserAddress, err error) {
	var n Followers
	err = c.get(ctx, fmt.Sprintf("%s%s/%s", host, pathFollowers, address), &n)
	if err != nil {
		return nil, err
	}
	for _, f := range n {
		result = append(result, &UserAddress{Username: f.Username, Address: f.Address})
	}
	return result, nil
}

func NewClient() (Client, error) {
	return &client{}, nil
}
