package main

import (
	"fmt"
)

var id []int
var sz []int
var count int

func Find(p int) int {
	for p != id[p] {
		p = id[p]
	}
	return p
}

func Union(p, q int) {
	r1 := Find(p)
	r2 := Find(q)

	if r1 != r2 {
		if sz[r1] <= sz[r2] {
			id[r1] = r2
			sz[r2] += sz[r1]
		} else {
			id[r2] = r1
			sz[r1] += sz[r2]
		}
	}
	count--
}

func Connected(p, q int) bool {
	var flag bool
	if Find(p) == Find(q) {
		flag = true
	}
	flag = false
	return flag
}

/*func Counter() int {
	return count
}
*/

func main() {
	var N int
	fmt.Scan(&N)
	count = N

	id = make([]int, N)
	sz = make([]int, N)
	for i := 0; i < N; i++ {
		id[i] = i
		sz[i] = 1
	}

	fmt.Println(id)
	dataSize := 8
	for i := 0; i < dataSize; i++ {
		var p, q int
		fmt.Scan(&p, &q)
		if Connected(p, q) == false {
			Union(p, q)
		}
	}

	fmt.Println(id)
	fmt.Println(count)
}
