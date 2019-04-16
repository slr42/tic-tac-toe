package main

import (
	"fmt"
)

type AnalyzeReport struct {
	weights [3][3]int
}

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"", "", ""},
		{"", "", ""},
		{"", "", ""},
	}

	printBoard(board)
	report := analyzeBoard(board)
	printReport(report)
}

func analyzeBoard(board [][]string) AnalyzeReport {
	analyzeReport := AnalyzeReport{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			weight, err := getPositionWeight(board, i, j)
			if len(err) > 0 {
				continue
			}
			analyzeReport.weights[i][j] = weight
		}
	}
	return analyzeReport
}

func getPositionWeight(board [][]string, x int, y int) (weight int, err string) {
	if len(board[x][y]) > 0 {
		return -1, "Position must be empty!"
	}
	// check all columns with x
	var rowFreeCount int
	for j := 0; j < len(board); j++ {
		if len(board[x][j]) != 0 {
			continue
		}
		rowFreeCount++
	}
	if rowFreeCount == len(board) {
		weight++
	}

	// check all rows with y
	var columnFreeCount int
	for i := 0; i < len(board); i++ {
		if len(board[i][y]) != 0 {
			continue
		}
		columnFreeCount++
	}
	if columnFreeCount == len(board) {
		weight++
	}

	if x == y {
		var diagonal1FreeCount int
		for i := 0; i < len(board); i++ {
			if len(board[i][i]) != 0 {
				continue
			}
			diagonal1FreeCount++
		}
		if diagonal1FreeCount == len(board) {
			weight++
		}
	}

	if len(board) - 1 - x == y {
		var diagonal2FreeCount int
		for i := 0; i < len(board); i++ {
			j := len(board) - 1 - i
			if len(board[i][j]) != 0 {
				continue
			}
			diagonal2FreeCount++
		}
		if diagonal2FreeCount == len(board) {
			weight++
		}
	}

	return weight, ""
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Print("[")
		for j := 0; j < len(board[i]); j++ {
			symbol := board[i][j]
			if len(symbol) == 0 {
				symbol = "_"
			}
			fmt.Printf(" %s ", symbol)
		}
		fmt.Print("]\n")
	}
}

func printReport(report AnalyzeReport) {
	fmt.Println("*** Analyze report ***")
	for i := 0; i < len(report.weights); i++ {
		fmt.Print("[")
		for j := 0; j < len(report.weights[i]); j++ {
			fmt.Printf(" %d ", report.weights[i][j])
		}
		fmt.Print("]\n")
	}
}
