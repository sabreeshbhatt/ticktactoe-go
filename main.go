package main

import (
	"fmt"
	"strings"
)

func main() {
	// creating a tictactoe board
	board := [][]string{
		[]string{" ", "|", " ", "|", " "},
		[]string{"-", "-", "-", "-", "-"},
		[]string{" ", "|", " ", "|", " "},
		[]string{"-", "-", "-", "-", "-"},
		[]string{" ", "|", " ", "|", " "},
	}
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
