package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print(INSTRUCTIONS)
	var survey *Survey
	switch getByte() {
	case 'a':
		survey = New(ADULT)
	case 'b':
		survey = New(BOY)
	case 'c':
		survey = New(GIRL)
	default:
		fmt.Println("Invalid choice.")
		return
	}

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
