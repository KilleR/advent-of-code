package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

type bingoCell struct {
	value  int
	marked bool
}

type bingoBoard struct {
	board [][]*bingoCell
}

func (b *bingoBoard) mark(n int) {
	for _, row := range b.board {
		for _, cell := range row {
			if cell.value == n {
				cell.marked = true
			}
		}
	}
}

func (b bingoBoard) print() {
	for _, row := range b.board {
		for _, cell := range row {
			fmt.Print(*cell)
		}
		fmt.Print("\n")
	}
}

func (b bingoBoard) sumUnmarked() (total int) {
	for _, row := range b.board {
		for _, cell := range row {
			if !cell.marked {
				total += cell.value
			}
		}
	}

	return total
}

func day4() {
	input := utils.GetInput(getInputFile("day4", false))

	var gameNumbers []int
	var boards []*bingoBoard
	var currentBoard *bingoBoard
	var err error

	for _, line := range input {

		if len(gameNumbers) == 0 {
			gameNumbers, err = utils.StringSliceToInts(strings.Split(line, ","))
			if err != nil {
				log.Fatalln("Failed to parse game numbers", err)
			}
			continue
		}

		if len(line) == 0 {
			currentBoard = new(bingoBoard)
			boards = append(boards, currentBoard)
			continue
		}

		cells, err := utils.StringSliceToInts(strings.Split(strings.Trim(strings.Replace(line, "  ", " ", -1), " "), " "))
		if err != nil {
			log.Fatalln("Invalid line:", line, err)
		}
		var boardRow []*bingoCell

		for _, cell := range cells {
			boardRow = append(boardRow, &bingoCell{
				value:  cell,
				marked: false,
			})
		}
		currentBoard.board = append(currentBoard.board, boardRow)
	}

	var winningIndex int
	var winningBoard *bingoBoard
	var lastWinningScore int

gameLoop:
	for _, num := range gameNumbers {
		for _, board := range boards {
			board.mark(num)
		}
		fmt.Println("remaining boards:", len(boards))
		winningIndex = 1
		for winningIndex != 0 {
			winningIndex, winningBoard = checkForGameEnd(boards)
			if winningBoard != nil {
				lastWinningScore = winningBoard.sumUnmarked() * num
				fmt.Println("score", lastWinningScore, num, winningBoard.sumUnmarked())
				if len(boards) > 1 {
					boards = append(boards[:winningIndex], boards[(winningIndex+1):]...)
				} else {
					fmt.Println("Victory with score:", winningBoard.sumUnmarked()*num)
					break gameLoop
				}
			}
		}
	}

	fmt.Println("Final Victory with score:", lastWinningScore)
}

func checkForGameEnd(boards []*bingoBoard) (int, *bingoBoard) {
	for i, board := range boards {
		colVictory := make([]bool, len(board.board[0]))
		for i := range colVictory {
			colVictory[i] = true
		}
		for _, row := range board.board {
			rowVictory := true
			for i, cell := range row {
				if !cell.marked {
					rowVictory = false
					colVictory[i] = false
				}
			}
			if rowVictory {
				return i, board
			}
		}
		for _, victory := range colVictory {
			if victory {
				return i, board
			}
		}
	}
	return 0, nil
}

func add() {

}
