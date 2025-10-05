package mapx_test

import (
	"fmt"
	"slices"

	"github.com/SharkByteSoftware/go-snk/mapx"
)

func ExampleKeys() {
	var numberMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	keys := mapx.Keys(numberMap)

	slices.Sort(keys)

	fmt.Println(keys)
	// Output: [0 2 3 8 12 256]
}

func ExampleValues() {
	var numberMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	values := mapx.Values(numberMap)

	slices.Sort(values)

	fmt.Println(values)
	// Output: [five four one three two zero]
}

func ExampleUniqueValues() {
	var dupValueMap = map[int]string{
		0:    "zero",
		8:    "zero",
		2:    "two",
		3:    "five",
		12:   "five",
		256:  "five",
		8192: "five",
	}

	values := mapx.UniqueValues(dupValueMap)

	slices.Sort(values)

	fmt.Println(values)
	// Output: [five two zero]
}

func ExampleContains() {
	var numberMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	result1 := mapx.Contains(numberMap, 0)
	result2 := mapx.Contains(numberMap, 1)
	result3 := mapx.Contains(numberMap, 2, 8, 256)

	fmt.Println(result1, result2, result3)
	// Output: true false true
}

func ExampleValueOr() {
	var numberMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	result1 := mapx.ValueOr(numberMap, 0, "nothing")
	result2 := mapx.ValueOr(numberMap, 1, "nothing")

	fmt.Println(result1, result2)
	// Output: zero nothing
}

func ExampleInvert() {
	var numberMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	inverted := mapx.Invert(numberMap)

	fmt.Println(inverted)
	// Output: map[five:256 four:12 one:8 three:3 two:2 zero:0]
}
