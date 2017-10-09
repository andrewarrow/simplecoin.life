package crypto

import "encoding/json"

type Transaction struct {
	//Index    uint64 `json:"index"`
	Owner      string `json:"owner"`
	Id         string `json:"id"`
	Previous   string `json:"previous"`
	Created    int64  `json:"created"`
	Transfered int64  `json:"transfered"`
}
type TransactionList struct {
	Items []Transaction `json:"items"`
}

func NewTransaction(id, owner, previous string, created, transfered int64) Transaction {
	t := Transaction{}
	t.Id = id
	t.Owner = owner
	t.Previous = previous
	t.Created = created
	t.Transfered = transfered
	return t
}
func (t Transaction) Encode() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
func (t TransactionList) Encode() string {
	b, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
