package domain

type body struct {
	AddressActivityUrl string `json:"addressActivityUrl"`
	AvatarUrl          string `json:"avatarUrl"`
	DisplayName        string `json:"displayName"`
	ProofUrl           string `json:"proofUrl"`
	Timestamp          int64  `json:"timestamp"`
	Version            int    `json:"version"`
}

type Directory struct {
	Body       *body  `json:"body"`
	MerkleRoot string `json:"merkleRoot"`
	Signature  string `json:"signature"`
}
