package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/NoTabaco/nomadcoin/db"
	"github.com/NoTabaco/nomadcoin/utils"
)

type blockchain struct {
	Height     int    `json:"height"`
	NewestHash string `json:"newestHash"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	utils.HandleErr(decoder.Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{0, ""}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				fmt.Println("Restoring...")
				b.restore(checkpoint)
			}
		})
	}
	return b
}
