package main

import (
	"bufio"
	"fmt"
	"lab1"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input number: ")

	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(text[:len(text)-1])
	if err != nil {
		panic(err)
	}

	numeral, err := lab1.DecimalToRomanNumeral(num)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %s", numeral)

}
