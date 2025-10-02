package main

import "fmt"

func printcap(s []int) {
	fmt.Println(cap(s))
}

func main() {
	s := make([]int, 0, 4)
	printcap(s)
	s = append(s, 0)
	printcap(s)
	s = append(s, 0)
	s[0] = 1
	printcap(s)
	s2 := append(s, 0)
	s2[0] = 2
	fmt.Println("s", s2)
	printcap(s2)
	s = append(s2, 0)
	s[0] = 3
	fmt.Println("s", s)
	fmt.Println("s2", s2)
	printcap(s)
	printcap(s2)
	s2 = append(s2, 1)
	fmt.Println("s", s) // s is affected
	fmt.Println("s2", s2)
	printcap(s)
	printcap(s2)
	s = append(s, 0)
	printcap(s)

}
