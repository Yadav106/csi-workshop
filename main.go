// main.go
package main

import "fmt"

func main() {
	newblockchain := NewBlockchain()
	newblockchain.AddBlock("first transaction")
	newblockchain.AddBlock("second transaction")

	for i, block := range newblockchain.Blocks {
		fmt.Println("----------------------------------------------------------")
		fmt.Println("Block id: ", i)
		fmt.Println("Timestamp: ", block.Timestamp+int64(i))
		fmt.Println("Hash of the block: ", block.MyBlockHash)
		fmt.Println("Previous block hash: ", block.PreviousBlockHash)
		fmt.Println("All the transactions: ", block.AllData)
		fmt.Println("----------------------------------------------------------")
	}
}
