package farcaster

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_GetProfile(t *testing.T) {
	c, err := NewClient()
	assert.NoError(t, err)

	ctx := context.Background()
	p, err := c.GetProfile(ctx, "0x58975eA416fAA998E1A88921B1e889aa7ecD9a41")
	assert.NoError(t, err)

	assert.Equal(t, "slanders", p.Username)
}

func TestClient_GetFollowers(t *testing.T) {
	c, err := NewClient()
	assert.NoError(t, err)

	ctx := context.Background()
	f, err := c.GetFollowers(ctx, "0x58975eA416fAA998E1A88921B1e889aa7ecD9a41")
	assert.NoError(t, err)

	assert.Equal(t, 10, len(f))
}
