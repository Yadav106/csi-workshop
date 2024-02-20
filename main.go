// main.go
package main

import "fmt"

func main() {
	newblockchain := NewBlockchain()
	newblockchain.AddBlock("first transaction")
	newblockchain.AddBlock("second transaction")

	for i, block := range newblockchain.Blocks {
		fmt.Println("Block id:%d\n", i)
		fmt.Println("Timestamp: %d\n", block.Timestamp+int64(i))
		fmt.Println("Hash of the block: % \n", block.MyBlockHash)
		fmt.Println("Previous block hash: %x \n", block.PreviousBlockHash)
		fmt.Println("All the transactions: %× \n", block.AllData)
	}
}
