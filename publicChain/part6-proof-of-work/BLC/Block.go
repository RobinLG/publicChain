package BLC

import (
	"time"
)

type Block struct {
	//1. 区块高度
	Height int64
	//2. 上一个区块的Hash值
	PrevBlockHash []byte
	//3. 交易数据
	Data []byte
	//4. 时间戳
	Timestamp int64
	//5. Hash
	Hash []byte
	//6. Nonce
	Nonce int64
}

//1. 工厂方法：创建新区块
func NewBlock(height int64, prevBlockHash []byte, data string) *Block{

	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	//创建工作量证明对象
	pow := NewProofOfWork(block)

	hash, nonce := pow.run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

//2. 单独写创建创世区块的方法
func CreateGenesisBlock(data string) *Block {
	return NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, data)
}

