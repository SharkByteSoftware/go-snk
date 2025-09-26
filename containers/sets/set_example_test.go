package sets_test

import (
	"fmt"
	"slices"

	"github.com/SharkByteSoftware/go-snk/containers/sets"
)

func ExampleNew() {
	set := sets.New[int](1, 2, 3, 3, 4, 4)

	values := set.Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2 3 4]
}

func ExampleSet_Add() {

}

//func ExampleAdd() {
//	s := sets.New[int]()
//
//	s.Add(1, 2, 2, 3, 3)
//
//	values := s.Values()
//	slices.Sort(values)
//
//	fmt.Println(values)
//	// Output [1 2 3]
//}
