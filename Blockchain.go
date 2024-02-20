// Blockchain.go
package main

func (blockchain *Blockchain) AddBlock(data string) {
	Previousblock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := NewBlock(data, Previousblock.MyBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
