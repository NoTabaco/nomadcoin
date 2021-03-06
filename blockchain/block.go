package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/NoTabaco/nomadcoin/db"
	"github.com/NoTabaco/nomadcoin/utils"
)

type Block struct {
	Height       int    `json:"height"`
	Hash         string `json:"hash"`
	PrevHash     string `json:"prevHash,omitempty"`
	Difficulty   int    `json:"difficulty"`
	Nonce        int    `json:"nonce"`
	Timestamp    int    `json:"timestamp"`
	Transactions []*Tx  `json:"transactions"`
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

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(prevHash string, height int) *Block {
	block := &Block{
		Height:     height + 1,
		Hash:       "",
		PrevHash:   prevHash,
		Difficulty: Blockchain().difficulty(),
		Nonce:      0,
	}
	block.Transactions = Mempool.TxToConfirm()
	block.mine()
	block.persist()
	return block
}
