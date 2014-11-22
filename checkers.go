package main

func main() {
	board := new(Board)
	board.NewGame()
	board.PrintGame()
	game := new(Game)
	game.board = *board
	blackPlayer := &Player{BLACK}
	game.doMove(&Square{2, 1}, &Square{3, 2}, blackPlayer)
	game.board.PrintGame()
}
