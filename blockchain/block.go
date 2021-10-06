package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/NoTabaco/nomadcoin/db"
	"github.com/NoTabaco/nomadcoin/utils"
)

const difficulty int = 2

type Block struct {
	Height     int    `json:"height"`
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

var ErrNotFound = errors.New("Block not found")

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
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
	block.persist()
	return block
}
