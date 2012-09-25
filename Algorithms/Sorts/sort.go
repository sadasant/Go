// Daniel R. (sadasant.com)
// 25/09/2012
//
// Example test:
//   go run sort.go 40000000
//   Testing Insertion Sort
//   
//   Testing Bubble Sort
//   
//   Testing Quick Sort
//   
//   Testing Merge Sort
//   
//   
//   Merge Sort Lasted:  3us
//   Quick Sort Lasted:  2us
//   Bubble Sort Lasted:  3us
//   Insertion Sort Lasted:  3us
//   
//   Arrays were too large to print
//
// License:
//   Public Domain
//

package main // TODO Make an Algorithms package

import (
	"flag"
	"strconv"
	"fmt"
	"time"
	"math/rand"
)

func Insertion(array []int) []int {
	var now int
	answer := make([]int, len(array))
	copy(answer, array)
	for i, v := range answer {
		now = i
		for now > 0 && answer[now-1] > v {
			answer[now] = answer[now-1]
			now--
		}
		answer[now] = v
	}
	return answer
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
	flag.Parse()
	max, _ := strconv.Atoi(flag.Arg(0))

	// Too damn high
	tooLarge := false

	if (max > 10000) {
		defer println("\nArrays were too large to print\n")
		tooLarge = true
	}

	// Array
	array := make([]int, max)
	for i := 0; i < max; i++ {
		array[i] = rand.Int()
	}

	// Time
	start := time.Now()

	println("Testing Insertion Sort")
	start = time.Now()
	if !tooLarge {
		fmt.Printf("  %v\n", Insertion(array))
	}
	defer fmt.Println("Insertion Sort Lasted: ", time.Since(start))

	println("\nTesting Bubble Sort")
	start = time.Now()
	if !tooLarge {
		fmt.Printf("  %v\n", Bubble(array))
	}
	defer fmt.Println("Bubble Sort Lasted: ", time.Since(start))

	println("\nTesting Quick Sort")
	start = time.Now()
	if !tooLarge {
		fmt.Printf("  %v\n", Quick(array))
	}
	defer fmt.Println("Quick Sort Lasted: ", time.Since(start))

	println("\nTesting Merge Sort")
	start = time.Now()
	if !tooLarge {
		fmt.Printf("  %v\n", Merge(array))
	}
	defer fmt.Println("Merge Sort Lasted: ", time.Since(start))

	println("\n")
}
