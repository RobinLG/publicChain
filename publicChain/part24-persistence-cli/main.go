package main

import (
	"robin/publicChain/part24-persistence-cli/BLC"
)

func main() {
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	cli := &BLC.CLI{blockchain}

	cli.Run()

}
