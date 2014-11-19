package main

const RED_FORWARD = -1
const BLACK_FORWARD = 1

const SIMPLE = 7
const JUMP = 8
const KING = 9
const ILLEGAL = -1

type Direction int8
type MoveType int8

type Game struct {
	board       Board
	redPlayer   Player
	blackPlayer Player
	moves       []Move
}

type Player struct {
	color int8
}

func (player *Player) getPlayDirection() Direction {
	if player.color == RED {
		return RED_FORWARD
	} else if player.color == BLACK {
		return BLACK_FORWARD
	} else {
		return -1
	}
}