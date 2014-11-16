package main

type Move struct {
	start  Square
	finish Square
	player Player
}

type Square struct {
	row, col uint8
}

func (move *Move) getDirection() Direction {
	if move.start.row < move.finish.row {
		return DIRECTION_UP
	} else {
		return DIRECTION_DOWN
	}
}
