package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Op int

const (
	Insert Op = iota
	Search
	StartsWith
)

type instruction struct {
	op       Op
	str      string
	expected bool
}

func testcase(t *testing.T, insts []instruction) {
	trie := Constructor()
	for i, inst := range insts {
		switch inst.op {
		case Insert:
			trie.Insert(inst.str)
		case Search:
			assert.Equal(t, inst.expected, trie.Search(inst.str), fmt.Sprintln("inst", i, inst))
		case StartsWith:
			assert.Equal(t, inst.expected, trie.StartsWith(inst.str), fmt.Sprintln("inst", i, inst))
		}
	}
}

func TestCase1(t *testing.T) {
	testcase(
		t,
		[]instruction{
			{Insert, "apple", false},
			{Search, "apple", true},
			{Search, "app", false},
			{StartsWith, "app", true},
			{Insert, "app", false},
			{Search, "app", true},
		},
	)
}
