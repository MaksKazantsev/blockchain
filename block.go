package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block describes a model of the block in blockchain
type Block struct {
	CreatedAt int64
	Hash      []byte
	PrevHash  []byte
	Data      []byte
	Nonce     int64
}

// NewBlock created a block which will be used in blockchain
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{CreatedAt: time.Now().Unix(), Hash: []byte{}, PrevHash: prevBlockHash, Data: []byte(data), Nonce: 0}
	block.SetHash()
	return block
}

// SetHash sets hash of a new block
func (b *Block) SetHash() {
	timeCreated := []byte(strconv.FormatInt(b.CreatedAt, 10))
	data := bytes.Join([][]byte{b.Data, b.PrevHash, timeCreated}, []byte{})
	hash := sha256.Sum256(data)

	b.Hash = hash[:]
}
