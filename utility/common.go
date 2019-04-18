package tic_tac_toe

import (
	"fmt"
)

type Board struct {
	Field [3][3]string
}

type Player struct {
	Name   string
	Mark   string
	Result Result
}

type Position struct {
	X, Y int
}

type Result struct {
	XCount         [3]int
	YCount         [3]int
	Diagonal1Count int
	Diagonal2Count int
}

type AnalyzeReport struct {
	Attacks []Position
	Defends []Position
	Weights [3][3]int
}

func Turn(board *Board, player *Player, position *Position) {
	markPosition(board, player, position)
}

func ResetPlayerResult(player *Player) {
	player.Result = Result{}
}

func CheckWinCondition(board *Board, player *Player) bool {
	isWin := false
	boardSize := len(board.Field)
	for i := 0; i < len(player.Result.XCount); i++ {
		if player.Result.XCount[i] >= boardSize {
			isWin = true
		}
	}
	for j := 0; j < len(player.Result.YCount); j++ {
		if player.Result.YCount[j] >= boardSize {
			isWin = true
		}
	}
	if player.Result.Diagonal1Count >= boardSize || player.Result.Diagonal2Count >= boardSize {
		isWin = true
	}
	if isWin {
		fmt.Println("Player " + player.Mark + " isWin!")
	}
	return isWin
}

func markPosition(board *Board, player *Player, position *Position) {
	board.Field[position.X][position.Y] = player.Mark
	player.Result.XCount[position.X]++
	player.Result.YCount[position.Y]++
	if position.X == position.Y {
		player.Result.Diagonal1Count++
	}
	if position.X+position.Y == len(board.Field)-1 {
		player.Result.Diagonal2Count++
	}
}
