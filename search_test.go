package main

import (
	"fmt"
	"testing"
)

func TestFindMove(t *testing.T) {
	board := new(Board)
	board.NewGame()

	game := new(Game)
	game.board = board

	blackPlayer := &Player{BLACK}
	redPlayer := &Player{RED}

	game.DoMove(&Square{2, 1}, &Square{3, 2}, blackPlayer)
	game.DoMove(&Square{5, 0}, &Square{4, 1}, redPlayer)
	game.DoMove(&Square{2, 3}, &Square{3, 4}, blackPlayer)
	game.DoMove(&Square{5, 6}, &Square{4, 5}, redPlayer)
	game.DoMove(&Square{3, 2}, &Square{5, 0}, blackPlayer)

	blackMove, err := FindMove(game, BLACK)
	if err == nil {
		fmt.Printf("Best move for black: %v\n", blackMove)
	} else {
		fmt.Println("No moves for black.")
	}
	redMove, err := FindMove(game, RED)
	if err == nil {
		fmt.Printf("Best move for red: %v\n", redMove)
	} else {
		fmt.Println("No moves for red.")
	}

}
