package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Words []Word

type Word struct {
	Len int
	Str string
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineNumber int
	var words Words

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if lineNumber == 0 {
			lineNumber, _ = strconv.Atoi(scanner.Text())
			continue
		}

		words = append(words, Word{Len: len([]rune(scanner.Text())), Str: scanner.Text()})
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	sort.Sort(words)

	if lineNumber > len(words) {
		lineNumber = len(words)
	}

	for i := 0; i < lineNumber; i++ {
		fmt.Println(words[i].Str)
	}
}

func (w Words) Len() int { return len(w) }

func (w Words) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func (w Words) Less(i, j int) bool { return w[i].Len > w[j].Len }
