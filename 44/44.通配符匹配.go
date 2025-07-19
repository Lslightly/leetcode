package main

/*
 * @lc app=leetcode.cn id=44 lang=golang
 *
 * [44] 通配符匹配
 */
// @lc code=start

type Node struct {
	accept bool
	out    map[byte]*Node
}

func (n *Node) setNext(ch byte, nextN *Node) {
	n.out[ch] = nextN
}

func createNFA(p string) *Node {
	root := &Node{
		accept: false,
		out:    make(map[byte]*Node),
	}
	if len(p) == 0 {
		root.accept = true
		return root
	}
	cur := root
	for i := range len(p) {
		ch := p[i]
		switch ch {
		case '*':
			cur.setNext('*', cur)
		default: // '?'/a/b/c
			node := &Node{
				accept: false,
				out:    make(map[byte]*Node),
			}
			cur.setNext(ch, node)
			cur = node
		}
		if i == len(p)-1 {
			cur.accept = true
		}
	}
	return root
}

func walkNFA(s string, root *Node) bool {
	wl := make([]*Node, 0)
	wl = append(wl, root)
	for i := range s {
		ch := s[i]
		if len(wl) == 0 {
			break
		}
		nextWL := make([]*Node, 0, len(wl))
		for _, n := range wl {
			for _, ch := range []byte{ch, '?', '*'} {
				if next, ok := n.out[ch]; ok {
					nextWL = append(nextWL, next)
				}
			}
			wl = nextWL
		}
	}
	for _, n := range wl {
		if n.accept {
			return true
		}
	}
	return false
}

func isMatch(s string, p string) bool {
	return walkNFA(s, createNFA(p))
}

// @lc code=end
