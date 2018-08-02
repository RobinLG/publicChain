package main

import (
	"robin/publicChain/part14-block-boltdb/BLC"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	block := BLC.NewBlock(1,[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, }, "Test")
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	//打开或创建数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("blocks"))

		if b == nil {
			b, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic("Blocks table create fail.....")
			}
		}

		err = b.Put([]byte("l"), block.Serialize())
		if err != nil {
			log.Panic(err)
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	err = db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte("blocks"))
		if b != nil {
			blockData := b.Get([]byte("l"))
			//fmt.Println(blockData)
			//fmt.Printf("%s\n", blockData)
			block := BLC.DeserializeBlock(blockData)
			fmt.Println(block)
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}
