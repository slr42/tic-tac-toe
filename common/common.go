package common

import (
	"fmt"
)

type Board struct {
	Field [3][3]string
}

type Player struct {
	Name string
	Mark string
	result Result
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
	Attacks []*Position
	Defends []*Position
	Weights [3][3]int
}

func Turn(board *Board, player *Player, position *Position) {
	markPosition(board, player, position)
}

func ResetPlayerResult(player *Player) {
	player.result = Result{}
}

func CheckWinCondition(board *Board, player *Player) bool {
	isWin := false
	boardSize := len(board.Field)
	for i := 0; i < len(player.result.XCount); i++ {
		if player.result.XCount[i] >= boardSize {
			isWin = true
		}
	}
	for j := 0; j < len(player.result.YCount); j++ {
		if player.result.YCount[j] >= boardSize {
			isWin = true
		}
	}
	if player.result.Diagonal1Count >= boardSize || player.result.Diagonal2Count >= boardSize {
		isWin = true
	}
	if isWin {
		fmt.Println("Player " + player.Mark + " isWin!")
	}
	return isWin
}

func markPosition(board *Board, player *Player, position *Position) {
	board.Field[position.X][position.Y] = player.Mark
	player.result.XCount[position.X]++
	player.result.YCount[position.Y]++
	if position.X == position.Y {
		player.result.Diagonal1Count++
	}
	if position.X + position.Y == len(board.Field)-1 {
		player.result.Diagonal2Count++
	}
}