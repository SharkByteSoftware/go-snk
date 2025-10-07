package mapx_test

import (
	"fmt"
	"slices"

	"github.com/SharkByteSoftware/go-snk/mapx"
)

func ExampleKeys() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	keys := mapx.Keys(numMap)

	slices.Sort(keys)

	fmt.Println(keys)
	// Output: [0 2 3 8 12 256]
}

func ExampleValues() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	values := mapx.Values(numMap)

	slices.Sort(values)

	fmt.Println(values)
	// Output: [five four one three two zero]
}

func ExampleUniqueValues() {
	var dupMap = map[int]string{
		0:    "zero",
		8:    "zero",
		2:    "two",
		3:    "five",
		12:   "five",
		256:  "five",
		8192: "five",
	}

	values := mapx.UniqueValues(dupMap)

	slices.Sort(values)

	fmt.Println(values)
	// Output: [five two zero]
}

func ExampleContains() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	result1 := mapx.Contains(numMap, 0)
	result2 := mapx.Contains(numMap, 1)
	result3 := mapx.Contains(numMap, 2, 8, 256)

	fmt.Println(result1, result2, result3)
	// Output: true false true
}

func ExampleValueOr() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	result1 := mapx.ValueOr(numMap, 0, "nothing")
	result2 := mapx.ValueOr(numMap, 1, "nothing")

	fmt.Println(result1, result2)
	// Output: zero nothing
}

func ExampleInvert() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	inverted := mapx.Invert(numMap)

	fmt.Println(inverted)
	// Output: map[five:256 four:12 one:8 three:3 two:2 zero:0]
}

func ExampleCombine() {
	var numMap1 = map[int]string{0: "zero", 8: "one", 2: "two"}
	var numMap2 = map[int]string{3: "three", 12: "four", 256: "five"}

	result := mapx.Combine(numMap1, numMap2)

	fmt.Println(result)
	// Output: map[0:zero 2:two 3:three 8:one 12:four 256:five]
}

func ExampleToSlice() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	values := mapx.ToSlice(numMap, func(key int, value string) string {
		return fmt.Sprintf("%d-%s", key, value)
	})

	slices.Sort(values)

	fmt.Println(values)
	// Output: [0-zero 12-four 2-two 256-five 3-three 8-one]
}

func ExampleFilter() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	result := mapx.Filter(numMap, func(key int, value string) bool {
		return key%2 != 0
	})

	fmt.Println(result)
	// Output: map[3:three]
}

func ExampleApply() {
	var numMap = map[int]string{
		0:   "zero",
		8:   "one",
		2:   "two",
		3:   "three",
		12:  "four",
		256: "five",
	}

	var result string
	mapx.Apply(numMap, func(key int, value string) {
		if key == 256 {
			result = value
		}
	})

	fmt.Println(result)
	// Output: five
}
