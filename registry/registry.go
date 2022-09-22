package registry

import (
	"bytes"
	"context"
	"encoding/json"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"

	"farcaster/client"
	"farcaster/contracts/farcaster_registry"
	"farcaster/domain"
)

const (
	defaultRegAddress = "0xe3Be01D99bAa8dB9905b33a3cA391238234B79D1"
)

type Registry interface {
	ForEachUser(ctx context.Context, handler func(u *domain.User) error) error
}

type registry struct {
	c  *ethclient.Client
	fr *farcaster_registry.FarcasterRegistryCaller
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
	uc, err := r.fr.UsernamesLength(nil)
	if err != nil {
		log.WithError(err).Fatal("error getting length")
	}

	grp, ctx := errgroup.WithContext(ctx)
	idx := make(chan uint32)
	for i := 0; i < 20; i++ {
		grp.Go(func() error {
			for i := range idx {
				u, err := r.fr.UsernameAtIndex(nil, uint32(i))
				if err != nil {
					return err
				}
				username := string(bytes.Trim(u[:], "\x00"))
				url, err := r.fr.UsernameToUrl(nil, u)

				if err != nil {
					return err
				}
				d, err := getDirectory(url.Url)
				if err != nil || d.Body == nil {
					continue
				}

				parts := strings.Split(d.Body.AddressActivityUrl, "/")
				addr := parts[len(parts)-1]
				if !strings.HasPrefix(addr, "0x") {
					continue
				}
				if err := handler(&domain.User{
					Address:      parts[len(parts)-1],
					Username:     username,
					DirectoryURL: url.Url,
					Directory:    d,
				}); err != nil {
					return err
				}
			}
			return nil
		})
	}

	for i := int64(0); i < uc.Int64(); i++ {
		idx <- uint32(i)
	}

	close(idx)
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
	c, err := client.NewClient(cfg.JsonRpcUrl)
	if err != nil {
		return nil, err
	}

	reg, err := farcaster_registry.NewFarcasterRegistryCaller(common.HexToAddress(regAddr), c)
	if err != nil {
		return nil, err
	}
	return &registry{
		c:  c,
		fr: reg,
	}, nil
}
