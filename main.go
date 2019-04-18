package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	. "tic_tac_toe/utility"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputRegexp := regexp.MustCompile("^[0-2],[0-2]$")
	fmt.Println("*** Tic-tac-toe game ***")
	fmt.Println("`restart` -- to restart game")
	fmt.Println("`exit` -- exit game")

	player := Player{Name: "Player", Mark: "X"}
	ticTacBot := Player{Name: "TicTacBot", Mark: "0"}

	board := initialize(&player, &ticTacBot)

	PrintBoard(board)
	fmt.Println("Enter position in X,Y-format (ex. 2,3)")
	var x, y int
	for scanner.Scan() && scanner.Text() != "exit" {
		if scanner.Text() == "restart" {
			board = initialize(&player, &ticTacBot)
			PrintBoard(board)
			continue
		}
		xyStr := scanner.Text()
		if inputIsValid := inputRegexp.Match([]byte(xyStr)); !inputIsValid {
			fmt.Println("Invalid board position! Only two digits are required")
			continue
		}
		fmt.Println("Enter position in X,Y-format (ex. 2,3)")

		position := Position{}
		_, _ = fmt.Sscanf(xyStr, "%d,%d", &position.X, &position.Y)
		if len(board.Field[x][y]) > 0 {
			fmt.Println("Invalid board position! Position is already chosen")
			continue
		}

		Turn(&board, &player, &position)
		if CheckWinCondition(&board, &player) {
			break
		}

		ComputerTurn(&board, &ticTacBot, &player)
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
