package main

import (
	"fmt"
)

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

func (b *Board) CheckWinner() bool {
	
	for i := 0; i < 3; i++ {
		
		if b.Board[i][0] == b.Turn && b.Board[i][1] == b.Turn && b.Board[i][2] == b.Turn {
			b.Winner = b.Turn
			return true
		}
		
		if b.Board[0][i] == b.Turn && b.Board[1][i] == b.Turn && b.Board[2][i] == b.Turn {
			b.Winner = b.Turn
			return true
		}
	}

	if b.Board[0][0] == b.Turn && b.Board[1][1] == b.Turn && b.Board[2][2] == b.Turn {
		b.Winner = b.Turn
		return true
	}
	if b.Board[0][2] == b.Turn && b.Board[1][1] == b.Turn && b.Board[2][0] == b.Turn {
		b.Winner = b.Turn
		return true
	}

	return false
}

func (b *Board) CheckDraw() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.Board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}

func (b *Board) SwitchTurn() {
	if b.Turn == "X" {
		b.Turn = "O"
	} else {
		b.Turn = "X"
	}
}

func (b *Board) GetInput() {
	var row, col int
	for {
		fmt.Printf("Player %s, enter your move (row and column): ", b.Turn)
		fmt.Scan(&row, &col)

		if row >= 0 && row < 3 && col >= 0 && col < 3 && b.Board[row][col] == "_" {
			b.Board[row][col] = b.Turn
			break
		} else {
			fmt.Println("Invalid move. Try again.")
		}
	}
}

func PlayGame() {
	game := StartGame()
	for game.GameRunning {
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
