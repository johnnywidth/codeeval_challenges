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
var output [][]string

func main() {
	if err := prepareInputData(); err != nil {
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
		output = append(output, _output)
	}
	if err := writeOutputDataToFile(); err != nil {
		log.Fatal(err)
	}
}

func prepareInputData() error {
	file, err := os.Open("./input.txt")
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

func writeOutputDataToFile() error {
	file, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, _output := range output {
		outputLine := strings.Join(_output, " ")
		fmt.Fprintln(w, outputLine)
		if err := w.Flush(); err != nil {
			return err
		}
	}
	return nil
}
