package kv

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"os"
)

type BBoltKv struct {
	BBkv *bbolt.DB
}

type ViewValueCallBack func(k,v []byte) error


//https://pkg.go.dev/go.etcd.io/bbolt#readme-bbolt
//db, err := bolt.Open("my.db", 0600, nil)
func NewBBolt(path string, mode os.FileMode, options *bbolt.Options) (*BBoltKv, error)  {
	b,er:= bbolt.Open(path,mode,options)
	return &BBoltKv{BBkv:b},er
}


func (bb *BBoltKv)PutWithBkt(bkt string,k,v []byte) error {

	return bb.BBkv.Update(func(tx *bbolt.Tx) error {
		bk:=tx.Bucket([]byte(bkt))
		return bk.Put(k,v)
	})
}

func (bb *BBoltKv)PutWithBktManyJson(bkt string,mp map[string]interface{}) error {

	return bb.BBkv.Update(func(tx *bbolt.Tx) error {
		bk:=tx.Bucket([]byte(bkt))
		for k,v:=range mp{
			bt,err:=json.Marshal(v)
			if err!=nil {
				return err
			}
			err=bk.Put([]byte(k),bt)
			if err!=nil {
				return err
			}
		}
		return nil
	})
}

func (bb *BBoltKv)PutWithBktStrKey(bkt string,k string,v []byte) error {
	return bb.PutWithBkt(bkt,[]byte(k),v)
}

func (bb *BBoltKv)PutWithBktStrKeyJson(bkt string,k string,v interface{}) error {
	bt,err:=json.Marshal(v)
	if err!=nil {
		return err
	}
	return bb.PutWithBktStrKey(bkt,k,bt)
}

func (bb *BBoltKv)PutWithBktByteKeyJson(bkt string,k []byte ,v interface{}) error {
	return bb.PutWithBktStrKeyJson(bkt,string(k),v)
}

func (bb *BBoltKv)GetByStrKeyValueJson(bkt string,k string,v interface{})  error {
	return bb.BBkv.View(func(tx *bbolt.Tx) error {
		bk:=tx.Bucket([]byte(bkt))
		bt:=bk.Get([]byte(k))
		return json.Unmarshal(bt,v)
	})
}

func (bb *BBoltKv)GteInter(bkt string,va ViewValueCallBack) error {
	 return  bb.BBkv.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bkt))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			err:=va(k,v)
			if err!=nil {
				return err
			}
		}
		return nil
	})
}

func (bb *BBoltKv)GetLast(bkt string,va ViewValueCallBack) error {
	return  bb.bkv.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bkt))
		c := b.Cursor()
		return va(c.Last())
	})
}