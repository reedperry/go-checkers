package main

import (
	"fmt"
	"math"
)

const EMPTY = 0
const BLACK = 1
const RED = 2
const BLACK_KING = 3
const RED_KING = 4

const SIZE = 8

type Board struct {
	state [SIZE][SIZE]int8
}

func (board *Board) NewGame() {
	fmt.Println("Starting new game...")

	var row, col int8
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
	var row int8
	for row = 0; row < SIZE; row++ {
		printRow(board.state[row])
	}
	fmt.Println()
}

func printRow(row [8]int8) {
	var col int8
	fmt.Printf("|")
	for col = 0; col < SIZE; col++ {
		fmt.Printf("%d|", row[col])
	}
	fmt.Println()
}

func (board *Board) statusOfSquare(square *Square) int8 {
	return board.state[square.row][square.col]
}

func (board *Board) validSquare(square *Square) bool {
	return square.row >= 0 && square.row < SIZE && square.col >= 0 && square.col < SIZE
}

func (board *Board) validMove(move *Move) bool {

	playerColor := move.player.color
	kingMove := board.isKingMove(&move.start, playerColor)

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

	availableMoves := board.getAvailableMoves(&move.start, playerColor)
	for _, option := range availableMoves {
		if &move.finish == option {
			return true
		}
	}

	return false
}

func (board *Board) isPlayableSquare(square *Square) bool {
	return board.isPlayableLocation(square.row, square.col)
}

func (board *Board) isPlayableLocation(row, col int8) bool {
	return board.validSquare(&Square{row, col}) && (row+col)%2 == 1
}

func (board *Board) isKingMove(start *Square, playerColor int8) bool {
	return (playerColor == RED && board.statusOfSquare(start) == RED_KING) ||
		(playerColor == BLACK && board.statusOfSquare(start) == BLACK_KING)
}

func (board *Board) getAvailableMoves(start *Square, playerColor int8) []*Square {
	var playDirection Direction

	if playerColor == RED {
		playDirection = RED_FORWARD
	} else {
		playDirection = BLACK_FORWARD
	}

	options := make([]*Square, 0)
	var destination *Square
	if destination = board.searchForMoveInDirection(int8(playDirection), -1, start, playerColor); destination != nil {
		options = append(options, destination)
	}
	if destination = board.searchForMoveInDirection(int8(playDirection), 1, start, playerColor); destination != nil {
		options = append(options, destination)
	}

	if board.isKingMove(start, playerColor) {
		if destination = board.searchForMoveInDirection(-1*int8(playDirection), -1, start, playerColor); destination != nil {
			options = append(options, destination)
		}
		if destination = board.searchForMoveInDirection(-1*int8(playDirection), 1, start, playerColor); destination != nil {
			options = append(options, destination)
		}
	}

	return options
}

// TODO Better strategy for determining opponent pieces, including kings...
func (board *Board) searchForMoveInDirection(dRow int8, dCol int8, start *Square, playerColor int8) *Square {

	adjacentSquare := &Square{start.row + dRow, start.col + dCol}
	if board.isPlayableSquare(adjacentSquare) {
		status := board.statusOfSquare(adjacentSquare)
		if status == EMPTY {
			return adjacentSquare
		} else if status == opponentOf(playerColor) {
			jumpSquare := &Square{start.row + 2*dRow, start.col + 2*dCol}
			if board.isPlayableSquare(adjacentSquare) {
				status = board.statusOfSquare(jumpSquare)
				if status == EMPTY {
					return jumpSquare
				}
			}
		}
	}
	return nil
}

func (board *Board) getMoveType(move *Move) MoveType {
	if !board.isPlayableSquare(&move.start) || !board.isPlayableSquare(&move.finish) {
		return ILLEGAL
	}

	startStatus := board.statusOfSquare(&move.start)
	endStatus := board.statusOfSquare(&move.finish)

	// Again, doesn't handle kings yet...
	if startStatus != move.player.color {
		return ILLEGAL
	}
	if endStatus != EMPTY {
		return ILLEGAL
	}

	moveSize := math.Abs(float64(move.start.row - move.finish.row))
	if moveSize != math.Abs(float64(move.start.col-move.finish.col)) {
		return ILLEGAL
	}

	var moveType MoveType = ILLEGAL
	if moveSize == 1 {
		moveType = SIMPLE
	} else if moveSize == 2 {
		moveType = JUMP
	}

	// TODO Check for reaching opponent's back line
	if moveType != ILLEGAL {

	}

	return moveType
}

// FIXME Doesn't handle kings...
func opponentOf(color int8) int8 {
	if color == RED {
		return BLACK
	} else if color == BLACK {
		return RED
	}
	return -1
}

func colorForRow(row int8) int8 {
	if row < 3 {
		return BLACK
	} else {
		return RED
	}
}
