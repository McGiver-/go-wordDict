package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	dict "github.com/McGiver-/go-wordDict/dictionary"
)

type node struct {
	Val  int
	Next *node
}

const dictionaryFile = "./dictionary.gob"
const wordsFile = "/usr/share/dict/american-english"
var topN = 3
var routines = 100


func main() {
	d := dict.Dictionary()
	ws := []string{}
	for word := range readWordsFile(wordsFile) {
		ws = append(ws,word)
		d.Add(word)
	}
	nbWords := len(ws)
	wg := sync.WaitGroup{}
	wg.Add(nbWords)
	startTime := time.Now()
	for _, word := range ws{
		go func(word string,d *dict.Dict,w *sync.WaitGroup){
			d.SearchN(word,topN)
			w.Done()
		}(word,d,&wg)
	}
	// d := dict.Dictionary()
	// wordChan := readWordsFile(wordsFile)

	// for word := range wordChan {
	// 	d.Add(word)
	// }

	// for _, v := range d.String() {
	// 	fmt.Println(v)
	// }

	// err := writeGob(dictionaryFile, d)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	wg.Wait()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Printf("took %s to return top %d words for %d words" ,
	elapsed,
	topN,
	nbWords)
}

func print(str string, dict *dict.Dict, wg *sync.WaitGroup) {
	fmt.Println(fmt.Sprintf(strings.Join(dict.Search(str), "\n")))
	wg.Done()
}

func printN(str string, n int, dict *dict.Dict, wg *sync.WaitGroup) {
	fmt.Println(fmt.Sprintf(strings.Join(dict.SearchN(str, n), "\n")))
	wg.Done()
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
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
