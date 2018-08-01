package BLC

import "fmt"

type Blockchain struct {
	Block []*Block //储存有序区块
}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	//创建创世区块
	genesisBlock := CreateGenesisBlock("Genesis Data....")
	fmt.Println()
	fmt.Println([]*Block{genesisBlock})
	fmt.Println()
	//返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}