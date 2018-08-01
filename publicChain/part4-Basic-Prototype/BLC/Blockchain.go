package BLC

import "fmt"

type Blockchain struct {
	Blocks []*Block //储存有序区块
}

// 增加区块到区块链里
func (blc *Blockchain) AddBlockToBlockchain(hight int64, preHash []byte, data string) {
	// 创建新区块
	newBlock := NewBlock(hight, preHash, data)
	// 往区块链里添加区块
	blc.Blocks = append(blc.Blocks, newBlock)
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
