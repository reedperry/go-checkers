package main

import "fmt"

const EMPTY = 0
const BLACK = 1
const RED = 2
const RED_KING = 3
const BLACK_KING = 4

const SIZE = 8

type Board struct {
	state [SIZE][SIZE]uint8
}

func (board *Board) NewGame() {
	fmt.Println("Starting new game...")

	var row, col uint8
	for row = 0; row < SIZE; row++ {
		for col = 0; col < SIZE; col++ {
			if row == 3 || row == 4 || (row+col)%2 == 0 {
				board.state[row][col] = EMPTY
			} else {
				board.state[row][col] = colorForRow(row)
			}
		}
	}
}

func (board *Board) PrintGame() {
	var row uint8
	for row = 0; row < SIZE; row++ {
		printRow(board.state[row])
	}
}

func printRow(row [8]uint8) {
	var col uint8
	fmt.Printf("|")
	for col = 0; col < SIZE; col++ {
		fmt.Printf("%d|", row[col])
	}
	fmt.Println()
}

func (board *Board) statusOfSquare(square *Square) uint8 {
	return board.state[square.row][square.col]
}

func (board *Board) validSquare(square *Square) bool {
	return square.row >= 0 && square.row < SIZE && square.col >= 0 && square.col < SIZE
}

func (board *Board) validMove(move *Move) bool {

	kingMove := (move.player.color == RED && board.statusOfSquare(&move.start) == RED_KING) ||
		(move.player.color == BLACK && board.statusOfSquare(&move.start) == BLACK_KING)

	if !board.validSquare(&move.start) || !board.validSquare(&move.finish) {
		return false
	}

	if board.statusOfSquare(&move.start) == EMPTY {
		return false
	} else if board.statusOfSquare(&move.finish) != EMPTY {
		return false
	}

	if move.getDirection() != move.player.getPlayDirection() && !kingMove {
		return false
	}

	adjacent := board.getAdjacent(&move.start)
	for _, s := range adjacent {
		if board.statusOfSquare(s) != EMPTY {
			return false
		}
	}

	return true
}

func (board *Board) isPlayableSquare(square *Square) bool {
	return board.isPlayableLocation(square.row, square.col)
}

func (board *Board) isPlayableLocation(row, col uint8) bool {
	return (row+col)%2 == 1
}

func (board *Board) getAdjacent(square *Square) []*Square {
	adjacent := make([]*Square, 4)

	var col, row uint8
	for col = square.col - 1; col <= square.col+1; col++ {
		for row = square.row - 1; row <= square.row+1; row++ {
			if col == square.col && row == square.row {
				continue
			} else if !board.isPlayableLocation(row, col) {
				continue
			} else {
				adjacent = append(adjacent, &Square{row, col})
			}
		}
	}

	return adjacent
}

func colorForRow(row uint8) uint8 {
	if row < 3 {
		return BLACK
	} else {
		return RED
	}
}
