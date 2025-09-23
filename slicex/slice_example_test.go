package slicex_test

import (
	"fmt"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

func ExampleFilter() {
	list := []int64{1, 2, 3, 4}

	result := slicex.Filter(list, func(i int64) bool {
		return i%2 == 0
	})

	fmt.Printf("%v", result)
	// Output: [2 4]
}

//func ExampleFilterWithIndex() {
//	list := []int64{1, 2, 3, 4}
//
//	result := slicex.FilterWithIndex(list)
//}

//
//Map
//
//MapWithIndex
//
//UniqueMap
//
//Bind
