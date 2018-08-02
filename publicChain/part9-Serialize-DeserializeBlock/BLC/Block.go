package BLC

import (
	"time"
	"fmt"
	"bytes"
	"encoding/gob"
	"log"
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

// 将区块序列化成字节数组
func (block *Block) Serialize() []byte {
	var result bytes.Buffer

	enconder := gob.NewEncoder(&result)

	err := enconder.Encode(block)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

// 反序列化
func DeserializeBlock(blockBytes []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))

	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}

//1. 工厂方法：创建新区块
func NewBlock(height int64, prevBlockHash []byte, data string) *Block{

	// 创建新区块
	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil, 0}

	// 调用工作量证明的方法并且返回有效的Hash和Nonce
	pow := NewProofOfWork(block)

	// 挖矿验证
	hash, nonce := pow.run()

	block.Hash = hash[:]
	block.Nonce = nonce

	fmt.Println()

	return block
}

//2. 单独写创建创世区块的方法
func CreateGenesisBlock(data string) *Block {
	return NewBlock(1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, data)
}

