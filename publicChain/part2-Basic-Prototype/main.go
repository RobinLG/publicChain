package main

import (
	"fmt"
	"robin/publicChain/part2-Basic-Prototype/BLC"
)

func main() {
	// Hash值64位等于32个字节
	genesisBlock := BLC.CreateGenesisBlock("GenesisBlock")

	fmt.Println(genesisBlock)

	}
