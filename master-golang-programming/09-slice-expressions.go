package main

import (
	"fmt"
)

func main() {
	// ?? Arrays, Slices and Strings are sliceable
	a := [5]int{1, 2, 3, 4, 5} // [5]int
	s1 := []int{2, 4, 6, 8}    // []int
	str := "Winner"            // string

	b := a[1:3]    // slicing an array -> returns a slice
	ss := str[3:6] // slicing a string -> returns a string

	fmt.Printf("Type: %T, Value: %#v\n", b, b)   // Type: []int, Value: []int{2, 3}
	fmt.Printf("Type: %T, Value: %#v\n", ss, ss) // Type: string, Value: "ner"

	fmt.Println(s1[2:]) // [6 8], same as s1[2:len(s1)]
	fmt.Println(s1[:3]) // [2 4 6], same as s1[0:3]
	fmt.Println(s1[:])  // [2 4 6 8], same as s1[0:len(s1)]

	// ?? slice expression application
	// fmt.Printf("%p %p\n", &s1[0], s1[:3])
	// apd := s1[:3]
	// fmt.Printf("%p %p\n", apd, append(apd, 100))
	fmt.Println(append(s1[:3], 100)) // [2 4 6 100]

	// ** Go implements a slice as data structure called `Slice Header`.
	// ** Slice Header contains 3 fields (24bytes in total):
	// ** - the address of the backing array (pointer) -> 8bytes.
	// ** - the length of the slice.  The built-in function len() returns it -> 8bytes.
	// ** - the capacity of the slice. The size of the backing array after the slice first element.
	// **   cap() built-in function returns it -> 8bytes.

	// ** A nil slice doesn't have backing array, so all the fields in the slice header are equal to zero.

	// ?? when a slice is created by slicing an array,
	// ?? that array becomes the backing array of the new slice.
	arr1 := [4]int{10, 20, 30, 40}
	slice1, slice2 := arr1[0:2], arr1[1:3]

	arr1[1] = 2                       // ** modifying the array
	fmt.Println(arr1, slice1, slice2) // -> [10 2 30 40] [10 2] [2 30]

	// ?? a slice expression doesn't create a new backing array.
	// ?? The original and the returned slice are connected!
	s1 = []int{10, 20, 30, 40, 50}
	s3, s4 := s1[0:2], s1[1:3] // s3, s4 share the same backing array with s1

	s3[1] = 600
	fmt.Println(s1) // -> [10 600 30 40 50]
	fmt.Println(s4) // -> [600 30]
}
