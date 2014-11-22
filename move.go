package main

type Move struct {
	start  Square
	finish Square
	player Player
}

type Square struct {
	row, col int8
}

func (move *Move) Direction() Direction {
	if move.start.row < move.finish.row {
		return RED_FORWARD
	} else {
		return BLACK_FORWARD
	}
}
