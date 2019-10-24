package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileMap := make(map[string]map[string]int)
	files := os.Args[1:]

	for _, arg := range files {
		counts := make(map[string]int)
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, fileMap)
		f.Close()
	}

	for fileName, counts := range fileMap {
		fmt.Println(fileName)
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileMap map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileMap[f.Name()] = counts
	}
}
