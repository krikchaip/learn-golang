package main

import (
	"fmt"
	"reflect"
)

type names []string   // defined type
type uintarr = []uint // type alias

func main() {
	// ?? declaring a slice, equals "nil" by default
	var cities []string
	fmt.Printf("%v, %#v\n", cities, cities) // => [], []string(nil)

	// ?? creating a slice using the `make` built-in function
	// ** returns T with zero-values initilized
	nums := make([]int, 2)
	fmt.Println(nums) // => [0 0]

	// ?? creating a slice using the `new` built-in function
	// ** returns a pointer to T
	nums2 := new([]int)
	fmt.Println(nums2)                    // => &[]
	fmt.Printf("%#v, %T\n", nums2, nums2) // => &[]int(nil), *[]int

	// ?? declaring a slice using a slice literal
	ds := []int{2, 3, 4, 5}
	uints := uintarr{2, 4, 6, 8}
	friends := names{"Dan", "Maria"}

	fmt.Println(ds)
	fmt.Printf("%v, %#v\n", uints, uints)     // => []uint{0x2, 0x4, 0x6, 0x8}
	fmt.Printf("%v, %#v\n", friends, friends) // => main.names{"Dan", "Maria"}

	// ** slices are behave like lists in other language
	var n []int
	n = ds // n holds the SAME REFERENCE as ds
	n[0] = 555
	fmt.Println(ds) // => [555 3 4 5]

	// ?? COMPARING SLICES - can only compare slices to `nil`

	// ** uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// ** empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) // false

	// ** this will cause a compile error
	// ** as we can't compare two slices using `=` operator
	// fmt.Println(n == ds)

	// ** do this instead
	fmt.Println(SliceEquals(n, ds)) // true

	// ** or this
	fmt.Println(reflect.DeepEqual(n, ds)) // true
}

func SliceEquals[T comparable](xs, ys []T) bool {
	if len(xs) != len(ys) {
		return false
	}

	for i, v := range xs {
		if v != ys[i] {
			return false
		}
	}

	return true
}
