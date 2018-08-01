package main

import (
	"robin/publicChain/part8-proof-of-work/BLC"
	"fmt"
)

func main() {
	//blockchain := BLC.CreateBlockchainWithGenesisBlock()
	//
	////fmt.Println(blockchain.Blocks[0])
	//
	//blockchain.AddBlockToBlockchain(blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1, blockchain.Blocks[len(blockchain.Blocks) - 1].Hash, "a send 100RMB to b")
	//
	//blockchain.AddBlockToBlockchain(blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1, blockchain.Blocks[len(blockchain.Blocks) - 1].Hash, "b send 100RMB to c")
	//
	//blockchain.AddBlockToBlockchain(blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1, blockchain.Blocks[len(blockchain.Blocks) - 1].Hash, "c send 100RMB to d")
	//
	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)

	// height int64, prevBlockHash []byte, data string
	block := BLC.NewBlock(1,[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, "Test")
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	proofOfBlock := BLC.NewProofOfWork(block)
	fmt.Printf("%v", proofOfBlock.IsValue())
}
