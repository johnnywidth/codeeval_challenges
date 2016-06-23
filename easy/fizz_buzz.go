package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input [][]string

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumbers := strings.Split(scanner.Text(), " ")
		if len(lineNumbers) != 3 {
			log.Fatal("Invalid input data")
		}
		input = append(input, lineNumbers)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	for _, value := range input {
		fizzNumber, _ := strconv.Atoi(value[0])
		buzzNumber, _ := strconv.Atoi(value[1])
		length, _ := strconv.Atoi(value[2])

		var output []string
		for i := 1; i <= length; i++ {
			var fbOutput string
			if i%fizzNumber == 0 {
				fbOutput += "F"
			}
			if i%buzzNumber == 0 {
				fbOutput += "B"
			}
			if fbOutput != "" {
				output = append(output, fbOutput)
			} else {
				output = append(output, strconv.Itoa(i))
			}
		}

		fmt.Println(strings.Join(output, " "))
	}
}
