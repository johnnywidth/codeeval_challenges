package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ",")
		fmt.Println(RemoveCharacters(strings.TrimSpace(l[0]), strings.TrimSpace(l[1])))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func RemoveCharacters(s, e string) string {
	var b bytes.Buffer
	b.WriteString(s)

	for _, v := range strings.Split(e, "") {
		r := strings.Replace(b.String(), v, "", -1)
		b.Reset()
		b.WriteString(r)
	}

	return b.String()
}
