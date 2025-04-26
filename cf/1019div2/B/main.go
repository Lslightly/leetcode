package main

import (
	"fmt"
	"regexp"
)

func main() {
	t := 0
	fmt.Scan(&t)
	for range t {
		each()
	}
}

func each() {
	n := 0
	fmt.Scan(&n)
	s := ""
	fmt.Scan(&s)
	s = "0" + s
	diff := 0
	for i := 1; i <= n; i++ {
		if s[i] != s[i-1] {
			diff++
		}
	}
	total := n + diff
	for _, pat := range []string{"01", "10"} {
		regex, _ := regexp.Compile(pat)
		subs := regex.FindAllString(s, -1)
		if len(subs) >= 2 {
			total -= 2
			break
		}
	}
	fmt.Println(total)
}
