package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/nonlinear/hashtable/hash"
)

func main() {
	dict := hash.ValueHashTable{}
	for i := 10; i < 100; i++ {
		dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	fmt.Println(dict)
}
