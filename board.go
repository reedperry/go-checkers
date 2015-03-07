package main

import (
	"fmt"
	"math"
)

const SIZE int8 = 8

type Board [SIZE][SIZE]Piece

func (board *Board) set(row, col int8, piece Piece) {
	board[row][col] = piece
}

func (board *Board) get(row, col int8) Piece {
	return board[row][col]
}

func (board *Board) clearSquare(row, col int8) {
	board.set(row, col, EMPTY)
}

func (board *Board) Print() {
	fmt.Println()
	printColNumbers()
	var row int8
	for row = 0; row < SIZE; row++ {
		fmt.Printf("%v ", row)
		printRow(board[row])
	}
}

func printRow(row [8]Piece) {
	var col int8
	fmt.Printf("|")
	for col = 0; col < SIZE; col++ {
		fmt.Printf("%s|", symbol(row[col]))
	}
	fmt.Println()
}

func symbol(piece Piece) string {
	switch piece {
	case EMPTY:
		return " "
	case BLACK_MAN:
		return "b"
	case BLACK_KING:
		return "B"
	case RED_MAN:
		return "r"
	case RED_KING:
		return "R"
	default:
		return "?"
	}
}

func printColNumbers() {
	fmt.Print("  ")
	var col int8 = 0
	for col = 0; col < SIZE; col++ {
		fmt.Printf(" %d", col)
	}
	fmt.Println()
}

func (board *Board) PieceAtSquare(square *Square) Piece {
	return board[square.row][square.col]
}

func (board *Board) ValidSquare(square *Square) bool {
	return square.row >= 0 && square.row < SIZE && square.col >= 0 && square.col < SIZE
}

func (board *Board) PlayableSquare(square *Square) bool {
	return board.PlayableLocation(square.row, square.col)
}

func (board *Board) PlayableLocation(row, col int8) bool {
	return board.ValidSquare(&Square{row, col}) && (row+col)%2 == 1
}

func (board *Board) KingMove(start *Square, playerColor Color) bool {
	return (playerColor == RED && board.PieceAtSquare(start) == RED_KING) ||
		(playerColor == BLACK && board.PieceAtSquare(start) == BLACK_KING)
}

func (board *Board) MovePiece(move *Move) {
	board.set(move.finish.row, move.finish.col, board.get(move.start.row, move.start.col))
	board.clearSquare(move.start.row, move.start.col)
}

func isKing(piece Piece) bool {
	return math.Abs(float64(piece)) > 1
}

func startPieceForRow(row int8) Piece {
	rowsToFill := (SIZE / 2) - 1

	if row < rowsToFill {
		if BLACK_FORWARD > 0 {
			return BLACK_MAN
		} else {
			return RED_MAN
		}
	} else if row == rowsToFill || row == rowsToFill+1 {
		return EMPTY
	} else {
		if BLACK_FORWARD < 0 {
			return BLACK_MAN
		} else {
			return RED_MAN
		}
	}
}
