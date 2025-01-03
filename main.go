package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	chain []*Block
}

type Block struct {
	prev_hash []byte
	hash      []byte
	data      []byte
}

func (b *Block) derive_hash() {
	digest := bytes.Join([][]byte{b.data, b.prev_hash}, []byte{})
	hash := sha256.Sum256(digest)
	b.hash = hash[:]
}

func create_block(_prev_hash []byte, _data []byte) *Block {
	block := &Block{_prev_hash, []byte{}, _data}
	block.derive_hash()
	return block
}

func (c *Blockchain) add_block(_data string) {
	_prev_hash := c.chain[len(c.chain)-1].hash
	b := create_block(_prev_hash, []byte(_data))
	c.chain = append(c.chain, b)
}

func generate_genesis() *Block {
	genesis_data := []byte("GENESIS")
	b := &Block{[]byte{}, []byte{}, genesis_data}
	b.derive_hash()
	return b
}

func init_chain() *Blockchain {
	return &Blockchain{
		[]*Block{generate_genesis()},
	}
}

func main() {
	chain := init_chain()
	chain.add_block("First Block")
	chain.add_block("Second Block")
	chain.add_block("Third Block")

	for _, block := range chain.chain {
		fmt.Printf("Previous Hash: %x\n", block.prev_hash)
		fmt.Printf("Current Hash: %x\n", block.hash)
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("\n")
	}
}
