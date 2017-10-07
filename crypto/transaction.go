package crypto

import "encoding/json"

type Transaction struct {
	//Index    uint64 `json:"index"`
	Owner      string `json:"owner"`
	Id         string `json:"id"`
	Previous   string `json:"previous"`
	Created    int64  `json:"created"`
	Transfered int64  `json:"created"`
}
type TransactionList struct {
	Items []Transaction `json:"items"`
}

func NewTransaction(index int) Transaction {
	t := Transaction{}
	//t.Index = uint64(index)
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
