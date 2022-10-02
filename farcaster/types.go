package farcaster

// these structs are generated from example response json (for expediency)

type notificationsPage struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
	User      struct {
		Username    string `json:"username"`
		Address     string `json:"address"`
		DisplayName string `json:"displayName"`
		Avatar      struct {
			Url        string `json:"url"`
			IsVerified bool   `json:"isVerified"`
		} `json:"avatar"`
		IsViewerFollowing bool `json:"isViewerFollowing"`
		IsFollowingViewer bool `json:"isFollowingViewer"`
	} `json:"user"`
	AggregatedItems []struct {
		Id        string `json:"id"`
		Timestamp int64  `json:"timestamp"`
		User      struct {
			Username    string `json:"username"`
			Address     string `json:"address"`
			DisplayName string `json:"displayName"`
			Avatar      struct {
				Url        string `json:"url"`
				IsVerified bool   `json:"isVerified"`
			} `json:"avatar"`
			IsViewerFollowing bool `json:"isViewerFollowing"`
			IsFollowingViewer bool `json:"isFollowingViewer"`
		} `json:"user"`
	} `json:"aggregatedItems"`
}

type notificationsResult struct {
	Notifications map[string]notificationsPage
}

// Notifications obj is generated from json: https://guardian.farcaster.xyz/v1/notifications?address={address}
type Notifications struct {
	Result notificationsResult `json:"result"`
	Meta   struct {
		Next string `json:"next"`
	} `json:"meta"`
}

type UserAddress struct {
	Username string `json:"username"`
	Address  string `json:"address"`
}

func (n *Notifications) ExtractFollowers() []*UserAddress {
	var result []*UserAddress
	uniq := make(map[string]bool)
	for _, entry := range n.Result.Notifications {
		for _, n := range entry.AggregatedItems {
			if !n.User.IsFollowingViewer {
				continue
			}
			if _, ok := uniq[n.User.Address]; !ok {
				result = append(result, &UserAddress{
					Username: n.User.Username,
					Address:  n.User.Address,
				})
				uniq[n.User.Address] = true
			}
		}
	}
	return result
}

type User struct {
	Address           string `json:"address"`
	Username          string `json:"username"`
	DisplayName       string `json:"displayName"`
	Avatar            Avatar `json:"avatar"`
	FollowerCount     int    `json:"followerCount"`
	FollowingCount    int    `json:"followingCount"`
	IsViewerFollowing bool   `json:"isViewerFollowing"`
	IsFollowingViewer bool   `json:"isFollowingViewer"`
	Profile           struct {
		Bio struct {
			Text     string        `json:"text"`
			Mentions []interface{} `json:"mentions"`
		} `json:"bio"`
	} `json:"profile"`
}

// Avatar is found in User and Profile objects
type Avatar struct {
	Url        string `json:"url"`
	IsVerified bool   `json:"isVerified"`
}

// Profile is generated from json: https://guardian.farcaster.xyz/v1/profiles/{address}
type Profile struct {
	Result struct {
		User *User `json:"user"`
	} `json:"result"`
}

// Activity generated from activity struct
type Activity struct {
	Body struct {
		Type        string `json:"type"`
		PublishedAt int64  `json:"publishedAt"`
		Sequence    int    `json:"sequence"`
		Username    string `json:"username"`
		Address     string `json:"address"`
		Data        struct {
			Text                  string `json:"text"`
			ReplyParentMerkleRoot string `json:"replyParentMerkleRoot"`
		} `json:"data"`
		PrevMerkleRoot string `json:"prevMerkleRoot"`
	} `json:"body"`
	MerkleRoot string `json:"merkleRoot"`
	Signature  string `json:"signature"`
	Meta       struct {
		DisplayName      string `json:"displayName"`
		Avatar           string `json:"avatar"`
		IsVerifiedAvatar bool   `json:"isVerifiedAvatar"`
		NumReplyChildren int    `json:"numReplyChildren"`
		Reactions        struct {
			Count int    `json:"count"`
			Type  string `json:"type"`
			Self  bool   `json:"self"`
		} `json:"reactions"`
		Recasts struct {
			Count int  `json:"count"`
			Self  bool `json:"self"`
		} `json:"recasts"`
		Watches struct {
			Count int  `json:"count"`
			Self  bool `json:"self"`
		} `json:"watches"`
		Mentions []struct {
			Address  string `json:"address"`
			Username string `json:"username"`
		} `json:"mentions"`
		ReplyParentUsername struct {
			Address  string `json:"address"`
			Username string `json:"username"`
		} `json:"replyParentUsername"`
	} `json:"meta"`
}

// Followers are generated from json: https://guardian.farcaster.xyz/indexer/followers/{address}
type Followers []Follower

// Follower is one follower from the followers API
type Follower struct {
	// TODO(ek): define/use real type instead of bool for Verifications
	Address           string `json:"address"`
	Username          string `json:"username"`
	DisplayName       string `json:"displayName"`
	Avatar            Avatar `json:"avatar"`
	IsViewerFollowing bool   `json:"isViewerFollowing"'`
	Verifications     []bool `json:"verifications"`
}
