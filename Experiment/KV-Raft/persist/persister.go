package persist

import (
	"fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

type Persister struct {
	path string
	db   *leveldb.DB
}

func (p *Persister) Init(path string) {
	var err error
	p.path = path
	p.db, err = leveldb.OpenFile(path, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Persister) Put(key string, value string) {
	err := p.db.Put([]byte(key), []byte(value), nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Persister) Get(key string) string {
	value, err := p.db.Get([]byte(key), nil)
	if err != nil {
		return ""
	}
	return string(value)
}

func (p *Persister) Close() {
	err := p.db.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func (p *Persister) PrintStrVal(key string) {
	value := p.Get(key)
	fmt.Println(value)
}

// Test leveldb
//func main() {
//	p := &Persister{}
//	p.Init("../KV-Raft")
//	p.Put("1", "a")
//	fmt.Println(p.Get("1"))
//	err := p.db.Close()
//	if err != nil {
//		log.Fatalln(err)
//	}
//}
