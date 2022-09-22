package domain

import (
	"farcaster/farcaster"
)

type User struct {
	Address   string                   `json:"address"`
	Username  string                   `json:"username"`
	Profile   *farcaster.User          `json:"profile"`
	Followers []*farcaster.UserAddress `json:"followers"`
}
