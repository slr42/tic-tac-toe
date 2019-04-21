package tic_tac_toe

import (
	"fmt"
	"reflect"
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
	WinPositions []Position
	Defends []Position
	PreemptiveDefends []Position
	Attacks []Position
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

func FieldHasFreeCells(board Board) bool {
	for i:=0; i< len(board.Field); i++ {
		for j:=0; j< len(board.Field[i]); j++ {
			if len(board.Field[i][j]) == 0 {
				return true
			}
		}
	}
	fmt.Println("No one wins!")
	return false
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

func positionInPositionList(searchPosition Position, positionList []Position) bool {
	for _, position := range positionList {
		if reflect.DeepEqual(position, searchPosition) {
			return true
		}
	}
	return false
}

func isRowFree(board *Board, x int, orMark string) bool {
	var isRowFree = true
	for _, value := range board.Field[x] {
		if value != "" && value != orMark {
			isRowFree = false
		}
	}
	return isRowFree
}

func isColumnFree(board *Board, y int, orMark string) bool {
	var isColumnFree = true
	for i:=0; i<len(board.Field); i++ {
		if board.Field[i][y] != "" && board.Field[i][y] != orMark {
			isColumnFree = false
		}
	}
	return isColumnFree
}

func isDiagonal1Free(board *Board, orMark string) bool {
	var isDiagonal1Free = true
	for i:=0; i< len(board.Field); i++ {
		if board.Field[i][i] != "" && board.Field[i][i] != orMark {
			isDiagonal1Free = false

		}
	}
	return isDiagonal1Free
}

func isDiagonal2Free(board *Board, orMark string) bool {
	var isDiagonal2Free = true
	for i:=0; i< len(board.Field); i++ {
		j := len(board.Field)-1-i
		if board.Field[i][j] != "" && board.Field[i][j] != orMark {
			isDiagonal2Free = false

		}
	}
	return isDiagonal2Free
}

func clonePlayer(player *Player) *Player {
	clonePlayer := Player{
		Mark: player.Mark,
		Name: player.Name,
		Result: player.Result,
	}
	return &clonePlayer
}

func cloneBoard(board *Board) Board {
	cloneBoard := Board{}
	for i, row := range board.Field {
		for j, value := range row {
			cloneBoard.Field[i][j] = value
		}
	}
	return cloneBoard
}