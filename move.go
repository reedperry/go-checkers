package main

const SIMPLE = 7
const JUMP = 8
const ILLEGAL = -1

type MoveType int8

type Move struct {
	start  Square
	finish Square
	player Player
}

type Square struct {
	row, col int8
}

func (move *Move) Direction() int8 {
	if move.start.row < move.finish.row {
		return RED_FORWARD
	} else {
		return BLACK_FORWARD
	}
}
