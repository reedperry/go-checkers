package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	board := new(Board)
	board.NewGame()

	if board.state[0][0] != EMPTY {
		t.Errorf("Square 0,0 should be empty.")
	}

	if board.state[3][4] != EMPTY {
		t.Errorf("Square 3,4 should be empty.")
	}

	if board.state[SIZE-1][SIZE-1] != EMPTY {
		t.Errorf("Square %d, %d should be empty.", SIZE-1, SIZE-1)
	}
}
