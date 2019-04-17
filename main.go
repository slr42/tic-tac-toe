package tic_tac_toe

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"tic-tac-toe/bot"
	c "tic-tac-toe/common"
	"tic-tac-toe/printer"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputRegexp := regexp.MustCompile("^[0-2],[0-2]$")
	fmt.Println("*** Tic-tac-toe game ***")
	fmt.Println("`restart` -- to restart game")
	fmt.Println("`exit` -- exit game")

	player := c.Player{Name: "Player", Mark: "X"}
	ticTacBot := c.Player{Name: "TicTacBot", Mark: "0"}

	board := initialize(&player, &ticTacBot)

	printer.PrintBoard(board)
	fmt.Println("Enter position in X,Y-format (ex. 2,3)")
	var x, y int
	for scanner.Scan() && scanner.Text() != "exit" {
		if scanner.Text() == "restart" {
			board = initialize(&player, &ticTacBot)
			printer.PrintBoard(board)
			continue
		}
		xyStr := scanner.Text()
		if inputIsValid := inputRegexp.Match([]byte(xyStr)); !inputIsValid {
			fmt.Println("Invalid board position! Only two digits are required")
			continue
		}
		fmt.Println("Enter position in X,Y-format (ex. 2,3)")

		position := c.Position{}
		_, _ = fmt.Sscanf(xyStr, "%d,%d", &position.X, &position.Y)
		if len(board.Field[x][y]) > 0 {
			fmt.Println("Invalid board position! Position is already chosen")
			continue
		}

		c.Turn(&board, &player, &position)
		if c.CheckWinCondition(&board, &player) {
			break
		}

		bot.ComputerTurn(&board, &ticTacBot)
		if c.CheckWinCondition(&board, &ticTacBot) {
			break
		}
	}
}

func initialize(player1 *c.Player, player2 *c.Player) c.Board {
	board := c.Board{}
	c.ResetPlayerResult(player1)
	c.ResetPlayerResult(player2)
	return board
}
