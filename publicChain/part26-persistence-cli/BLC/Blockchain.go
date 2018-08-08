package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"fmt"
	"time"
	"os"
)

//数据库名字
const dbName = "blockchain.db"
//表名
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte //新区块的Hash
	DB *bolt.DB //数据库
}

func (blockchain *Blockchain) Iterator() *BlockIterator {
	return &BlockIterator{blockchain.Tip,blockchain.DB}
}

//判断数据库是否存在
func dbExists() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}

	return true
}

// 遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {

	blockIterator := blc.Iterator()

	for {

		block := blockIterator.Next()

		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timpstamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}
	}
}


// 增加区块到区块链里
func (blc *Blockchain) AddBlockToBlockchain(data string) {

	err := blc.DB.Update(func(tx *bolt.Tx) error {
		// 1. 获取表
		b := tx.Bucket([]byte(blockTableName))
		if b != nil {

			// 获取最新区块
			blockBytes := b.Get(blc.Tip)
			// 反序列化
			block := DeserializeBlock(blockBytes)

			// 2. 创建新区块
			newBlock := NewBlock(block.Height + 1, block.Hash, data)
			// 3. 将新区块序列化并且存储到数据库中
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			// 4. 更新数据库里"l"对应的hash
			err = b.Put([]byte("new"), newBlock.Hash)
			// 5. 更新Tip
			blc.Tip = newBlock.Hash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

}

//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock(data string) {
	if dbExists() {
		fmt.Println("创世区块已经存在....")
		os.Exit(1)
	}

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	//打开或创建数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("正在创建创世区块....")

	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(blockTableName))
		if  b == nil {
			b, err = tx.CreateBucket([]byte(blockTableName))

			if err != nil {
				log.Panic(err)
			}
		}


		if b != nil {
			// 创建创世区块
			genesisBlock := CreateGenesisBlock(data)
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

		}

		return nil
	})

}