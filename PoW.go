package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

var MaxNonce = math.MaxInt64

const Difficulty = 24

type PoW struct {
	block  *Block
	target *big.Int
}

func NewPoW(b *Block) *PoW {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow := &PoW{
		block:  b,
		target: target,
	}

	return pow
}

func (pow *PoW) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevHash,
			pow.block.Data,
			IntToHex(pow.block.CreatedAt),
			IntToHex(int64(Difficulty)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *PoW) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < MaxNonce {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

func (pow *PoW) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(int(pow.block.Nonce))
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
