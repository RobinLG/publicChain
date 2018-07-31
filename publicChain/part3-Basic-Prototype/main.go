package main

import (
	"robin/publicChain/part3-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()

	fmt.Println(genesisBlockchain)
	fmt.Println(genesisBlockchain.Block)
	fmt.Println(genesisBlockchain.Block[0])

}
