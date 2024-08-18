package main

import (
	"fmt"
	"github.com/InsulaLabs/testa"
)

func main() {

	db, err := testa.Open("my.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if e := db.Set("name", []byte("rudolph")); e != nil {
		panic(e)
	}

	v, e := db.Get("name")
	if e != nil {
		panic(e)
	}

	fmt.Println("Got: ", string(v))
}
