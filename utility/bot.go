package tic_tac_toe

import "fmt"

func ComputerTurn(board *Board, bot *Player, player *Player) {
	analyzeReport := AnalyzeResult(board, player.Result)

	PrintReport(analyzeReport)

	var positionList []Position
	if len(analyzeReport.Defends) > 0 {
		positionList = analyzeReport.Defends
	} else {
		positionList = analyzeReport.Attacks
	}

	chosenPosition := ChooseBestPosition(positionList)
	fmt.Println(positionList)
	fmt.Println(chosenPosition)
	Turn(board, bot, &chosenPosition)
}

func AnalyzeResult(board *Board, playerResult Result) AnalyzeReport {
	var analyzeReport AnalyzeReport
	boardSize := len(board.Field)
	winnableCount := boardSize - 1

	var maxWeight, bestAttackX, bestAttackY int

	for i := 0; i < boardSize; i++ {
		if playerResult.XCount[i] >= winnableCount {
			// check free space at i-row
			for j := 0; j < boardSize; j++ {
				if board.Field[i][j] == "" {
					computerPosition := Position{}
					computerPosition.X = i
					computerPosition.Y = j
					analyzeReport.Defends = append(analyzeReport.Defends, computerPosition)
				}
			}
		}
		if playerResult.YCount[i] >= winnableCount {
			for j := 0; j < boardSize; j++ {
				if board.Field[j][i] == "" {
					computerPosition := Position{}
					computerPosition.X = j
					computerPosition.Y = i
					analyzeReport.Defends = append(analyzeReport.Defends, computerPosition)
				}
			}
		}

		// check diagonals
		if playerResult.Diagonal1Count >= winnableCount {
			for i := 0; i < boardSize; i++ {
				if board.Field[i][i] == "" {
					computerPosition := Position{}
					computerPosition.X = i
					computerPosition.Y = i
					analyzeReport.Defends = append(analyzeReport.Defends, computerPosition)
				}
			}
		}
		if playerResult.Diagonal2Count >= winnableCount {
			for i := 0; i < boardSize; i++ {
				x := i
				y := winnableCount - i
				if board.Field[x][y] == "" {
					computerPosition := Position{}
					computerPosition.X = x
					computerPosition.Y = y
					analyzeReport.Defends = append(analyzeReport.Defends, computerPosition)
				}
			}
		}
	}

	for i := 0; i < len(board.Field); i++ {
		for j := 0; j < len(board.Field[i]); j++ {
			weight, err := getPositionWeight(board, i, j)
			if len(err) > 0 {
				continue
			}
			if weight > maxWeight {
				maxWeight = weight
				bestAttackX = i
				bestAttackY = j
			}
			analyzeReport.Weights[i][j] = weight
		}
	}
	analyzeReport.Attacks = append(analyzeReport.Attacks, Position{X:bestAttackX, Y:bestAttackY})

	return analyzeReport
}

func ChooseBestPosition(positionList []Position) Position {
	choices := make(map[Position]int)

	for i:=0; i < len(positionList); i++ {
		for j:=0; j < len(positionList); j++ {
			if i == j {
				continue
			}
			if positionList[i] == positionList[j] {
				choices[positionList[i]]++
			}
		}
	}
	var maxNumber int
	var chosenPosition Position
	for n := range choices {
		maxNumber = choices[n]
		chosenPosition = n
		break
	}
	for n := range choices {
		if choices[n] > maxNumber {
			chosenPosition = n
		}
	}
	return chosenPosition
}

func getPositionWeight(board *Board, x int, y int) (weight int, err string) {
	if len(board.Field[x][y]) > 0 {
		return -1, "Position must be empty!"
	}
	// check all columns with x
	var rowFreeCount int
	for j := 0; j < len(board.Field); j++ {
		if len(board.Field[x][j]) != 0 {
			continue
		}
		rowFreeCount++
	}
	if rowFreeCount == len(board.Field) {
		weight++
	}

	// check all rows with y
	var columnFreeCount int
	for i := 0; i < len(board.Field); i++ {
		if len(board.Field[i][y]) != 0 {
			continue
		}
		columnFreeCount++
	}
	if columnFreeCount == len(board.Field) {
		weight++
	}

	if x == y {
		var diagonal1FreeCount int
		for i := 0; i < len(board.Field); i++ {
			if len(board.Field[i][i]) != 0 {
				continue
			}
			diagonal1FreeCount++
		}
		if diagonal1FreeCount == len(board.Field) {
			weight++
		}
	}

	if len(board.Field) - 1 - x == y {
		var diagonal2FreeCount int
		for i := 0; i < len(board.Field); i++ {
			j := len(board.Field) - 1 - i
			if len(board.Field[i][j]) != 0 {
				continue
			}
			diagonal2FreeCount++
		}
		if diagonal2FreeCount == len(board.Field) {
			weight++
		}
	}

	return weight, ""
}
