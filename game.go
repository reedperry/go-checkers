package main

const DIRECTION_UP = 5
const DIRECTION_DOWN = 6
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
	color uint8
}

func (player *Player) getPlayDirection() Direction {
	if player.color == RED {
		return DIRECTION_UP
	} else if player.color == BLACK {
		return DIRECTION_DOWN
	} else {
		return -1
	}
}
