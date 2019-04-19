package tic_tac_toe

import (
	"fmt"
)

func PrintBoard(board Board) {
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

func PrintReport(report AnalyzeReport) {
	/*fmt.Println("*** Analyze report. Weights ***")
	for i := 0; i < len(report.Weights); i++ {
		fmt.Print("[")
		for j := 0; j < len(report.Weights[i]); j++ {
			fmt.Printf(" %d ", report.Weights[i][j])
		}
		fmt.Print("]\n")
	}*/
	fmt.Println("*** Analyze report. WinPositions ***")
	fmt.Println(report.WinPositions)
	fmt.Println("*** Analyze report. Defends ***")
	fmt.Println(report.Defends)
	fmt.Println("*** Analyze report. Preemptive defends ***")
	fmt.Println(report.PreemptiveDefends)
	fmt.Println("*** Analyze report. Attacks ***")
	fmt.Println(report.Attacks)
}

func PrintResult(result Result) {
	fmt.Println(result)
}