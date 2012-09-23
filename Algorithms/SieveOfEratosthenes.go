// Daniel R. (sadasant.com)
// 23/09/2012
//
// How to run:
//   go run SieveOfEratosthenes.go 11
//      [1 2 3 5 7]
//
// License:
//   Public Domain
//

package main

import (
	"flag"
	"fmt"
	"strconv"
)

// Sieve of Eratosthenes
func primesLowerThan(n int64) []int64 {
	bools := make([]bool, n+1)
	primes := make([]int64, n+1)
	primes[0] = 1
	c := 1
	for i := int64(2); i < n; i++ {
		if !bools[i] {
			bools[i] = true
			primes[c] = i
			c++
			for j := i + i; j <= n; j += i {
				bools[j] = true
			}
		}
	}
	return primes[:c]
}

func main() {
	flag.Parse()
	max, _ := strconv.ParseInt(flag.Arg(0), 10, 64)
	primes := primesLowerThan(max)
	fmt.Printf("%v", primes)
}
