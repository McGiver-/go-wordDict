package dictionary

import (
	"container/heap"
	"fmt"
)

type node struct {
	Char     byte
	Freq     int
	Children map[byte]*node
	IsWord   bool
}

type priorityQueue []*node

func (q priorityQueue) Len() int {
	return len(q)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].Freq < p[j].Freq
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (h *priorityQueue) Push(x interface{}) {
	*h = append(*h, x.(*node))
}

func (h *priorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Dict struct {
	Root *node
}

func Dictionary() *Dict {
	return &Dict{&node{0, 0, map[byte]*node{}, false}}
}

func (dict *Dict) Add(word string) bool {
	return dict.Root.add(word)
}

func (dict *Dict) Update(part string) bool {
	return dict.Root.update(part)
}

func (dict *Dict) String() []string {
	return dict.Root.string()
}

func (dict *Dict) Search(part string) (res []string) {
	p := part[:len(part)-1]
	for _, v := range dict.Root.search(part) {
		res = append(res, p+v)
	}
	return res
}

func (dict *Dict) SearchN(part string, topN int) (res []string) {
	p := part[:len(part)-1]
	pq := &priorityQueue{}
	heap.Init(pq)
	for _, v := range dict.Root.searchN(part, pq, &topN) {
		res = append(res, p+v)
	}
	return res
}

func (n *node) add(word string) bool {
	w := word[0]
	next, exist := n.Children[w]
	if !exist {
		next = &node{w, 0, map[byte]*node{}, false}
		n.Children[w] = next
	}

	if len(word) == 1 {
		next.IsWord = true
		return true
	}
	return next.add(word[1:])
}

func (n *node) string() (str []string) {
	if len(n.Children) == 0 {
		return []string{string(n.Char)}
	}

	if n.IsWord {
		str = append(str, string(n.Char))
	}

	for _, v := range n.Children {
		for _, l := range v.string() {
			str = append(str, string(n.Char)+l)
		}
	}
	return str
}

func (n *node) update(part string) bool {
	if n.Char != 0 {
		n.Freq++
	}
	if len(part) == 0 {
		return true
	}
	c := part[0]
	next := n.Children[c]
	return next.update(part[1:])
}

func (n *node) search(part string) (str []string) {
	start, err := n.findPart(part)
	if err != nil {
		return []string{}
	}
	return start.string()
}

func (n *node) searchN(part string, pq *priorityQueue, topN *int) (str []string) {
	start, err := n.findPart(part)
	if err != nil {
		return []string{}
	}

	if *topN <= 0 {
		return []string{}
	}

	for _, v := range n.Children {
		pq.Push(v)
	}

	for *topN > 0 {

	}
	for k := range children {
		if *topN <= 0 {
			return str
		}
		*topN--
		str = append(str, children[k].stringN(pq, topN)...)
	}
}

func (n *node) findPart(part string) (*node, error) {
	if len(part) == 0 {
		return n, nil
	}
	c := part[0]
	next, exist := n.Children[c]

	if !exist {
		return nil, fmt.Errorf("could not find this word")
	}
	return next.findPart(part[1:])
}
