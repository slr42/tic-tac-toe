package printer

import (
	"fmt"
	c "tic-tac-toe/common"
)

func PrintBoard(board c.Board) {
	for i := 0; i < len(board.Field); i++ {
		fmt.Print("[")
		for j := 0; j < len(board.Field[i]); j++ {
			symbol := board.Field[i][j]
			if len(symbol) == 0 {
				symbol = "_"
			}
			fmt.Printf(" %s ", symbol)
		}
		fmt.Print("]\n")
	}
}

func PrintReport(report c.AnalyzeReport) {
	fmt.Println("*** Analyze report ***")
	for i := 0; i < len(report.Weights); i++ {
		fmt.Print("[")
		for j := 0; j < len(report.Weights[i]); j++ {
			fmt.Printf(" %d ", report.Weights[i][j])
		}
		fmt.Print("]\n")
	}
}

func PrintResult(result c.Result) {
	fmt.Println(result)
}