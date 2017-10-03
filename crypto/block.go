package crypto

import "encoding/json"

type Block struct {
	Index        uint64 `json:"index"`
	Timestamp    uint64 `json:"timestamp"`
	Data         string `json:"data"`
	Hash         string `json:"hash"`
	PreviousHash string `json:"previous_hash"`
}

func NewBlock(index int) Block {
	b := Block{}
	b.Index = uint64(index)
	return b
}
func (b Block) Encode() string {
	data, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
