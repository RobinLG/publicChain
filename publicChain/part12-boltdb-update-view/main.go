package main

import (
	"github.com/boltdb/bolt"
	"log"
)

//安装boltdb，在终端输入：go get "github.com/boltdb/bolt"

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	//打开或创建数据库
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Read-write transactions
	err = db.Update(func(tx *bolt.Tx) error {
		//获取BlockBucket表对象
		b := tx.Bucket([]byte("BlockBucket"))

		//往表里插入数据
		if b != nil {
			err := b.Put([]byte("ll"),  []byte("Send 1000 BTC To Ronnie....."))
			if err != nil {
				log.Panic("数据存储失败.....")
			}
		}

		// 返回nil，以便数据库处理相应操作
		return nil
	})
	//更新失败
	if err != nil {
		log.Panic(err)
	}

}