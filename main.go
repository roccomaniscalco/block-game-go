package main

import "fmt"

var gameBoard [9][9]bool

func init() {
	for i := range gameBoard {
		for j := range gameBoard[i] {
			gameBoard[i][j] = false
		}
	}
}

func main() {
	for i := range gameBoard {
		fmt.Println()
		for j := range gameBoard[i] {
			if gameBoard[i][j]{
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
	}
}
