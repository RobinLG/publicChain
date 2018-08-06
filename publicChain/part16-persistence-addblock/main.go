package main

import (
	"robin/publicChain/part16-persistence-addblock/BLC"
)

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()


	blockchain.AddBlockToBlockchain("a send 100RMB to b")

	blockchain.AddBlockToBlockchain("b send 100RMB to c")

	blockchain.AddBlockToBlockchain("c send 100RMB to d")

	//blockchain.Printchain()

}
