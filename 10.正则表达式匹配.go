package main

import "log"

/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
// @lc code=start

var idCnt int = 0

type Node struct {
	id  int
	acc bool
	out map[byte]Nodes
}

func (n *Node) EpsilonOut() Nodes {
	if out, ok := n.out['_']; !ok {
		return Nodes{}
	} else {
		return out
	}
}

func (n *Node) SingleEpsilonOut() *Node {
	out := n.EpsilonOut()
	if len(out) != 1 {
		log.Fatalf("should be 1 single ε out")
	}
	return out[0]
}

func NewNode() *Node {
	idCnt += 1
	return &Node{
		id:  idCnt,
		acc: false,
		out: make(map[byte]Nodes),
	}
}

func (n *Node) closure() Nodes {
	res := make(Nodes, 0)
	for _, out := range n.EpsilonOut() {
		res = append(res, out.closure()...)
	}
	return append(res, n)
}

func (n *Node) Nexts(ch byte) Nodes {
	res := make(Nodes, 0)
	nClosure := n.closure()
	for _, n := range nClosure {
		if nexts, ok := n.out[ch]; ok {
			res.AppendClosureOf(nexts)
		}
		if nexts, ok := n.out['.']; ok {
			if len(nexts) != 1 {
				log.Fatalf("dot should points to itself")
			}
			res.AppendClosureOf(nexts)
		}
	}
	return res
}

// use nextStartIdx to judge whether at end
func (n *Node) createNexts(p string, idx int) (nextStartIdx int, nextNode *Node) {
	if idx+1 < len(p) && p[idx+1] == '*' { // judge whether next is asterisk
		iterNode := NewNode()
		n.out['_'] = Nodes{iterNode}
		iterNode.out[p[idx]] = append(iterNode.out[p[idx]], iterNode)
		nextNode = NewNode()
		iterNode.out['_'] = Nodes{nextNode}
		if idx+2 == len(p) {
			nextNode.acc = true
		}
		return idx + 2, nextNode
	}
	nextNode = NewNode()
	n.out[p[idx]] = append(n.out[p[idx]], nextNode)
	if idx+1 == len(p) {
		nextNode.acc = true
	}
	return idx + 1, nextNode
}

type Nodes []*Node

func (ns Nodes) Empty() bool {
	return len(ns) == 0
}

func (ns *Nodes) AppendClosureOf(nexts Nodes) {
	for _, next := range nexts {
		*ns = append(*ns, next.closure()...)
	}
}

// return the start node of NFA
func CreateNFA(p string) *Node {
	start := NewNode()
	cur := start
	idx := 0
	for {
		idx, cur = cur.createNexts(p, idx)
		if idx == len(p) {
			return start
		}
	}
}

func panicAtCloseOk() {
	recover()
}

func WalkNFA(s string, startNode *Node) bool {
	succ := make(chan bool)
	add := make(chan int)   // add multiple
	done := make(chan bool) // sub once a time
	defer func() {
		close(succ)
		close(add)
		close(done)
	}()
	var counter int

	finish := func() {
		defer panicAtCloseOk() // recover for done
		if err := recover(); err == nil {
			done <- true
		}
	}
	var consume func(idx int, n *Node)
	consume = func(idx int, n *Node) {
		defer finish()

		if idx == len(s) {
			return
		}

		nexts := n.Nexts(s[idx]) // consume s[idx]
		if nexts.Empty() {
			return
		}

		add <- len(nexts) // block the remaining and done channel
		for _, next := range nexts {
			if next.acc && idx == len(s)-1 {
				succ <- true
				return
			}
			go consume(idx+1, next)
		}
	}

	counter = 1
	go consume(0, startNode)

	for {
		select {
		case <-succ:
			return true
		default:
			if counter == 0 {
				return false
			}
		}
		select {
		case k := <-add:
			counter += k
		case <-done:
			counter -= 1
		default: // non blocking here
		}
	}
}

func isMatch(s string, p string) bool {
	return WalkNFA(s, CreateNFA(p))
}

// @lc code=end
