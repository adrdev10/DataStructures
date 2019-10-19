package main

import (
	"fmt"

	"github.com/xdragon1015/DataStructures/linear/set"
)

func main() {
	s := set.NewSet()
	s.AddElement(10)
	s.AddElement(100)
	s.AddElement(10000)
	s.AddElement(5)
	s.AddElement(10)
	fmt.Println("First set:", s)

	fmt.Println(s.Contains(10))

	s2 := set.NewSet()
	s2.AddElement(10)
	s2.AddElement(100)
	s2.AddElement(1000000)
	fmt.Println("Another set", s2)

	fmt.Println("Intersecting sets")

	s3 := s.InterSect(s2)
	fmt.Println(s3)
	s4 := set.NewSet()
	s4.AddElement(10)
	s4.AddElement(190)
	s4.AddElement(79)
	s5 := s.UnionSet(s4)
	fmt.Println("Union operation", s5)
}
