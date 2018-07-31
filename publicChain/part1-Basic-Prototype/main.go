package main

import (
	"fmt"
	"robin/publicChain/part1-Basic-Prototype/BLC"
)

func main() {
	// Hash值64位等于32个字节
	block := BLC.NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, "GenesisBlock")

	fmt.Println(block)

}
