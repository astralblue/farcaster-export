package main

import (
	"context"
	"farcaster/domain"
	"farcaster/registry"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
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
	grp, ctx := errgroup.WithContext(ctx)
	nodes := make(chan *domain.User, 100)

	nodeFile, err := os.OpenFile("nodes.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.WithError(err).Fatal("error opening file")
	}
	linksFile, err := os.OpenFile("links.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.WithError(err).Fatal("error opening file")
	}

	grp.Go(func() error {
		for n := range nodes {
			_, err := nodeFile.WriteString(fmt.Sprintf("%s,\"%s\",User\n", n.Address, n.Profile.Username))
			if err != nil {
				return err
			}
			for _, f := range n.Followers {
				_, err := linksFile.WriteString(fmt.Sprintf("%s,%s,FOLLOWS\n", f.Address, n.Address))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	go func() {
		if err := grp.Wait(); err != nil {
			log.WithError(err).Fatal("errgroup error")
		}
	}()

	err = r.ForEachUser(ctx, func(u *domain.User) error {
		if err != nil {
			return err
		}
		log.WithFields(log.Fields{
			"username":  u.Username,
			"address":   u.Address,
			"followers": len(u.Followers),
		}).Info("user")
		nodes <- u
		return nil
	})
	close(nodes)

	if err != nil {
		log.WithError(err).Fatal("error creating registry")
	}
}
