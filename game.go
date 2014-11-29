package main

import (
	"errors"
	"fmt"
)

const SIMPLE = 7
const JUMP = 8
const KING = 9
const ILLEGAL = -1

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

func (game *Game) DoMove(start *Square, end *Square, player *Player) error {
	move := &Move{*start, *end, *player}
	moveType := game.board.MoveType(move)
	if moveType != ILLEGAL {
		game.board.MovePiece(move)
		if moveType == JUMP {
			game.board.CapturePiece(move)
		} else if moveType == KING {
			game.board.MakeKing(&move.finish)
		}
		game.moves = append(game.moves, *move)
		return nil
	}

	color := "Black"
	if player.color == RED {
		color = "Red"
	}
	return errors.New(fmt.Sprintf("Illegal move attempted by %v player: %v -> %v", color, *start, *end))
}

func (game *Game) Print() {
	game.board.PrintGame()
}

func (game *Game) PrintMoves() {
	fmt.Println()
	for i, m := range game.moves {
		fmt.Print(i+1, ": ")
		if m.player.color == BLACK {
			fmt.Print("Black ")
		} else {
			fmt.Print("Red ")
		}
		fmt.Println(m.start, "->", m.finish)
	}
}
