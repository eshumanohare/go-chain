package main

import (
	"fmt"

	"github.com/eshumanohare/go-chain/gochain"
)

func main() {
	blockchain := gochain.Init_chain()
	blockchain.Add_block("First Block")
	blockchain.Add_block("Second Block")
	blockchain.Add_block("Third Block")

	for _, block := range blockchain.Chain {
		fmt.Printf("Block Number: %d\n", block.Block_number)
		fmt.Printf("Previous Hash: %x\n", block.Prev_hash)
		fmt.Printf("Current Hash: %x\n", block.Hash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("\n")
	}
}
