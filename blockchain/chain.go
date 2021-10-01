package blockchain

import (
	"sync"
)

type blockchain struct {
	Height     int    `json:"height"`
	NewestHash string `json:"newestHash"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{0, ""}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}
