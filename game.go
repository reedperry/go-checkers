package main

import (
	"errors"
	"fmt"
	"math"
)

type Color int8
type Piece int8
type Direction int8

const (
	BLACK_KING Piece = -2
	BLACK_MAN  Piece = -1
	EMPTY      Piece = 0
	RED_MAN    Piece = 1
	RED_KING   Piece = 2

	RED   Color = 1
	BLACK Color = -1
	NONE  Color = 0

	RED_FORWARD   Direction = -1
	BLACK_FORWARD Direction = 1
)

type Game struct {
	board       *Board
	redPlayer   Player
	blackPlayer Player
	moves       []Move
}

type Player struct {
	color Color
}

func (p Piece) Color() Color {
	if p < 0 {
		return BLACK
	} else if p > 0 {
		return RED
	} else {
		return NONE
	}
}

func (player *Player) PlayDirection() Direction {
	if player.color == RED {
		return RED_FORWARD
	} else if player.color == BLACK {
		return BLACK_FORWARD
	} else {
		// TODO Fail here?
		return -1
	}
}

func PlayDirectionOfPiece(piece Piece) Direction {
	if piece == BLACK_MAN || piece == BLACK_KING {
		return BLACK_FORWARD
	} else if piece == RED_MAN || piece == RED_KING {
		return RED_FORWARD
	}
	return 0
}

func (game *Game) NewGame() {
	fmt.Println("Starting new game...")

	var row, col int8
	for row = 0; row < SIZE; row++ {
		for col = 0; col < SIZE; col++ {
			if row == 3 || row == 4 || (row+col)%2 == 0 {
				game.board[row][col] = EMPTY
			} else {
				game.board[row][col] = startPieceForRow(row)
			}
		}
	}
}

func (game *Game) DoMove(start *Square, end *Square, player *Player) error {
	move := &Move{*start, *end, *player}
	moveType, kingMove := game.board.MoveType(move)
	if moveType != ILLEGAL {
		game.board.MovePiece(move)
		if moveType == JUMP {
			game.board.CapturePiece(move)
		}
		if kingMove {
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

func (board *Board) CapturePiece(move *Move) bool {
	captureRow := (move.finish.row + move.start.row) / 2
	captureCol := (move.finish.col + move.start.col) / 2
	captureSquare := &Square{captureRow, captureCol}

	if !areOpponents(move.player.color, board.PieceAtSquare(captureSquare).Color()) {
		return false
	}

	board[captureRow][captureCol] = EMPTY
	return true
}

func (board *Board) MakeKing(square *Square) bool {
	currentPiece := board.PieceAtSquare(square)
	if isKing(currentPiece) {
		return false
	} else if currentPiece == EMPTY {
		return false
	} else {
		// TODO Refactor
		board[square.row][square.col] = currentPiece * 2
		return true
	}
}

func (board *Board) AvailableMoves(start *Square, playerColor Color) []*Square {
	var playDirection Direction

	if playerColor == RED {
		playDirection = RED_FORWARD
	} else {
		playDirection = BLACK_FORWARD
	}

	options := make([]*Square, 0)
	var destination *Square
	if destination = board.FindMoveInDirection(int8(playDirection), -1, start, playerColor); destination != nil {
		options = append(options, destination)
	}
	if destination = board.FindMoveInDirection(int8(playDirection), 1, start, playerColor); destination != nil {
		options = append(options, destination)
	}

	if board.KingMove(start, playerColor) {
		if destination = board.FindMoveInDirection(-1*int8(playDirection), -1, start, playerColor); destination != nil {
			options = append(options, destination)
		}
		if destination = board.FindMoveInDirection(-1*int8(playDirection), 1, start, playerColor); destination != nil {
			options = append(options, destination)
		}
	}

	return options
}

func (board *Board) FindMoveInDirection(dRow int8, dCol int8, start *Square, playerColor Color) *Square {

	adjacentSquare := &Square{start.row + dRow, start.col + dCol}
	if board.PlayableSquare(adjacentSquare) {
		piece := board.PieceAtSquare(adjacentSquare)
		if piece == EMPTY {
			return adjacentSquare
		} else if areOpponents(playerColor, piece.Color()) {
			jumpSquare := &Square{start.row + 2*dRow, start.col + 2*dCol}
			if board.PlayableSquare(jumpSquare) {
				piece = board.PieceAtSquare(jumpSquare)
				if piece == EMPTY {
					return jumpSquare
				}
			}
		}
	}
	return nil
}

func (board *Board) MoveType(move *Move) (MoveType, bool) {
	if !board.PlayableSquare(&move.start) || !board.PlayableSquare(&move.finish) {
		return ILLEGAL, false
	}

	startPiece := board.PieceAtSquare(&move.start)
	endPiece := board.PieceAtSquare(&move.finish)

	if !areTeammates(startPiece.Color(), move.player.color) {
		return ILLEGAL, false
	}
	if endPiece != EMPTY {
		return ILLEGAL, false
	}

	moveSize := math.Abs(float64(move.start.row - move.finish.row))
	if moveSize != math.Abs(float64(move.start.col-move.finish.col)) {
		return ILLEGAL, false
	}

	var moveType MoveType = ILLEGAL
	if moveSize == 1 {
		moveType = SINGLE
	} else if moveSize == 2 {
		moveType = JUMP
	}

	kingMove := false
	// Check for reaching opponent's back line
	if moveType != ILLEGAL && !isKing(startPiece) {
		playerColor := playerColorOf(startPiece)
		playDirection := PlayDirectionOfPiece(Piece(playerColor))

		if playDirection < 0 && move.finish.row == 0 {
			kingMove = true
		} else if playDirection > 0 && move.finish.row == SIZE-1 {
			kingMove = true
		}
	}

	return moveType, kingMove
}

func (game *Game) ValidMove(move *Move) bool {

	playerColor := move.player.color
	kingMove := game.board.KingMove(&move.start, playerColor)
	board := game.board

	if !board.ValidSquare(&move.start) || !board.ValidSquare(&move.finish) {
		return false
	}

	if board.PieceAtSquare(&move.start) == EMPTY {
		return false
	} else if board.PieceAtSquare(&move.finish) != EMPTY {
		return false
	}

	if move.Direction() != move.player.PlayDirection() && !kingMove {
		return false
	}

	availableMoves := board.AvailableMoves(&move.start, playerColor)
	for _, option := range availableMoves {
		if &move.finish == option {
			return true
		}
	}

	return false
}

func playerColorOf(piece Piece) Color {
	if piece == BLACK_MAN || piece == BLACK_KING {
		return BLACK
	} else if piece == RED_MAN || piece == RED_KING {
		return RED
	} else {
		return NONE
	}
}

func areOpponents(color1, color2 Color) bool {
	if color1 > NONE {
		return color2 < NONE
	} else if color1 < NONE {
		return color2 > NONE
	}
	return false
}

func areTeammates(color1, color2 Color) bool {
	if color1 > NONE {
		return color2 > NONE
	} else if color1 < NONE {
		return color2 < NONE
	}
	return false
}

func (game *Game) Print() {
	game.board.Print()
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
