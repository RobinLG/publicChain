package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"math/big"
)

//数据库名字
const dbName = "blockchain.db"
//表名
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte //新区块的Hash
	DB *bolt.DB //数据库
}

// 遍历输出所有区块的信息
func (blc *Blockchain) Printchain() {

	var block *Block
	var currentHash []byte = blc.Tip

	for {
		err := blc.DB.View(func(tx *bolt.Tx) error {
			// 1. 表
			b := tx.Bucket([]byte(blockTableName))
			if b != nil {
				// 获取当前区块的字节数组
				blockBytes := b.Get(currentHash)
				// 反序列化
				block = DeserializeBlock(blockBytes)

				fmt.Printf("Height: %d\n", block.Height)
				fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)
				fmt.Printf("Nonce: %d\n", block.Nonce)

			}
			return nil
		})

		if err != nil {
			log.Panic(err)
		}

		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

		currentHash = block.PrevBlockHash
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
			err = b.Put([]byte("l"), newBlock.Hash)
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
func CreateBlockchainWithGenesisBlock() *Blockchain {

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	//打开或创建数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	var blockHash []byte

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