package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var input [][]string

func main() {
	if err := prepareInputData(os.Args[1]); err != nil {
		log.Fatal(err)
	}
	for _, _input := range input {
		fizzNumber, _ := strconv.Atoi(_input[0])
		buzzNumber, _ := strconv.Atoi(_input[1])
		length, _ := strconv.Atoi(_input[2])
		var _output []string
		for i := 1; i <= length; i++ {
			var _fbOutput string
			if i%fizzNumber == 0 {
				_fbOutput += "F"
			}
			if i%buzzNumber == 0 {
				_fbOutput += "B"
			}
			if _fbOutput != "" {
				_output = append(_output, _fbOutput)
			} else {
				_output = append(_output, strconv.Itoa(i))
			}
		}
		outputLine := strings.Join(_output, " ")
		fmt.Println(outputLine)
	}
}

func prepareInputData(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineNumbers := strings.Split(scanner.Text(), " ")
		if len(lineNumbers) != 3 {
			return errors.New("Invalid input data")
		}
		input = append(input, lineNumbers)
	}
	return scanner.Err()
}
