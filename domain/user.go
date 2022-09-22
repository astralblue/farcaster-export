package domain

import (
	"encoding/json"
	"errors"
	"net/http"
)

type User struct {
	Address      string     `json:"address"`
	Username     string     `json:"username"`
	DirectoryURL string     `json:"directoryUrl"`
	Directory    *Directory `json:"directory"`
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

func (u *User) FetchActivity() ([]*Activity, error) {
	if u.Directory == nil || u.Directory.Body == nil || u.Directory.Body.AddressActivityUrl == "" {
		return nil, errors.New("no address activity url defined")
	}
	resp, err := http.Get(u.Directory.Body.AddressActivityUrl)
	if err != nil {
		return nil, err
	}
	var a []*Activity
	if err := json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return nil, err
	}
	return a, nil
}
