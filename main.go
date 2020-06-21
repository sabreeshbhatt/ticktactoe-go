package main

import (
	"fmt"
	"strings"
)

func DisplayBoard(board [][]string) { //DisplayBoard...Function to display TicTacToe Board
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func main() {
	// creating a tictactoe board
	board := [][]string{
		[]string{" ", "|", " ", "|", " "},
		[]string{"-", "-", "-", "-", "-"},
		[]string{" ", "|", " ", "|", " "},
		[]string{"-", "-", "-", "-", "-"},
		[]string{" ", "|", " ", "|", " "},
	}
	DisplayBoard(board)
}
