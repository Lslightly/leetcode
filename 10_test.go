package main

import "testing"

func TestCreateNFA1(t *testing.T) {
	start := CreateNFA("a*")
	if start == nil {
		t.Fatalf("create nfa failed")
	}
	/*
		start -> 1 -> 2
		1 -a-> 1
		1 -epsilon-> 2
	*/
	s1 := start.EpsilonOut()[0]
	if len(s1.EpsilonOut()) != 1 {
		t.Fatalf("s1's out len should be 1")
	}
	if s1.out['a'][0] != s1 {
		t.Fatalf("s1 should transfer to itself by a")
	}
	s2 := s1.EpsilonOut()[0]
	if !s2.acc {
		t.Fatalf("s2.acc should be true")
	}
}

func TestCraeteNFA(t *testing.T) {
	start := CreateNFA("a")
	if start.acc == true {
		t.Fatalf("start should not be acc")
	}
	if len(start.out['a']) != 1 {
		t.Fatalf("should transfer to another state")
	}
	if start.out['a'][0].acc != true {
		t.Fatalf("next state should be acc")
	}
}

func TestCreateNFADotAsterisk(t *testing.T) {
	start := CreateNFA(".*")
	if start == nil {
		t.Fatalf("create nfa failed")
	}
	if start.acc != true {
		t.Fatalf("start should be acc")
	}
	nexts, ok := start.out['.']
	if !ok {
		t.Fatalf("should contain .")
	}
	if nexts.Empty() {
		t.Fatalf("nexts should not be empty")
	}
	if len(nexts) != 1 {
		t.Fatalf("len of nexts should be 1")
	}
	if nexts[0] != start {
		t.Fatalf("nexts[0] should be start")
	}
}

func Test10(t *testing.T) {
	type pair struct {
		s, p  string
		match bool
	}
	inout := []pair{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
	}
	for _, pair := range inout {
		if !(isMatch(pair.s, pair.p) == pair.match) {
			t.Fail()
		}
	}
}

func Test10_1(t *testing.T) {
	if isMatch("aa", "a") {
		t.Fail()
	}
}

func Test10_2(t *testing.T) {
	if !isMatch("aa", "a*") {
		t.Fail()
	}
}

func Test10_3(t *testing.T) {
	if !isMatch("aa", ".*") {
		t.Fail()
	}
}

func Test36(t *testing.T) {
	if isMatch("aaba", "ab*a*c*a") {
		t.Fail()
	}
}

func Test28FNA(t *testing.T) {
	start := CreateNFA("c*a*b")
	c := start.SingleEpsilonOut()
	ca := c.SingleEpsilonOut()
	a := ca.SingleEpsilonOut()
	ab := a.SingleEpsilonOut()
	end := ab.out['b'][0]
	if !end.acc {
		t.Fail()
	}
	if c.out['c'][0] != c {
		t.Fail()
	}
	if a.out['a'][0] != a {
		t.Fail()
	}
}

func Test28(t *testing.T) {
	if !isMatch("aab", "c*a*b") {
		t.Fail()
	}
}

func Test34(t *testing.T) {
	if !isMatch("abbbcd", "ab*bbbcd") {
		t.Fail()
	}
}
