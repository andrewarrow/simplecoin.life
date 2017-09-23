package crypto

import "encoding/json"

type Transaction struct {
	Index                   uint64   `json:"index"`
	Hash                    string   `json:"hash"`
	OwnersPublicAccount     []uint64 `json:"owners_public_account"`
	PreviousOwnersSignature []uint64 `json:"previous_owners_signature"`
}

func NewTransaction(index int) Transaction {
	t := Transaction{}
	t.Index = uint64(index)
	return t
}
func (t Transaction) Encode() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
