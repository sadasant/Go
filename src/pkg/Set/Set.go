// Daniel R. (sadasant.com)
// 2012-10-02
//

package main //set

import "fmt"

func Union(a, b []int) []int {
	la, lb := len(a), len(b)
	lc := la + lb
	r := make([]int, lc)
	i, j, c := 0, 0, 0
	for c < lc {
		r[c] = a[i]
		r[c+1] = b[j]
		c += 2
		i++
		j++
	}
	return r
}

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8}
	fmt.Printf("%v", Union(a, b))
}
