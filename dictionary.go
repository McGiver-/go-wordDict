package dictionary

type Dict struct {
	Root *node
}

//Dictionary returns the pointer to an empty dictionary
func Dictionary() *Dict {
	return &Dict{&node{0, 0, make(map[byte]*node), false}}
}

//Add adds a string to the dictionary
func (dict *Dict) Add(word string) bool {
	return dict.Root.add(word)
}

//Update increases the frequency of the input word by 1
func (dict *Dict) Update(word string) bool {
	return dict.Root.update(word)
}

//String returns a slice of all the strings in the dictionary
func (dict *Dict) String() (str []string) {
	for _, b := range dict.Root.string() {
		str = append(str, b.string())
	}
	return str
}

//Search returns a slice of strings that have the input string as a suffix
func (dict *Dict) Search(word string) (res []string) {
	p := word[:len(word)-1]
	for _, v := range dict.Root.search(word) {
		res = append(res, p+v.string())
	}
	return res
}

//SearchN returns a slice of the top n strings that have the input string as a suffix. These are returned
//in order of highest frequency first
func (dict *Dict) SearchN(word string, topN int) (res []string) {
	p := word[:len(word)-1]
	for _, v := range dict.Root.searchN(word, &topN) {
		res = append(res, p+v.string())
	}
	return res
}
