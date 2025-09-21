package examples

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

//FilterWithIndex
//
//Map
//
//MapWithIndex
//
//UniqueMap
//
//Bind
