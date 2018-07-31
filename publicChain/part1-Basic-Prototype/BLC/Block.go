package BLC

import (
	"time"
	"strconv"
	"fmt"
	"bytes"
	"crypto/sha256"
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
}

func (block *Block) setHash() {
	//1. Height -> byte[]
	heightBytes := IntToHex(block.Height)
	fmt.Println(heightBytes)

	//2. Timestamp -> byte[]
	timeString := strconv.FormatInt(block.Timestamp, 2)
	fmt.Println(timeString)

	timeBytes := []byte(timeString)
	fmt.Println(timeBytes)

	//3. 拼接所有属性
	blockBytes := bytes.Join([][]byte{heightBytes, block.PrevBlockHash, block.Data, timeBytes, block.Hash}, []byte{})

	//4. 生成Hash
	hash := sha256.Sum256(blockBytes)

	block.Hash = hash[:]

	}

//1. 工厂方法：创建新区块
func NewBlock(height int64, prevBlockHash []byte, data string) *Block{

	block := &Block{height, prevBlockHash, []byte(data), time.Now().Unix(), nil}
	fmt.Println(block)
	block.setHash()

	return block
}

