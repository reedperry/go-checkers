package main

import "fmt"

func main() {
	board := new(Board)

	board.NewGame()
	board.PrintGame()
	moves := board.getAvailableMoves(&Square{5, 0}, RED)
	for _, m := range moves {
		fmt.Println("Available Move", *m)
	}
}
