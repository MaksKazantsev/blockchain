package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	CreatedAt int64
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{Hash: []byte{}, PrevHash: prevBlockHash, Data: []byte(data), CreatedAt: time.Now().Unix()}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	timeCreated := []byte(strconv.FormatInt(b.CreatedAt, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timeCreated}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}
