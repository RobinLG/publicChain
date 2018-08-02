package main

import (
	"github.com/boltdb/bolt"
	"log"
)

//安装boltdb，在终端输入：go get "github.com/boltdb/bolt"

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}