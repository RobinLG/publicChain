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
	fmt.Println("\naddblock -data data -- 交易数据")
	fmt.Println("\nprintchain -- 交易数据")

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


func (cli *CLI) Run() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data","http://roblog.tech", "新增区块....")

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
}