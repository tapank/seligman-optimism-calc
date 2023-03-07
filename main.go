package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print(INSTRUCTIONS)
	getByte()
	survey := New(ADULT)
	for q := survey.Next(); q != ""; q = survey.Next() {
		fmt.Print(q)
		for {
			if survey.Add(getByte()) {
				break
			}
			fmt.Print("Invalid. Enter a or b: ")
		}
	}
	survey.PrintResult()
}

func getByte() rune {
	reader := bufio.NewReaderSize(os.Stdin, 1)
	in, _ := reader.ReadByte()
	return rune(in)
}
