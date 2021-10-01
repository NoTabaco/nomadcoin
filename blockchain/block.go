package blockchain

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Height   int    `json:"height"`
	Data     string `json:"data"`
	Hash     string `json:"hash"`
	PrevHash string `json:"prevHash,omitempty"`
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Height:   height + 1,
		Data:     data,
		PrevHash: prevHash,
		Hash:     "",
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	return block
}
