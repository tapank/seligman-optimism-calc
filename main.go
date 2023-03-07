package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(INSTRUCTIONS)
	reader := bufio.NewReaderSize(os.Stdin, 1)
	getByte(reader)
	survey := New(ADULT)
	for q := survey.Next(); q != ""; q = survey.Next() {
		fmt.Println(q)
		for {
			if survey.Add(getByte(reader)) {
				break
			}
		}
	}
	survey.PrintResult()
}

func getByte(reader *bufio.Reader) rune {
	in, err := reader.ReadByte()
	for err != nil {
		_, err = reader.ReadByte()
	}
	fmt.Printf("DEBUG: you pressed:'%c'\n", in)
	return rune(in)
}
