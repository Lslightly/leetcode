package main

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	root Node
}

type Node struct {
	outs   map[rune]*Node
	isWord bool
}

func Constructor() Trie {
	return Trie{
		root: Node{
			outs: make(map[rune]*Node),
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := &this.root
	for _, ch := range word {
		if next, ok := cur.outs[ch]; ok {
			cur = next
		} else {
			newNode := &Node{
				outs: make(map[rune]*Node),
			}
			cur.outs[ch] = newNode
			cur = newNode
		}
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	cur := &this.root
	for _, ch := range word {
		if next, ok := cur.outs[ch]; !ok {
			return false
		} else {
			cur = next
		}
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := &this.root
	for _, ch := range prefix {
		if next, ok := cur.outs[ch]; !ok {
			return false
		} else {
			cur = next
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
