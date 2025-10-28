package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCase1(t *testing.T) {
	assert.Equal(t, "aaabcbc", decodeString("3[a]2[bc]"))
}

func TestCase2(t *testing.T) {
	assert.Equal(t, "accaccacc", decodeString("3[a2[c]]"))
}

func TestCase3(t *testing.T) {
	assert.Equal(t, "abc", decodeString("abc"))
}
