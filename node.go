package dictionary

import (
	"sort"
)

type node struct {
	Char     byte
	Freq     int
	Children map[byte]*node
	IsWord   bool
}

func (n *node) add(word string) bool {
	w := word[0]
	next, exist := n.Children[w]
	if !exist {
		next = &node{w, 0, make(map[byte]*node), false}
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
		if n.Char == 0 {
			for _, l := range v.string() {
				str = append(str, l)
			}
		} else {
			for _, l := range v.string() {
				str = append(str, string(n.Char)+l)
			}
		}
	}
	return str
}

func (n *node) stringN(topN *int) (str []string) {
	if n == nil {
		return nil
	}
	var children []*node
	if *topN <= 0 {
		return nil
	}
	if len(n.Children) == 0 {
		*topN--
		return []string{string(n.Char)}
	}

	if n.IsWord {
		*topN--
		str = append(str, string(n.Char))
	}

	for _, v := range n.Children {
		children = append(children, v)
	}

	sort.Slice(children, func(i, j int) bool {
		return children[i].Freq > children[j].Freq
	})

	for _, v := range children {
		for _, l := range v.stringN(topN) {
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
	next, exist := n.Children[c]
	if !exist {
		return false
	}
	return next.update(part[1:])
}

func (n *node) search(part string) (str []string) {
	start := n.findPart(part)
	if start == nil {
		return nil
	}
	return start.string()
}

func (n *node) searchN(part string, topN *int) (str []string) {
	start := n.findPart(part)
	if start == nil {
		return nil
	}
	return start.stringN(topN)
}

func (n *node) findPart(part string) *node {
	if len(part) == 0 {
		return n
	}
	c := part[0]
	next, exist := n.Children[c]
	if !exist {
		return nil
	}
	return next.findPart(part[1:])
}
