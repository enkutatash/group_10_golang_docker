package cmd

import "fmt"

func StartGame() *Board {
	game := &Board{
		GameRunning: true,
		Turn:        "X",
		Winner:      "",
	}

	for i := range game.Board {
		for j := range game.Board[i] {
			game.Board[i][j] = "_"
		}
	}
	return game
}
func (b *Board) PrintBoard() {
	for _, row := range b.Board {
		fmt.Println(row)
	}
	fmt.Println()

}

func (b *Board)CheckWinner() bool{
	panic("unimplemented")
}

func (b *Board)CheckDraw() bool{
	panic("unimplemented")
}

func (b *Board)SwitchTurn() {}

func (b *Board)GetInput() {}

func PlayGame(){
	game := StartGame()
	for game.GameRunning{
		game.PrintBoard()
		game.GetInput()
		if game.CheckWinner() {
			game.GameRunning = false
			game.PrintBoard()
			fmt.Printf("Player %s wins!\n", game.Winner)
			return
		}
		if game.CheckDraw() {
			game.GameRunning = false
			game.PrintBoard()
			fmt.Println("It's a draw!")
			return
		}
		game.SwitchTurn()

	}
}