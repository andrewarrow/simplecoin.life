package crypto

import "encoding/json"

type Transaction struct {
	Hash                    string   `json:"hash"`
	OwnersPublicAccount     []uint64 `json:"owners_public_account"`
	PreviousOwnersSignature []uint64 `json:"previous_owners_signature"`
}

func NewTransaction() Transaction {
	t := Transaction{}
	return t
}
func (t Transaction) Encode() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
