package registry

import (
	"bytes"
	"context"
	"encoding/json"
	"farcaster/farcaster"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"

	"farcaster/client"
	"farcaster/contracts/farcaster_registry"
	"farcaster/domain"
)

const (
	defaultRegAddress = "0xe3Be01D99bAa8dB9905b33a3cA391238234B79D1"
	registerNameTopic = "0x00af56d6ef7b1de1b37e4636c3e255d78dc3e9359036fb04ca51e32d946a6834"
	genesisBlock      = 9145238
)

type Registry interface {
	ForEachUser(ctx context.Context, handler func(u *domain.User) error) error
}

type registry struct {
	fc farcaster.Client
	c  *ethclient.Client
	fr *farcaster_registry.FarcasterRegistryCaller
	ff *farcaster_registry.FarcasterRegistryFilterer
}

func getDirectory(url string) (*domain.Directory, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var d domain.Directory
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}
	return &d, nil
}

func (r *registry) ForEachUser(ctx context.Context, handler func(u *domain.User) error) error {
	grp, ctx := errgroup.WithContext(ctx)
	users := make(chan *farcaster_registry.FarcasterRegistryRegisterName)
	for i := 0; i < 1; i++ {
		grp.Go(func() error {
			for u := range users {
				username := string(bytes.Trim(u.Username[:], "\x00"))
				logger := log.WithFields(log.Fields{
					"user":    username,
					"address": u.Owner.String(),
				})
				profile, err := r.fc.GetProfile(ctx, u.Owner.String())
				if err != nil {
					logger.WithError(err).Error("error getting user")
					continue
				}
				followers, err := r.fc.GetFollowers(ctx, u.Owner.String())
				if err != nil {
					logger.WithError(err).Error("error getting followers")
					continue
				}

				if err := handler(&domain.User{
					Address:   u.Owner.String(),
					Username:  username,
					Profile:   profile,
					Followers: followers,
				}); err != nil {
					return err
				}
			}
			return nil
		})
	}

	bn, err := r.c.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	evts, err := r.ff.FilterRegisterName(&bind.FilterOpts{
		Start:   uint64(genesisBlock),
		End:     &bn,
		Context: ctx,
	}, nil, nil)
	if err != nil {
		return err
	}

	for evts.Next() {
		users <- evts.Event
	}

	close(users)
	return grp.Wait()
}

type Config struct {
	JsonRpcUrl      string
	RegistryAddress string
}

func NewRegistry(cfg Config) (*registry, error) {
	regAddr := defaultRegAddress
	if cfg.RegistryAddress != "" {
		regAddr = cfg.RegistryAddress
	}
	c, err := client.NewEthClient(cfg.JsonRpcUrl)
	if err != nil {
		return nil, err
	}

	reg, err := farcaster_registry.NewFarcasterRegistryCaller(common.HexToAddress(regAddr), c)
	if err != nil {
		return nil, err
	}
	ff, err := farcaster_registry.NewFarcasterRegistryFilterer(common.HexToAddress(regAddr), c)
	if err != nil {
		return nil, err
	}
	fc, err := farcaster.NewClient()
	if err != nil {
		return nil, err
	}
	return &registry{
		fc: fc,
		c:  c,
		fr: reg,
		ff: ff,
	}, nil
}
