package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"strings"
)

var (
	channel   string
	chaincode string
	key       string

	dbpath string
)

func init() {
	flag.StringVar(&channel, "channel", "mychannel", "Channel name")
	flag.StringVar(&chaincode, "chaincode", "mychaincode", "Chaincode name")
	flag.StringVar(&key, "key", "", "Key to query; empty query all keys")

	flag.StringVar(&dbpath, "dbpath", "", "Path to LevelDB")
}

func readAll(db *leveldb.DB) {
	var b bytes.Buffer
	b.WriteString(channel)
	b.WriteByte(0)
	b.WriteString(chaincode)
	prefix := b.String()

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		if strings.HasPrefix(key, prefix) {
			value := string(iter.Value())
			fmt.Printf("Key[%s]=[%s]\n", key, value)
		}
	}
	iter.Release()
	//err := iter.Error()
}

func main() {

}
