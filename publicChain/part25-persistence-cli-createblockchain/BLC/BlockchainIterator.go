package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//迭代器
type BlockIterator struct {
	CurrentHash []byte
	DB *bolt.DB
}

func (blockchainIterator *BlockIterator) Next() *Block {

	var block *Block

	err := blockchainIterator.DB.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))

		if b != nil {
			currentHash := b.Get(blockchainIterator.CurrentHash)
			//获取当前迭代器里currentHash对应的区块
			block = DeserializeBlock(currentHash)

			//更新迭代器里的currentHash
			blockchainIterator.CurrentHash = block.PrevBlockHash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return block
}
