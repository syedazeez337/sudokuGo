package main

import (
	"fmt"
)

const N = 9
const EMPTY = 0

func main() {
	board := [N][N]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	printBoard(board)
}

func printBoard(board [N][N]int) {
	fmt.Println("-------------------------")
	for i:=0; i < N; i++ {
		fmt.Print("| ")
		for j := 0; j < N; j++ {
			if board[i][j] == EMPTY {
				fmt.Print(". ")
			} else {
				fmt.Printf("%d ", board[i][j])
			}
			if (j+1)%3 == 0 {
				fmt.Print("| ")
			}
		}
		fmt.Println()
		if (i+1) % 3 == 0 {
			fmt.Println("-------------------------")
		}
	}
}