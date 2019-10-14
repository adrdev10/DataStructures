package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/linear/set"
)

func main() {
	im := make(map[int]bool)
	s := set.NewSet(im)
	s.AddElement(10)
	s.AddElement(100)
	s.AddElement(10000)
	s.AddElement(5)
	s.AddElement(10)

	fmt.Println(s.Contains(10))

	fmt.Println(s)
}
