// Daniel R. (sadasant.com)
// 25/09/2012
//
// Test it with:
//   go run sort.go
//
// License:
//   Public Domain
//

package main // TODO Make an Algorithms package

import (
	"fmt"
	"time"
)

func Insertion(array []int) []int {
	var now int
	for i, v := range array {
		now = i
		for now > 0 && array[now-1] > v {
			array[now] = array[now-1]
			now--
		}
		array[now] = v
	}
	return array
}

func Bubble(array []int) []int {
	var current, prev int
	var swapped bool
	l := len(array)
repeat:
	swapped = false
	for i := 1; i < l; i++ {
		current = array[i]
		prev = array[i-1]
		if prev > current {
			array[i] = prev
			array[i-1] = current
			swapped = true
		}
	}
	if swapped {
		goto repeat
	}
	return array
}

func Quick(array []int) []int {
	if (len(array) <= 1) {
		return array
	}
	var less, equal, greater []int
	pivot := array[0]
	for _, v := range array[0:] {
		switch {
		case v < pivot:
			less = append(less, v)
		case v == pivot:
			equal = append(equal, v)
		case v > pivot:
			greater = append(greater, v)
		}
	}
	mix := make([]int, 0)
	mix = append(mix, Quick(less)...)
	mix = append(mix, equal...)
	return append(mix, Quick(greater)...)
}

func Merge(array []int) []int {
	l := len(array)
	if l < 2 {
		return array
	}
	half := l / 2
	return merge(Merge(array[:half]), Merge(array[half:]))
}
func merge(a, b []int) []int {
	la, lb := len(a), len(b)
	if la == 0 {
		return b
	}
	if lb == 0 {
		return a
	}
	mix := make([]int, 1)
	switch {
	case a[0] < b[0]:
		mix[0] = a[0]
		return append(mix, merge(a[1:], b)...)
	default:
		mix[0] = b[0]
		return append(mix, merge(b[1:], a)...)
	}
	return mix
}

func main() {
	// Arrays
	array := []int{2, 5, 4, 16, -15, -22, 239323, -123123}

	// Time
	start := time.Now()

	println("Testing Insertion Sort")
	start = time.Now()
	fmt.Printf("  %v\n", Insertion(array))
	fmt.Println("  Lasted: ", time.Since(start))

	// Functions affect the original arrays
	array = []int{2, 5, 4, 16, -15, -22}

	println("\nTesting Bubble Sort")
	start = time.Now()
	fmt.Printf("  %v\n", Bubble(array))
	fmt.Println("  Lasted: ", time.Since(start))

	// Functions affect the original arrays
	array = []int{2, 5, 4, 16, -15, -22}

	println("\nTesting Quick Sort")
	start = time.Now()
	fmt.Printf("  %v\n", Quick(array))
	fmt.Println("  Lasted: ", time.Since(start))

	// Functions affect the original arrays
	array = []int{2, 5, 4, 16, -15, -22}

	println("\nTesting Merge Sort")
	start = time.Now()
	fmt.Printf("  %v\n", Merge(array))
	fmt.Println("  Lasted: ", time.Since(start))
}
