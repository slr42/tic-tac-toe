package tic_tac_toe

import (
	"fmt"
	"math/rand"
	"time"
)

func ComputerTurn(board *Board, bot *Player, player *Player) {
	analyzeReport := AnalyzeResult(board, bot, player)

	PrintReport(analyzeReport)

	var positionList []Position
	switch {
	case len(analyzeReport.WinPositions) > 0:
		positionList = analyzeReport.WinPositions
	case len(analyzeReport.Defends) > 0:
		positionList = analyzeReport.Defends
	case len(analyzeReport.PreemptiveDefends) > 0:
		positionList = analyzeReport.PreemptiveDefends
	default:
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

	var bestPositions []Position
	for n := range choices {
		if choices[n] > maxNumber {
			bestPositions = append(bestPositions, n)
		}
	}

	if len(bestPositions) > 1 {
		s := rand.NewSource(time.Now().Unix())
		r := rand.New(s)
		chosenPosition = bestPositions[r.Int() % len(bestPositions)]
	} else if len(bestPositions) == 1 {
		chosenPosition = bestPositions[0]
	}

	return chosenPosition
}

func AnalyzeResult(board *Board, bot *Player, player *Player) AnalyzeReport {
	var analyzeReport AnalyzeReport

	boardSize := len(board.Field)
	analyzeReport.Defends = getWinnablePositions(board, player, boardSize - 1)
	analyzeReport.WinPositions = getWinnablePositions(board, bot, boardSize - 1)

	analyzeReport.PreemptiveDefends = getPreemptiveDefends(board, player, bot)

	analyzeReport.Attacks = getAttacks(board, bot)
	if len(analyzeReport.Attacks) == 0 {
		analyzeReport.Attacks = getAttacks(board, player)
	}

	return analyzeReport
}

func getPreemptiveDefends(board *Board, player *Player, bot *Player) []Position {
	playerForkPositions := getForkPositions(board, player)

	var preemptiveDefendList []Position

	// for each free cell we need to calc like this cell is marked by bot
	for i, row := range board.Field {
		for j, value := range row {
			if len(value) != 0 {
				continue
			}
			// need to calc such win position on next turn that not in forkPositions
			// find 2-cells winnable positions
			twoCellPositions := getWinnablePositions(board, bot, len(board.Field)-2)
			for _, twoCellPosition := range twoCellPositions {
				possibleWinnablePositions := getPossibleWinnablePositions(board, bot, twoCellPosition.X, twoCellPosition.Y)
				for _, possibleWinnablePosition := range possibleWinnablePositions {
					// if possibleWinnablePosition is in player fork position list
					if positionInPositionList(possibleWinnablePosition, playerForkPositions) {
						// then skip it
						continue
					}
					// else it means we have found anti preemptive defend position
					if !positionInPositionList(twoCellPosition, preemptiveDefendList) {
						preemptiveDefendList = append(preemptiveDefendList, twoCellPosition)
					}
				}
			}

			if isForkPosition(board, bot, i, j) {
				preemptiveDefend := Position{i, j}
				preemptiveDefendList = append(preemptiveDefendList, preemptiveDefend)
			}
		}
	}
	return preemptiveDefendList
}

func getForkPositions(board *Board, player *Player) []Position {
	// boardSize := len(board.Field)
	var forkPositionList []Position

	// for each free cell we need to calc like this cell is marked by player
	for i, row := range board.Field {
		for j, value := range row {
			if len(value) != 0 {
				continue
			}
			if isForkPosition(board, player, i, j) {
				forkPosition := Position{i, j}
				forkPositionList = append(forkPositionList, forkPosition)
			}
		}
	}
	return forkPositionList
}

func isForkPosition(board *Board, player *Player, x int, y int) bool {
	possibleWinnablePositions := getPossibleWinnablePositions(board, player, x, y)
	return len(possibleWinnablePositions) > 1
}

func getPossibleWinnablePositions(board *Board, player *Player, x int, y int) []Position {
	boardSize := len(board.Field)

	possibleBoard := cloneBoard(board)
	possiblePlayer := clonePlayer(player)
	possiblePosition := Position{x, y}
	markPosition(&possibleBoard, possiblePlayer, &possiblePosition)
	possibleWinnablePositions := getWinnablePositions(&possibleBoard, possiblePlayer, boardSize - 1)

	return possibleWinnablePositions
}

func getWinnablePositions(board *Board, player *Player, markCount int) []Position {
	var winnablePositionList []Position
	boardSize := len(board.Field)
	for i := 0; i < boardSize; i++ {
		if player.Result.XCount[i] >= markCount {
			// if row is not free -- skip it
			if isRowFree(board, i, player.Mark) {
				// check free space at i-row
				for j := 0; j < boardSize; j++ {
					if board.Field[i][j] == "" {
						winnablePositionList = append(winnablePositionList, Position{i, j})
					}
				}
			}
		}
		if player.Result.YCount[i] >= markCount {
			// if column is not free -- skip it
			if isColumnFree(board, i, player.Mark) {
				for j := 0; j < boardSize; j++ {
					if board.Field[j][i] == "" {
						winnablePositionList = append(winnablePositionList, Position{j, i})
					}
				}
			}
		}

	}
	// check diagonals
	if player.Result.Diagonal1Count >= markCount && isDiagonal1Free(board, player.Mark) {
		for i, row := range board.Field {
			if row[i] == "" {
				winnablePositionList = append(winnablePositionList, Position{i,i})
			}
		}
	}
	if player.Result.Diagonal2Count >= markCount && isDiagonal2Free(board, player.Mark) {
		for x, row := range board.Field {
			y := len(board.Field)-1-x
			if row[y] == "" {
				winnablePositionList = append(winnablePositionList, Position{x, y})
			}
		}
	}
	return winnablePositionList
}

func getAttacks(board *Board, player *Player) []Position {
	attackPositions := getForkPositions(board, player)

	if len(attackPositions) == 0 {
		var maxWeight int
		for i, row := range board.Field {
			for j, value := range row {
				if len(value) > 0 {
					continue
				}
				weight, err := GetPositionWeight(board, player.Mark, i, j)
				if len(err) > 0 {
					continue
				}
				if weight >= maxWeight {
					maxWeight = weight
					attackPositions = append(attackPositions, Position{i, j})
				}
			}
		}
	}
	return attackPositions
}

func GetPositionWeight(board *Board, mark string, x int, y int) (weight int, err string) {
	if len(board.Field[x][y]) > 0 && board.Field[x][y] != mark {
		return -1, "Position must be empty!"
	}
	// check all columns with x
	var rowMarkedOrFreeCount int
	boardSize := len(board.Field)
	for j := 0; j < boardSize; j++ {
		if len(board.Field[x][j]) != 0 && board.Field[x][j] != mark {
			continue
		}
		rowMarkedOrFreeCount++
	}
	if rowMarkedOrFreeCount == boardSize {
		weight++
	}

	// check all rows with y
	var columnMarkedOrFreeCount int
	for _, row := range board.Field {
		if len(row[y]) != 0 && row[y] != mark {
			continue
		}
		columnMarkedOrFreeCount++
	}
	if columnMarkedOrFreeCount == boardSize {
		weight++
	}

	if x == y {
		var diagonal1MarkedOrFreeCount int
		for i, row := range board.Field {
			if len(row[i]) != 0 && row[i] != mark {
				continue
			}
			diagonal1MarkedOrFreeCount++
		}
		if diagonal1MarkedOrFreeCount == boardSize {
			weight++
		}
	}

	if boardSize-1-x == y {
		var diagonal2MarkedOrFreeCount int
		for i, row := range board.Field {
			j := len(board.Field) - 1 - i
			if len(row[j]) != 0 && row[j] != mark {
				continue
			}
			diagonal2MarkedOrFreeCount++
		}
		if diagonal2MarkedOrFreeCount == boardSize {
			weight++
		}
	}
	return weight, ""
}
