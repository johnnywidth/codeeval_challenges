package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		r, try := Add(n, 1)
		fmt.Printf("%d %d\n", try, r)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func Add(n int, try int) (int, int) {
	rn := strconv.Itoa(n)
	rn = Reverse(rn)

	rni, _ := strconv.Atoi(rn)
	r := n + rni

	rs := strconv.Itoa(r)
	left, right := HalfAndHalf(rs)

	if left == Reverse(right) {
		return r, try
	}

	return Add(r, try+1)
}

func Reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func HalfAndHalf(s string) (left string, right string) {
	l := len(s)
	switch {
	case l == 1:
		return s, s
	case l%2 == 0:
		left = s[:(l / 2)]
		right = s[(l / 2):]
	case l%2 == 1:
		left = s[:(l / 2)]
		right = s[(l/2 + 1):]
	}
	return left, right
}
