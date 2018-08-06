package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//数据库名字
const dbName = "blockchain.db"
//表名
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte //新区块的Hash
	DB *bolt.DB
}

//// 增加区块到区块链里
//func (blc *Blockchain) AddBlockToBlockchain(hight int64, preHash []byte, data string) {
//	// 创建新区块
//	newBlock := NewBlock(hight, preHash, data)
//	// 往区块链里添加区块
//	blc.Blocks = append(blc.Blocks, newBlock)
//}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	//打开或创建数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucket([]byte(blockTableName))

		if err != nil {
			log.Panic(err)
		}

		if b == nil {
			// 创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data....")
			// 将创世区块存储到表中
			err := b.Put(genesisBlock.Hash, genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			//存储最新区块的Hash
			err = b.Put([]byte("new"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}

		return nil
	})

	return &Blockchain{blockHash, db}
}