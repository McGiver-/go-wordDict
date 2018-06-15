package dictionary

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
)

var words = []string{
	"cat",
	"call",
	"caller",
	"offer",
	"offend",
	"timeline",
}

func TestDict_Add(te *testing.T) {

	words := []string{
		"cat",
		"cap",
		"ca",
		"con",
	}

	test := Dictionary()
	expect := Dictionary()

	for _, v := range words {
		test.Add(v)
	}
	p := &node{[]byte("p")[0], 0, map[byte]*node{}, true}
	t := &node{[]byte("t")[0], 0, map[byte]*node{}, true}
	n := &node{[]byte("n")[0], 0, map[byte]*node{}, true}
	a := &node{[]byte("a")[0], 0, map[byte]*node{
		[]byte("t")[0]: t,
		[]byte("p")[0]: p}, true}
	o := &node{[]byte("o")[0], 0, map[byte]*node{[]byte("n")[0]: n}, false}
	c := &node{[]byte("c")[0], 0, map[byte]*node{
		[]byte("o")[0]: o,
		[]byte("a")[0]: a}, false}

	root := &node{0, 0, map[byte]*node{
		[]byte("c")[0]: c,
	}, false}

	expect.Root = root
	if !reflect.DeepEqual(test, expect) {
		testStr := test.String()
		expectStr := expect.String()
		sort.Strings(testStr)
		sort.Strings(expectStr)
		te.Errorf("test : %s \n and expected : %s \n not equal", testStr, expectStr)
	}
}

func TestDict_Update(te *testing.T) {
	words := []string{
		"cat",
		"cap",
		"ca",
		"con",
	}

	test := Dictionary()
	expect := Dictionary()

	for _, v := range words {
		test.Add(v)
	}
	for _, v := range words {
		test.Update(v)
	}

	test.Update("cat")
	test.Update("cat")
	test.Update("cat")
	test.Update("con")
	test.Update("con")
	p := &node{[]byte("p")[0], 1, map[byte]*node{}, true}
	t := &node{[]byte("t")[0], 4, map[byte]*node{}, true}
	n := &node{[]byte("n")[0], 3, map[byte]*node{}, true}
	a := &node{[]byte("a")[0], 6, map[byte]*node{
		[]byte("t")[0]: t,
		[]byte("p")[0]: p}, true}
	o := &node{[]byte("o")[0], 3, map[byte]*node{[]byte("n")[0]: n}, false}
	c := &node{[]byte("c")[0], 9, map[byte]*node{
		[]byte("o")[0]: o,
		[]byte("a")[0]: a}, false}

	root := &node{0, 0, map[byte]*node{
		[]byte("c")[0]: c,
	}, false}

	expect.Root = root
	if !reflect.DeepEqual(test, expect) {
		te.Errorf("test : %v \n and expected : %v \n not equal", test, expect)
	}
}

func TestDict_String(t *testing.T) {

	test := Dictionary()

	for _, v := range words {
		test.Add(v)
	}
	expect := test.String()
	sort.Strings(words)
	sort.Strings(expect)

	if !reflect.DeepEqual(words, expect) {
		t.Errorf("words : %s \n and expected : %s \n not equal", words, expect)
	}
}

func TestDict_Search(t *testing.T) {
	test := Dictionary()
	for _, word := range words {
		test.Add(word)
	}
	for _, word := range words {
		if len(test.Search(word)) == 0 {
			t.Errorf("could not search word %s", word)
		}
	}

	if len(test.Search("c")) != 3 {
		t.Errorf("could not search word %s", "c")
	}
	if len(test.Search("o")) != 2 {
		t.Errorf("could not search word %s", "o")
	}
	if len(test.Search("q")) != 0 {
		t.Errorf("could not search word %s", "o")
	}
}

func TestDict_SearchN(t *testing.T) {
	test := Dictionary()
	for _, word := range words {
		test.Add(word)
	}
	for _, word := range words {
		test.Update(word)
	}
	for _, word := range words {
		if len(test.SearchN(word, 1)) != 1 {
			t.Errorf("could not search word %s", word)
		}
	}

	test.Update("caller")
	test.Update("caller")
	test.Update("cat")
	test.Update("con")
	test.Update("con")

	if !reflect.DeepEqual(test.SearchN("ca", 2), []string{
		"call",
		"caller",
	}) {
		t.Errorf("expecting call and caller, got %s", test.SearchN("ca", 2))
	}

	if !reflect.DeepEqual(test.SearchN("ca", 10), []string{
		"call",
		"caller",
		"cat",
	}) {
		t.Errorf("expecting call, caller and cat got %s", test.SearchN("ca", 3))
	}

	if !reflect.DeepEqual(test.SearchN("cal", 1), []string{
		"call",
	}) {
		t.Errorf("expecting call got %s", test.SearchN("cal", 1))
	}

	if len(test.SearchN("call", 0)) != 0 {
		t.Errorf("could not search word %s", "call")
	}
	if len(test.SearchN("q", 1)) != 0 {
		t.Errorf("could not search word %s", "o")
	}
}

/*-----------------------------BENCHMARKS ----------------------------------*/

func BenchmarkAdd(b *testing.B) {
	var wordsFile = "/usr/share/dict/american-english"
	d := Dictionary()
	wordChan := readWordsFile(wordsFile)
	ws := []string{}
	// wc := make(chan string)
	f := func(r rune) bool {
		return r < 'a' || r > 'z' || r == '\''
	}
	for word := range wordChan {
		if strings.IndexFunc(word, f) == -1 {
			ws = append(ws, word)
		}
	}
	b.Run("chan", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, word := range ws {
				d.Add(word)
			}
		}
	})
}

func BenchmarkSearch(b *testing.B) {
	var wordsFile = "/usr/share/dict/american-english"
	d := Dictionary()
	wordChan := readWordsFile(wordsFile)
	ws := []string{}
	for word := range wordChan {
		ws = append(ws, word)
		d.Add(word)
	}

	b.Run("chan", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			for _, word := range ws {
				d.Search(word)
			}
		}
	})
}

func BenchmarkSearchN(b *testing.B) {
	var wordsFile = "/usr/share/dict/american-english"
	d := Dictionary()
	wordChan := readWordsFile(wordsFile)
	ws := []string{}
	for word := range wordChan {
		ws = append(ws, word)
		d.Add(word)
	}

	b.Run("chan", func(b *testing.B) {
		for m := 0; m < b.N; m++ {
			for _, word := range ws {
				d.SearchN(word, 1)
			}
		}
	})

}

func readWordsFile(filePath string) <-chan string {
	out := make(chan string)
	go func() {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		close(out)
	}()

	return out
}
