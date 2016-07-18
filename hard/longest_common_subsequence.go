package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

type LCSResult []string

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), ";")
		if len(a) < 2 {
			continue
		}

		fmt.Println(LCS(a[0], a[1]))
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func LCS(l, r string) string {
	l, r = SpaceMap(l), SpaceMap(r)
	result := Fill(l, r)

	var lcs LCSResult
	x, y := len(l), len(r)

	for x >= 1 && y >= 1 {
		if l[x-1] == r[y-1] {
			lcs = append(lcs, string(l[x-1]))
			x, y = x-1, y-1
		} else if result[x-1][y] > result[x][y-1] {
			x -= 1
		} else {
			y -= 1
		}
	}

	sort.Sort(lcs)

	return strings.Join(lcs, "")
}

func Fill(l, r string) [][]int {
	result := make([][]int, len(l)+1)
	result[0] = make([]int, len(r)+1)

	for i := 1; i <= len(l); i++ {
		for j := 1; j <= len(r); j++ {
			if len(result[i]) == 0 {
				result[i] = make([]int, len(r)+1)
			}

			if l[i-1] == r[j-1] {
				result[i][j] = result[i-1][j-1] + 1
			} else {
				result[i][j] = Max(result[i][j-1], result[i-1][j])
			}
		}
	}

	return result
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (l LCSResult) Len() int { return len(l) }

func (l LCSResult) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

func (l LCSResult) Less(i, j int) bool { return i > j }
