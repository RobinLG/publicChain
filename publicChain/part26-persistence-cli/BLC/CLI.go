package BLC

import (
	"flag"
	"os"
	"log"
	"fmt"
)

type CLI struct {
	Blockchain *Blockchain
}

func printUsage() {

	fmt.Println("Usage:")
	fmt.Println("createblock -data data -- 交易数据")
	fmt.Println("addblock -data data -- 交易数据")
	fmt.Println("printchain -- 输出区块交易信息")

}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.Blockchain.AddBlockToBlockchain(data)
}

func (cli *CLI) printchain() {
	cli.Blockchain.Printchain()
}

func (cli *CLI) createBlock(data string) {
	CreateBlockchainWithGenesisBlock(data)
}


func (cli *CLI) Run() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblock", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","http://roblog.tech", "新增区块....")
	createBlockchainWithGenesis :=  createBlockchainCmd.String("data", "Genesis block data....", "创世区块交易数据....")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblock":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)

	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == ""{
			printUsage()
			os.Exit(1)
		}

		//fmt.Println(*flagAddBlockData)
		cli.addBlock(*flagAddBlockData)
	}

	if printChainCmd.Parsed() {
		//fmt.Println("输出所有区块数据....")
		cli.printchain()
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainWithGenesis == "" {
			fmt.Println("交易数据不能为空....")
			printUsage()
			os.Exit(1)
		}

		cli.createBlock(*createBlockchainWithGenesis)
	}
}