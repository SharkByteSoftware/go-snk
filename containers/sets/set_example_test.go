package sets_test

import (
	"fmt"
	"slices"
	"sort"

	"github.com/SharkByteSoftware/go-snk/containers/sets"
)

func ExampleNew() {
	set := sets.New(1, 2, 3, 3, 4, 4)

	values := set.Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2 3 4]
}

func ExampleSet_Add() {
	set := sets.New[int]()

	set.Add(1, 2, 3)

	values := set.Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2 3]
}

func ExampleSet_Contains() {
	set := sets.New(1, 2, 3, 3, 4, 4)

	fmt.Println(set.Contains(1), set.Contains(5))
	// Output: true false
}

func ExampleSet_Remove() {
	set := sets.New(1, 2, 3, 3, 4, 4)

	set.Remove(1)
	set.Remove(4)

	values := set.Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [2 3]
}

func ExampleSet_Intersect() {
	set1 := sets.New(1, 2, 3, 3, 4, 4)
	set2 := sets.New(3, 4, 5, 6, 7, 8)

	values := set1.Intersect(set2).Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [3 4]
}

func ExampleSet_Union() {
	set1 := sets.New(1, 2, 3, 3, 4, 4)
	set2 := sets.New(3, 4, 5, 6, 7, 8)

	values := set1.Union(set2).Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2 3 4 5 6 7 8]
}

func ExampleSet_Difference() {
	set1 := sets.New(1, 2, 3, 3, 4, 4)
	set2 := sets.New(3, 4, 5, 6, 7, 8)

	values := set1.Difference(set2).Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2]
}

func ExampleSet_SymmetricDifference() {
	set1 := sets.New(1, 2, 3, 3, 4, 4)
	set2 := sets.New(3, 4, 5, 6, 7, 8)

	values := set1.SymmetricDifference(set2).Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2 5 6 7 8]
}

func ExampleSet_Subset() {
	set1 := sets.New(1, 2, 3, 3, 4, 4)
	set2 := sets.New(3, 4, 5, 6, 7, 8)
	set3 := sets.New(3, 4)

	fmt.Println(set3.Subset(set1), set3.Subset(set2), set1.Subset(set2))
	// Output: true true false
}

func ExampleSet_Apply() {
	set := sets.New(1, 2, 3)

	var values []int
	set.Apply(func(i int) {
		values = append(values, i)
	})

	sort.Ints(values)

	fmt.Println(values)
	// Output: [1 2 3]
}

func ExampleSet_Equals() {
	set1 := sets.New(1, 2, 3)
	set2 := sets.New(3, 4, 5)
	set3 := sets.New(1, 2, 3)

	fmt.Println(set1.Equals(set2), set1.Equals(set3))
	// Output: false true
}

func ExampleSet_Clone() {
	original := sets.New(1, 2, 3)
	clone := original.Clone()

	clone.Remove(1)

	originalValues := original.Values()
	slices.Sort(originalValues)

	cloneValues := clone.Values()
	slices.Sort(cloneValues)

	fmt.Println(originalValues)
	fmt.Println(cloneValues)
	// Output:
	// [1 2 3]
	// [2 3]
}

func ExampleSet_IsEmpty() {
	set := sets.New[int]()

	fmt.Println(set.IsEmpty())
	set.Add(1)
	fmt.Println(set.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleSet_Clear() {
	set := sets.New(1, 2, 3)

	set.Clear()

	fmt.Println(set.IsEmpty())
	// Output: true
}

func ExampleSet_ToJSON() {
	set := sets.New(42)

	b, _ := set.ToJSON()

	fmt.Println(string(b))
	// Output: [42]
}

func ExampleSet_FromJSON() {
	set := sets.New[int]()

	_ = set.FromJSON([]byte(`[1,2,3]`))

	values := set.Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [1 2 3]
}

func ExampleSet_MarshalJSON() {
	set := sets.New(99)

	b, _ := set.MarshalJSON()

	fmt.Println(string(b))
	// Output: [99]
}

func ExampleSet_UnmarshalJSON() {
	set := sets.New[string]()

	_ = set.UnmarshalJSON([]byte(`["a","b","c"]`))

	values := set.Values()
	slices.Sort(values)

	fmt.Println(values)
	// Output: [a b c]
}
