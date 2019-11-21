package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/thimi0412/gopl.io/ch2/q2/conv"
)

const defaultConv = "t"

var c string

func main() {
	flag.StringVar(&c, "type", defaultConv, "convert type")
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		args = append(args, input.Text())
	}
	for _, arg := range args {
		intInput, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		showConv(c, intInput)
	}
}

func showConv(convType string, value int) {
	if convType == "t" {
		f := conv.Fahrenheit(value)
		c := conv.Celsius(value)
		fmt.Printf("%s = %s, %s = %s\n", f, conv.FToC(f), c, conv.CToF(c))
	} else if convType == "l" {
		m := conv.Meter(value)
		f := conv.Feet(value)
		fmt.Printf("%s = %s, %s = %s\n", m, conv.MToFe(m), f, conv.FeToM(f))
	} else if convType == "w" {
		k := conv.Kilogram(value)
		p := conv.Pound(value)
		fmt.Printf("%s = %s, %s = %s\n", k, conv.KtoP(k), p, conv.PtoK(p))
	} else {
		fmt.Println("No Type!")
	}
}
