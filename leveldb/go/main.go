package main

import (
    "fmt"
    "github.com/syndtr/goleveldb/leveldb"
)

func main() {
    db, _ := leveldb.OpenFile("/Users/JC/Documents/project/go/leveldb/db/test.db", nil)

    key := []byte("hello")
    // value := []byte("hi i'm levelDB by go")
    // e := db.Put(key, value, nil)
    // fmt.Println("e==>", e)
    data, _ := db.Get(key, nil)
    fmt.Println("data==>", data)
    fmt.Printf("%c\n",data)
    
    defer db.Close()
}



