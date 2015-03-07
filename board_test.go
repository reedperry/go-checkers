package main

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	board := new(Board)
	board.NewGame()

	if board[0][0] != EMPTY {
		t.Errorf("Square 0,0 should be empty.")
	}

	if board[0][1] == EMPTY {
		t.Errorf("Square 0,1 should not be empty.")
	}

	if board[3][4] != EMPTY {
		t.Errorf("Square 3,4 should be empty.")
	}

	if board[SIZE-1][SIZE-1] != EMPTY {
		t.Errorf("Square %d, %d should be empty.", SIZE-1, SIZE-1)
	}
}

func TestAreOpponents(t *testing.T) {
	if !areOpponents(BLACK, RED) {
		t.Errorf("BLACK and RED should be opponents.")
	}

	if areOpponents(RED, RED) {
		t.Errorf("RED and RED should not be opponents.")
	}

	if !areOpponents(BLACK_KING, RED) {
		t.Errorf("BLACK_KING and RED should be opponents.")
	}

	if !areOpponents(BLACK_KING, RED_KING) {
		t.Errorf("BLACK_KING and RED_KING should be opponents.")
	}

	if areOpponents(BLACK, EMPTY) {
		t.Errorf("BLACK and EMPTY should not be opponents.")
	}

	if areOpponents(RED_KING, EMPTY) {
		t.Errorf("RED_KING and EMPTY should not be opponents.")
	}

	if areOpponents(EMPTY, EMPTY) {
		t.Errorf("EMPTY and EMPTY should not be opponents.")
	}
}

func TestAreTeammates(t *testing.T) {
	if areTeammates(BLACK, RED) {
		t.Errorf("BLACK and RED should not be teammates.")
	}

	if !areTeammates(RED, RED) {
		t.Errorf("RED and RED should be teammates.")
	}

	if areTeammates(BLACK_KING, RED) {
		t.Errorf("BLACK_KING and RED should not be teammates.")
	}

	if areTeammates(BLACK_KING, RED_KING) {
		t.Errorf("BLACK_KING and RED_KING should not be teammates.")
	}

	if areTeammates(BLACK, EMPTY) {
		t.Errorf("BLACK and EMPTY should not be teammates.")
	}

	if areTeammates(RED_KING, EMPTY) {
		t.Errorf("RED_KING and EMPTY should not be teammates.")
	}

	if areTeammates(EMPTY, EMPTY) {
		t.Errorf("EMPTY and EMPTY should not be teammates.")
	}
}

func TestStartColorForRow(t *testing.T) {
	if startColorForRow(0) != BLACK {
		t.Errorf("Start color for row 0 should be BLACK.")
	}

	if startColorForRow((SIZE/2)-2) != BLACK {
		t.Errorf(fmt.Sprintf("Start color for row %d should be BLACK.", (SIZE/2)-2))
	}

	if startColorForRow((SIZE/2)-1) != EMPTY {
		t.Errorf(fmt.Sprintf("Start color for row %d should be EMPTY.", (SIZE/2)-1))
	}

	if startColorForRow(SIZE/2) != EMPTY {
		t.Errorf(fmt.Sprintf("Start color for row %d should be EMPTY.", (SIZE / 2)))
	}

	if startColorForRow(SIZE) != RED {
		t.Errorf(fmt.Sprintf("Start color for row %d should be RED.", SIZE))
	}
}

func TestIsKing(t *testing.T) {
	if isKing(BLACK) {
		t.Errorf("BLACK is not a king.")
	}

	if isKing(RED) {
		t.Errorf("RED is not a king.")
	}

	if isKing(EMPTY) {
		t.Errorf("EMPTY is not a king.")
	}

	if !isKing(BLACK_KING) {
		t.Errorf("BLACK_KING is a king.")
	}

	if !isKing(RED_KING) {
		t.Errorf("RED_KING is a king.")
	}
}

func TestMakeKing(t *testing.T) {
	board := new(Board)
	board[5][6] = BLACK
	board[6][5] = RED

	game := new(Game)
	game.board = board
	blackPlayer := &Player{BLACK}

	game.Print()
	err := game.DoMove(&Square{5, 6}, &Square{7, 4}, blackPlayer)
	game.Print()

	if err != nil {
		t.Errorf("Error during valid capture move.")
	}
	if board.StatusOfSquare(&Square{7, 4}) != BLACK_KING {
		t.Errorf("Square [7, 4] should contain a black king.")
	}
}
