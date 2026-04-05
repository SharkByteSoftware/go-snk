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
	var (
		numMap1 = map[int]string{0: "zero", 8: "one", 2: "two"}
		numMap2 = map[int]string{3: "three", 12: "four", 256: "five"}
	)

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

	result := mapx.Filter(numMap, func(key int, _ string) bool {
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

func ExampleMapKeys() {
	scores := map[string]int{
		"alice": 90,
		"bob":   75,
		"carol": 88,
	}

	result := mapx.MapKeys(scores, func(key string) string {
		return "user_" + key
	})

	fmt.Println(result)
	// Output: map[user_alice:90 user_bob:75 user_carol:88]
}

func ExamplePartition() {
	var numMap = map[int]string{
		0:   "zero",
		2:   "two",
		3:   "three",
		8:   "one",
		12:  "four",
		256: "five",
	}

	even, odd := mapx.Partition(numMap, func(key int, _ string) bool {
		return key%2 == 0
	})

	fmt.Println(even)
	fmt.Println(odd)
	// Output: map[0:zero 2:two 8:one 12:four 256:five]
	// map[3:three]
}

func ExampleCountBy() {
	inventory := map[string]int{
		"apple":      5,
		"banana":     12,
		"cherry":     3,
		"date":       8,
		"elderberry": 1,
	}

	result := mapx.CountBy(inventory, func(_ string, qty int) string {
		if qty < 5 {
			return "low"
		}

		return "ok"
	})

	fmt.Println(result)
	// Output: map[low:2 ok:3]
}

func ExampleMerge_keepLeft() {
	defaults := map[string]int{"timeout": 30, "retries": 3, "workers": 5}
	overrides := map[string]int{"timeout": 60, "debug": 1}

	// caller-supplied values win; defaults fill in the rest
	result := mapx.Merge(overrides, defaults, func(_ string, left, _ int) int {
		return left
	})

	fmt.Println(result["timeout"]) // from overrides
	fmt.Println(result["retries"]) // from defaults, no conflict
	fmt.Println(result["debug"])   // from overrides, no conflict
	// Output:
	// 60
	// 3
	// 1
}

func ExampleMerge_sum() {
	pageViews := map[string]int{"home": 100, "about": 40}
	moreViews := map[string]int{"home": 50, "contact": 20}

	result := mapx.Merge(pageViews, moreViews, func(_ string, left, right int) int {
		return left + right
	})

	fmt.Println(result["home"])    // combined
	fmt.Println(result["about"])   // no conflict
	fmt.Println(result["contact"]) // no conflict
	// Output:
	// 150
	// 40
	// 20
}
