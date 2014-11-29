package main

import (
	"fmt"
	"math"
)

// Make it easy to tell what color occupies a space
const BLACK_KING int8 = -2
const BLACK int8 = -1
const EMPTY int8 = 0
const RED int8 = 1
const RED_KING int8 = 2

const RED_FORWARD int8 = -1
const BLACK_FORWARD int8 = 1

const SIZE int8 = 8

type Board struct {
	state [SIZE][SIZE]int8
}

func (player *Player) PlayDirection() int8 {
	if player.color == RED {
		return RED_FORWARD
	} else if player.color == BLACK {
		return BLACK_FORWARD
	} else {
		return -1
	}
}

func PlayDirectionOfColor(color int8) int8 {
	if color == BLACK || color == BLACK_KING {
		return BLACK_FORWARD
	} else if color == RED || color == RED_KING {
		return RED_FORWARD
	}
	return 0
}

func (board *Board) NewGame() {
	fmt.Println("Starting new game...")

	var row, col int8
	for row = 0; row < SIZE; row++ {
		for col = 0; col < SIZE; col++ {
			if row == 3 || row == 4 || (row+col)%2 == 0 {
				board.state[row][col] = EMPTY
			} else {
				board.state[row][col] = startColorForRow(row)
			}
		}
	}
}

func (board *Board) PrintGame() {
	fmt.Println()
	var row int8
	for row = 0; row < SIZE; row++ {
		printRow(board.state[row])
	}
}

func printRow(row [8]int8) {
	var col int8
	fmt.Printf("|")
	for col = 0; col < SIZE; col++ {
		fmt.Printf("%s|", symbolForStatus(row[col]))
	}
	fmt.Println()
}

func symbolForStatus(status int8) string {
	switch status {
	case EMPTY:
		return " "
	case BLACK:
		return "b"
	case BLACK_KING:
		return "B"
	case RED:
		return "r"
	case RED_KING:
		return "R"
	default:
		return "?"
	}
}

func (board *Board) StatusOfSquare(square *Square) int8 {
	return board.state[square.row][square.col]
}

func (board *Board) ValidSquare(square *Square) bool {
	return square.row >= 0 && square.row < SIZE && square.col >= 0 && square.col < SIZE
}

func (board *Board) ValidMove(move *Move) bool {

	playerColor := move.player.color
	kingMove := board.KingMove(&move.start, playerColor)

	if !board.ValidSquare(&move.start) || !board.ValidSquare(&move.finish) {
		return false
	}

	if board.StatusOfSquare(&move.start) == EMPTY {
		return false
	} else if board.StatusOfSquare(&move.finish) != EMPTY {
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

func (board *Board) PlayableSquare(square *Square) bool {
	return board.PlayableLocation(square.row, square.col)
}

func (board *Board) PlayableLocation(row, col int8) bool {
	return board.ValidSquare(&Square{row, col}) && (row+col)%2 == 1
}

func (board *Board) KingMove(start *Square, playerColor int8) bool {
	return (playerColor == RED && board.StatusOfSquare(start) == RED_KING) ||
		(playerColor == BLACK && board.StatusOfSquare(start) == BLACK_KING)
}

func (board *Board) MovePiece(move *Move) {
	board.state[move.finish.row][move.finish.col] = board.state[move.start.row][move.start.col]
	board.state[move.start.row][move.start.col] = EMPTY
}

func (board *Board) AvailableMoves(start *Square, playerColor int8) []*Square {
	var playDirection int8

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

func (board *Board) FindMoveInDirection(dRow int8, dCol int8, start *Square, playerColor int8) *Square {

	adjacentSquare := &Square{start.row + dRow, start.col + dCol}
	if board.PlayableSquare(adjacentSquare) {
		status := board.StatusOfSquare(adjacentSquare)
		if status == EMPTY {
			return adjacentSquare
		} else if areOpponents(playerColor, status) {
			jumpSquare := &Square{start.row + 2*dRow, start.col + 2*dCol}
			if board.PlayableSquare(adjacentSquare) {
				status = board.StatusOfSquare(jumpSquare)
				if status == EMPTY {
					return jumpSquare
				}
			}
		}
	}
	return nil
}

func (board *Board) MoveType(move *Move) MoveType {
	if !board.PlayableSquare(&move.start) || !board.PlayableSquare(&move.finish) {
		return ILLEGAL
	}

	startStatus := board.StatusOfSquare(&move.start)
	endStatus := board.StatusOfSquare(&move.finish)

	if !areTeammates(startStatus, move.player.color) {
		return ILLEGAL
	}
	if endStatus != EMPTY {
		return ILLEGAL
	}

	moveSize := math.Abs(float64(move.start.row - move.finish.row))
	if moveSize != math.Abs(float64(move.start.col-move.finish.col)) {
		return ILLEGAL
	}

	var moveType MoveType = ILLEGAL
	if moveSize == 1 {
		moveType = SIMPLE
	} else if moveSize == 2 {
		moveType = JUMP
	}

	// Check for reaching opponent's back line
	if moveType != ILLEGAL && !isKing(startStatus) {
		playerColor := playerColorOf(startStatus)
		playDirection := PlayDirectionOfColor(playerColor)

		if playDirection < 0 && move.finish.row == 0 {
			moveType = KING
		} else if playDirection > 0 && move.finish.row == SIZE-1 {
			moveType = KING
		}
	}

	return moveType
}

func (board *Board) CapturePiece(move *Move) bool {
	captureRow := (move.finish.row + move.start.row) / 2
	captureCol := (move.finish.col + move.start.col) / 2
	captureSquare := &Square{captureRow, captureCol}

	if !areOpponents(move.player.color, board.StatusOfSquare(captureSquare)) {
		return false
	}

	board.state[captureRow][captureCol] = EMPTY
	return true
}

func (board *Board) MakeKing(square *Square) bool {
	currentStatus := board.StatusOfSquare(square)
	if isKing(currentStatus) {
		return false
	} else if currentStatus == EMPTY {
		return false
	} else {
		board.state[square.row][square.col] = currentStatus * 2
		return true
	}
}

func isKing(status int8) bool {
	return math.Abs(float64(status)) > 1
}

func playerColorOf(status int8) int8 {
	if status == BLACK || status == BLACK_KING {
		return BLACK
	} else if status == RED || status == RED_KING {
		return RED
	} else {
		return EMPTY
	}
}

func areOpponents(color1, color2 int8) bool {
	if color1 > 0 {
		return color2 < 0
	} else if color1 < 0 {
		return color2 > 0
	}
	return false
}

func areTeammates(color1, color2 int8) bool {
	if color1 > 0 {
		return color2 > 0
	} else if color1 < 0 {
		return color2 < 0
	}
	return false
}

func startColorForRow(row int8) int8 {
	rowsToFill := (SIZE / 2) - 1

	if row < rowsToFill {
		if BLACK_FORWARD > 0 {
			return BLACK
		} else {
			return RED
		}
	} else if row == rowsToFill || row == rowsToFill+1 {
		return EMPTY
	} else {
		if BLACK_FORWARD < 0 {
			return BLACK
		} else {
			return RED
		}
	}
}
