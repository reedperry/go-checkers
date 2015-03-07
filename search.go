package main

import (
	"errors"
	"fmt"
	"math"
)

func SelectMove(board *Board, color Color) ([]*Move, error) {
	var moves []*Move
	myGame := new(Game)
	myGame.board = board

	return moves, nil
}

func FindOneMove(board *Board, color Color) (*Move, int8, error) {

	myPieces := findMyPieces(board, color)

	var moves = make(map[*Square][]*Square)
	for _, piece := range myPieces {
		moves[piece] = findMovesForPiece(piece, board, color)
	}

	var score, max int8
	var best *Move
	player := &Player{color}

	for start, destinations := range moves {
		for _, dest := range destinations {
			//fmt.Printf("Scoring move from %v to %v...", start, dest)
			score = scoreMove(start, dest, color, board)
			//fmt.Printf("%v\n", score)
			if score > max {
				max = score
				best = &Move{*start, *dest, *player}
			}
		}
	}

	if max > 0 {
		return best, max, nil
	} else {
		return nil, -1, errors.New(fmt.Sprintf("No move found for player %v", color))
	}
}

func findMyPieces(board *Board, color Color) []*Square {
	var row, col int8
	var myPieces []*Square

	for row = 0; row < SIZE; row++ {
		for col = 0; col < SIZE; col++ {
			if areTeammates(color, board[row][col].Color()) {
				myPieces = append(myPieces, &Square{row, col})
			}
		}
	}

	return myPieces
}

func findMovesForPiece(piece *Square, board *Board, color Color) []*Square {
	moves := board.AvailableMoves(piece, color)
	return moves
}

func scoreMove(start *Square, end *Square, color Color, board *Board) int8 {
	player := &Player{color}
	move := &Move{*start, *end, *player}
	moveType, kingMove := board.MoveType(move)

	if moveType == ILLEGAL {
		return math.MinInt8
	}

	if moveType == SINGLE {
		if kingMove {
			return 5
		} else {
			return 1
		}
	}

	if moveType == JUMP {
		if kingMove {
			return 8
		} else {
			return 3
		}
	}

	return math.MinInt8
}
