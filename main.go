package main

import (
	"context"
	"farcaster/domain"
	"farcaster/registry"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	url := os.Getenv("ALCHEMY_RINKEBY")
	r, err := registry.NewRegistry(registry.Config{
		JsonRpcUrl: url,
	})
	if err != nil {
		log.WithError(err).Fatal("error creating registry")
	}
	ctx := context.Background()
	err = r.ForEachUser(ctx, func(u *domain.User) error {
		act, err := u.FetchActivity()
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"username": u.Username,
			"address":  u.Address,
			"casts":    len(act),
		}).Info("user")
		return nil
	})
	if err != nil {
		log.WithError(err).Fatal("error creating registry")
	}
}
