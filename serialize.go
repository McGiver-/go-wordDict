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

	dict "local/wordDictionary/dictionary"
)

type node struct {
	Val  int
	Next *node
}

const dictionaryFile = "./dictionary.gob"
const wordsFile = "/usr/share/dict/american-english"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
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

	var n = new(dict.Dict)
	err := readGob(dictionaryFile, n)
	startTime := time.Now()
	if err != nil {
		fmt.Println(err)
	} else {
		go print("helpful", n, &wg)
		go print("miracle", n, &wg)
	}
	wg.Wait()
	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	fmt.Println("elapsed time ", elapsed)
}

func print(str string, dict *dict.Dict, wg *sync.WaitGroup) {
	fmt.Println(fmt.Sprintf(strings.Join(dict.Search(str), "\n")))
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
