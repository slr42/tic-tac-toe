package main

import (
	"bufio"
	"fmt"
	. "github.com/slr42/tic-tac-toe/utility"
	"os"
	"regexp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputRegexp := regexp.MustCompile("^[1-3],[1-3]$")
	fmt.Println("*** Tic-tac-toe game ***")
	fmt.Println("`restart` -- to restart game")
	fmt.Println("`exit` -- exit game")

	player := Player{Name: "Player", Mark: "X"}
	ticTacBot := Player{Name: "TicTacBot", Mark: "0"}

	board := initialize(&player, &ticTacBot)
	var position Position

	PrintBoard(board)
	fmt.Println("Enter position in X,Y-format (ex. 2,3)")

	for scanner.Scan() && scanner.Text() != "exit" && FieldHasFreeCells(board) {
		if scanner.Text() == "restart" {
			board = initialize(&player, &ticTacBot)
			PrintBoard(board)
			continue
		}
		xyStr := scanner.Text()
		if inputIsValid := inputRegexp.Match([]byte(xyStr)); !inputIsValid {
			fmt.Println("Invalid board position! Digits from 1 to 3 are required")
			continue
		}
		fmt.Println("Enter position in X,Y-format (ex. 2,3)")
		position = Position{}
		_, _ = fmt.Sscanf(xyStr, "%d,%d", &position.X, &position.Y)
		position.X = position.X - 1
		position.Y = position.Y - 1
		if len(board.Field[position.X][position.Y]) > 0 {
			fmt.Println("Invalid board position! Position is already chosen")
			continue
		}

		Turn(&board, &player, &position)
		if CheckWinCondition(&board, &player) {
			break
		}

		if !FieldHasFreeCells(board) {
			PrintBoard(board)
			break
		}

		ComputerTurn(&board, &ticTacBot, &player)
		PrintBoard(board)
		if CheckWinCondition(&board, &ticTacBot) {
			break
		}
	}
}

func initialize(player1 *Player, player2 *Player) Board {
	board := Board{}
	ResetPlayerResult(player1)
	ResetPlayerResult(player2)
	return board
}
