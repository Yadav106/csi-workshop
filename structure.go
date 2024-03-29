// structure.go
package main

type Block struct {
	Timestamp         int64
	PreviousBlockHash []byte
	MyBlockHash       []byte
	AllData           []byte
}

type Blockchain struct {
	Blocks []*Block
}
