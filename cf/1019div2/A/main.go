package main

import "fmt"

func main() {
	numCases := 0
	fmt.Scan(&numCases)
	for range numCases {
		l := 0
		fmt.Scan(&l)
		m := deduplicate(l)
		fmt.Println(len(m))
	}
}

func deduplicate(l int) map[int]struct{} {
	m := make(map[int]struct{})
	for range l {
		tmp := 0
		fmt.Scan(&tmp)
		m[tmp] = struct{}{}
	}
	return m
}
