// http://people.bath.ac.uk/ensab/Primes/
//
// Diophantine Representation of the Set of Prime Numbers
//
// By Daniel Rodríguez <http://sadasant.com/>
// 2012-10-01
//

package main

import (
	M "math"
	"flag"
	"fmt"
	"strconv"
)

func factorial(n int64) int64 {
	if n > 1 {
		return factorial(n-1) * n
	}
	return n
}

func Diophantine(n int64) float64 {
	return M.Abs(M.Pow(M.Cos(M.Pi), float64(2))*(float64(factorial(n-1)+1)/float64(n)))
}

func main() {
	flag.Parse()
	n, _ := strconv.ParseInt(flag.Arg(0), 10, 32)
	isPrime := Diophantine(n)
	fmt.Printf("Is prime? %v", isPrime)
}
