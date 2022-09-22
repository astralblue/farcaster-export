package main

import (
	"context"
	"farcaster/domain"
	"farcaster/registry"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	url := os.Getenv("RINKEBY_JSON_RPC")
	if url == "" {
		log.Fatal("set RINKEBY_JSON_RPC to a valid rinkeby json rpc url")
	}
	r, err := registry.NewRegistry(registry.Config{
		JsonRpcUrl: url,
	})
	if err != nil {
		log.WithError(err).Fatal("error creating registry")
	}
	ctx := context.Background()
	err = r.ForEachUser(ctx, func(u *domain.User) error {
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"username":  u.Username,
			"address":   u.Address,
			"followers": len(u.Followers),
		}).Info("user")
		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("error creating registry")
	}
}
