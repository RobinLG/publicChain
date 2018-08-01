package main

import (
	"robin/publicChain/part7-proof-of-work/BLC"
	"fmt"
)

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	//fmt.Println(blockchain.Blocks[0])

	blockchain.AddBlockToBlockchain(blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1, blockchain.Blocks[len(blockchain.Blocks) - 1].Hash, "a send 100RMB to b")

	blockchain.AddBlockToBlockchain(blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1, blockchain.Blocks[len(blockchain.Blocks) - 1].Hash, "b send 100RMB to c")

	blockchain.AddBlockToBlockchain(blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1, blockchain.Blocks[len(blockchain.Blocks) - 1].Hash, "c send 100RMB to d")

	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)
}
