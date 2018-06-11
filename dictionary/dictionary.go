package dictionary

type Dict struct {
	Root *node
}

func Dictionary() *Dict {
	return &Dict{&node{0, 0, make(map[byte]*node), false}}
}

func (dict *Dict) Add(word string) bool {
	return dict.Root.add(word)
}

func (dict *Dict) Update(word string) bool {
	return dict.Root.update(word)
}

func (dict *Dict) String() []string {
	return dict.Root.string()
}

func (dict *Dict) Search(word string) (res []string) {
	p := word[:len(word)-1]
	for _, v := range dict.Root.search(word) {
		res = append(res, p+v)
	}
	return res
}

func (dict *Dict) SearchN(word string, topN int) (res []string) {
	p := word[:len(word)-1]
	for _, v := range dict.Root.searchN(word, &topN) {
		res = append(res, p+v)
	}
	return res
}
