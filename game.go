package main

import (
	"errors"
)

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

func (game *Game) doMove(start *Square, end *Square, player *Player) error {
	move := &Move{*start, *end, *player}
	moveType := game.board.getMoveType(move)
	if moveType != ILLEGAL {
		game.board.state[move.finish.row][move.finish.col] = game.board.state[move.start.row][move.start.col]
		game.board.state[move.start.row][move.start.col] = EMPTY
		return nil
	}
	return errors.New("Illegal move")
}
