package main

// Blockchain describes the model of a new blockchain
type Blockchain struct {
	blocks []*Block
}

// NewBlockChain inits new example of a blockchain
func (b *Blockchain) NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// AddBlock adds new block to a blockchain
func (b *Blockchain) AddBlock(data string) {
	prevBlock := b.blocks[len(b.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	b.blocks = append(b.blocks, newBlock)
}

// NewGenesisBlock creates first block in a current blockchain
func NewGenesisBlock() *Block {
	return NewBlock("Genesis block", []byte{})
}
