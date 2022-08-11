package main

import (
	"bufio"
	"fmt"
	"gopl.io/ch2/convformat"
	"gopl.io/ch2/tempconv"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) > 0 {
		for _, arg := range arguments {
			printArg(arg)
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			printArg(input.Text())
		}
	}
}

func printArg(argument string) {
	float, err := strconv.ParseFloat(argument, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "convformatter: %v\n", err)
		os.Exit(1)
	}
	printTemp(float)
	printDist(float)
	printMass(float)
}

func printTemp(temp float64) {
	f := tempconv.Fahrenheit(temp)
	c := tempconv.Celsius(temp)
	k := tempconv.Kelvin(temp)
	fmt.Printf("Temp: %s = %s, %s;  %s = %s, %s;  %s = %s, %s\n",
		c, tempconv.CToF(c), tempconv.CToK(c),
		f, tempconv.FToC(f), tempconv.FToK(f),
		k, tempconv.KToC(k), tempconv.KToF(k),
	)
}

func printDist(dist float64) {
	f := convformat.Foot(dist)
	m := convformat.Metre(dist)
	fmt.Printf("Distance: %s = %s;  %s = %s\n",
		f, convformat.FToM(f),
		m, convformat.MToF(m),
	)
}

func printMass(mass float64) {
	pound := convformat.Pound(mass)
	kg := convformat.Kilogram(mass)
	fmt.Printf("Mass: %s = %s;  %s = %s\n",
		pound, convformat.PToK(pound),
		kg, convformat.KToP(kg),
	)
}
