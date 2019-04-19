package tic_tac_toe

import "fmt"

func ComputerTurn(board *Board, bot *Player, player *Player) {
	analyzeReport := AnalyzeResult(board, bot, player)

	PrintReport(analyzeReport)

	var positionList []Position
	if len(analyzeReport.WinPositions) > 0 {
		positionList = analyzeReport.WinPositions
	} else if len(analyzeReport.Defends) > 0 {
		positionList = analyzeReport.Defends
	} else if len(analyzeReport.PreemptiveDefends) > 0 {
		positionList = analyzeReport.PreemptiveDefends
	} else {
		positionList = analyzeReport.Attacks
	}

	chosenPosition := ChooseBestPosition(positionList)

	fmt.Println(positionList)
	fmt.Println(chosenPosition)

	Turn(board, bot, &chosenPosition)
}

func ChooseBestPosition(positionList []Position) Position {
	choices := make(map[Position]int)

	for i := 0; i < len(positionList); i++ {
		for j := 0; j < len(positionList); j++ {
			if choices[positionList[i]] > 0 && i == j {
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

func AnalyzeResult(board *Board, bot *Player, player *Player) AnalyzeReport {
	var analyzeReport AnalyzeReport

	boardSize := len(board.Field)
	analyzeReport.Defends = getWinnablePositions(board, player, boardSize - 1)
	analyzeReport.WinPositions = getWinnablePositions(board, bot, boardSize - 1)

	bestAttackX, bestAttackY, maxWeight := getBestAttack(board, bot)
	// if there are no win position for bot
	if maxWeight == 0 {
		bestAttackX, bestAttackY, _ = getBestAttack(board, player)
	}
	analyzeReport.Attacks = append(analyzeReport.Attacks, Position{X: bestAttackX, Y: bestAttackY})

	return analyzeReport
}

func getWinnablePositions(board *Board, player *Player, markCount int) []Position {
	boardSize := len(board.Field)

	var winnablePositionList []Position
	for i := 0; i < boardSize; i++ {
		if player.Result.XCount[i] >= markCount {
			// check free space at i-row
			for j := 0; j < boardSize; j++ {
				if board.Field[i][j] == "" {
					computerPosition := Position{}
					computerPosition.X = i
					computerPosition.Y = j
					winnablePositionList = append(winnablePositionList, computerPosition)
				}
			}
		}
		if player.Result.YCount[i] >= markCount {
			for j := 0; j < boardSize; j++ {
				if board.Field[j][i] == "" {
					computerPosition := Position{}
					computerPosition.X = j
					computerPosition.Y = i
					winnablePositionList = append(winnablePositionList, computerPosition)
				}
			}
		}

		// check diagonals
		if player.Result.Diagonal1Count >= markCount {
			for i := 0; i < boardSize; i++ {
				if board.Field[i][i] == "" {
					computerPosition := Position{}
					computerPosition.X = i
					computerPosition.Y = i
					winnablePositionList = append(winnablePositionList, computerPosition)
				}
			}
		}
		if player.Result.Diagonal2Count >= markCount {
			for i := 0; i < boardSize; i++ {
				x := i
				y := len(board.Field)-1-i
				if board.Field[x][y] == "" {
					computerPosition := Position{}
					computerPosition.X = x
					computerPosition.Y = y
					winnablePositionList = append(winnablePositionList, computerPosition)
				}
			}
		}
	}
	return winnablePositionList
}

func getBestAttack(board *Board, player *Player) (bestAttackX int, bestAttackY int, maxWeight int) {
	for i := 0; i < len(board.Field); i++ {
		for j := 0; j < len(board.Field[i]); j++ {
			if len(board.Field[i][j]) > 0 {
				continue
			}
			weight, err := GetPositionWeight(board, player.Mark, i, j)
			if len(err) > 0 {
				continue
			}
			if weight > maxWeight {
				maxWeight = weight
				bestAttackX = i
				bestAttackY = j
			}
		}
	}
	return bestAttackX, bestAttackY, maxWeight
}

func GetPositionWeight(board *Board, mark string, x int, y int) (weight int, err string) {
	if len(board.Field[x][y]) > 0 && board.Field[x][y] != mark {
		return -1, "Position must be empty!"
	}
	// check all columns with x
	var rowMarkedOrFreeCount int
	for j := 0; j < len(board.Field); j++ {
		if len(board.Field[x][j]) != 0 && board.Field[x][j] != mark {
			continue
		}
		rowMarkedOrFreeCount++
	}
	if rowMarkedOrFreeCount == len(board.Field) {
		weight++
	}

	// check all rows with y
	var columnMarkedOrFreeCount int
	for i := 0; i < len(board.Field); i++ {
		if len(board.Field[i][y]) != 0 && board.Field[i][y] != mark {
			continue
		}
		columnMarkedOrFreeCount++
	}
	if columnMarkedOrFreeCount == len(board.Field) {
		weight++
	}

	if x == y {
		var diagonal1MarkedOrFreeCount int
		for i := 0; i < len(board.Field); i++ {
			if len(board.Field[i][i]) != 0 && board.Field[i][i] != mark {
				continue
			}
			diagonal1MarkedOrFreeCount++
		}
		if diagonal1MarkedOrFreeCount == len(board.Field) {
			weight++
		}
	}

	if len(board.Field)-1-x == y {
		var diagonal2MarkedOrFreeCount int
		for i := 0; i < len(board.Field); i++ {
			j := len(board.Field) - 1 - i
			if len(board.Field[i][j]) != 0 && board.Field[i][j] != mark {
				continue
			}
			diagonal2MarkedOrFreeCount++
		}
		if diagonal2MarkedOrFreeCount == len(board.Field) {
			weight++
		}
	}

	return weight, ""
}
