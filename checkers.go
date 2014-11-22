package main

func main() {
	board := new(Board)
	board.NewGame()

	game := new(Game)
	game.board = *board

	blackPlayer := &Player{BLACK}
	redPlayer := &Player{RED}

	game.Print()
	game.DoMove(&Square{2, 1}, &Square{3, 2}, blackPlayer)
	game.DoMove(&Square{5, 0}, &Square{4, 1}, redPlayer)
	game.DoMove(&Square{2, 3}, &Square{3, 4}, blackPlayer)
	game.DoMove(&Square{5, 6}, &Square{4, 5}, redPlayer)
	game.Print()

	game.PrintMoves()
}
