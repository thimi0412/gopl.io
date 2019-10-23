package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func printArgs(args []string) {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	secs := time.Since(start).Seconds()
	fmt.Println(s)
	fmt.Printf("実行時間: %g\n", secs)
}

func printArgsUseJoin(args []string) {
	start := time.Now()
	fmt.Println(strings.Join(args[1:], " "))
	secs := time.Since(start).Seconds()
	fmt.Printf("実行時間: %g\n", secs)
}

func main() {
	args := os.Args[1:]
	printArgs(args)
	printArgsUseJoin(args)
}
