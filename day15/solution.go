package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Robot struct {
	xpos int
	ypos int
}

func prettyPrintBoard(board [][]rune) {
	for _, line := range board {
		fmt.Println(string(line))
	}
}

func canMove(robot Robot, direction rune, board [][]rune) bool {
	for {
		switch direction {
		case '^':
			robot.ypos--
		case 'v':
			robot.ypos++
		case '<':
			robot.xpos--
		case '>':
			robot.xpos++
		}

		if robot.xpos < 0 || robot.xpos >= len(board[0]) || robot.ypos < 0 || robot.ypos >= len(board) {
			return false
		}
		if board[robot.ypos][robot.xpos] == '#' {
			return false
		}
		if direction == '^' || direction == 'v' {
			if board[robot.ypos][robot.xpos] == '[' {
				return canMove(robot, direction, board) && canMove(Robot{xpos: robot.xpos + 1, ypos: robot.ypos}, direction, board)
			}
			if board[robot.ypos][robot.xpos] == ']' {
				return canMove(robot, direction, board) && canMove(Robot{xpos: robot.xpos - 1, ypos: robot.ypos}, direction, board)
			}
		}

		if board[robot.ypos][robot.xpos] == '.' {
			return true
		}
	}
}
func moveRobot(robot *Robot, direction rune, board [][]rune, isBase bool, firstCall bool) {
	cloneRobot := *robot
	switch direction {
	case '^':
		cloneRobot.ypos--
	case 'v':
		cloneRobot.ypos++
	case '<':
		cloneRobot.xpos--
	case '>':
		cloneRobot.xpos++
	}
	if board[cloneRobot.ypos][cloneRobot.xpos] == '.' {
		board[cloneRobot.ypos][cloneRobot.xpos] = board[robot.ypos][robot.xpos]
	} else {
		moveRobot(&cloneRobot, direction, board, false,false)
		if direction == '^' {
			if board[cloneRobot.ypos][cloneRobot.xpos] == '[' {
				moveRobot(&Robot{xpos: cloneRobot.xpos + 1, ypos: cloneRobot.ypos}, direction, board, false,true)
			} else if board[cloneRobot.ypos][cloneRobot.xpos] == ']' {
				moveRobot(&Robot{xpos: cloneRobot.xpos - 1, ypos: cloneRobot.ypos}, direction, board, false,true)
			}
		}
		if direction == 'v' {
			if board[cloneRobot.ypos][cloneRobot.xpos] == '[' {
				moveRobot(&Robot{xpos: cloneRobot.xpos + 1, ypos: cloneRobot.ypos}, direction, board, false,true)
			} else if board[cloneRobot.ypos][cloneRobot.xpos] == ']' {
				moveRobot(&Robot{xpos: cloneRobot.xpos - 1, ypos: cloneRobot.ypos}, direction, board, false,true)
			}
		}
		board[cloneRobot.ypos][cloneRobot.xpos] = board[robot.ypos][robot.xpos]
	}
	if firstCall {
		board[robot.ypos][robot.xpos] = '.'
	}
	if isBase {
		robot.xpos = cloneRobot.xpos
		robot.ypos = cloneRobot.ypos
	}
}

func makeWideBoardline(line string) []rune {
	wideLine := ""
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '#':
			wideLine += "##"
		case '.':
			wideLine += ".."
		case 'O':
			wideLine += "[]"
		case '@':
			wideLine += "@."
		}
	}
	return []rune(wideLine)
}

func main() {
	file, err := os.ReadFile("../inputs/day15.txt")
	check(err)
	lines := strings.Split(string(file), "\n\n")

	boardLines := strings.Split(lines[0], "\n")
	board := make([][]rune, len(boardLines))
	board2 := make([][]rune, len(boardLines))
	robot := Robot{xpos: -1, ypos: -1}
	robot2 := Robot{xpos: -1, ypos: -1}
	for i, line := range boardLines {
		board[i] = []rune(line)
		board2[i] = makeWideBoardline(line)
		if robot.xpos == -1 && robot.ypos == -1 {
			for j, char := range board[i] {
				if char == '@' {
					robot.xpos = j
					robot.ypos = i
				}
			}
		}
		if robot2.xpos == -1 && robot2.ypos == -1 {
			for j, char := range board2[i] {
				if char == '@' {
					robot2.xpos = j
					robot2.ypos = i
				}
			}
		}
	}
	movelines := strings.Split(lines[1], "\n")
	for _, line := range movelines {
		for _, char := range line {
			if canMove(robot, char, board) {
				moveRobot(&robot, char, board, true,true)
			}
			if canMove(robot2, char, board2) {
				moveRobot(&robot2, char, board2, true,true)
			}
		}
	}
	totalOfGoods := 0
	totalOfGoods2 := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'O' {
				totalOfGoods += (i)*100 + (j)
			}
		}
	}
	for i := 0; i < len(board2); i++ {
		for j := 0; j < len(board2[i]); j++ {
			if board2[i][j] == '[' {
				totalOfGoods2 += (i)*100 + (j)
			}
		}
	}
	fmt.Println("Total part 1: ", totalOfGoods)
	fmt.Println("Total part 2: ", totalOfGoods2)
}
