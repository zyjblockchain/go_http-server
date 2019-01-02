package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {
	var t []byte
	var v []byte
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucketIfNotExists([]byte("bucket"))
		t = b.Get([]byte("name"))
		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("bucket"))
		v = b.Get([]byte("mm"))
		return nil
	})

	if t == nil {
		fmt.Println("t", t)
	}
	fmt.Println("v", v)
	// db.View(func(tx *bolt.Tx) error {
	//
	// })
}
