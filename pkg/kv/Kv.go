package kv

import (
	"go.etcd.io/bbolt"
	"os"
)

//https://pkg.go.dev/go.etcd.io/bbolt#readme-bbolt
//db, err := bolt.Open("my.db", 0600, nil)
func NewBBolt(path string, mode os.FileMode, options *bbolt.Options) (*bbolt.DB, error)  {
	return bbolt.Open(path,mode,options)
}
