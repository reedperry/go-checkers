package main

const (
	SINGLE  = 1
	JUMP    = 2
	ILLEGAL = -1
)

type MoveType int8

type Move struct {
	start  Square
	finish Square
	player Player
}

type Square struct {
	row, col int8
}

type Turn struct {
	moves []Move
	score int
}

func (move *Move) Direction() Direction {
	if move.start.row < move.finish.row {
		return RED_FORWARD
	} else {
		return BLACK_FORWARD
	}
}
