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

var directions = []vector{
	{x: 1, y: 0},
	{x: 0, y: 1},
	{x: -1, y: 0},
	{x: 0, y: -1},
}

type vector struct{
	x int
	y int
}

type robot struct{
	location vector
	velocity vector
}

func parseRobot(line string) robot{
	split := strings.Split(line," ")
	pos := strings.TrimPrefix(split[0],"p=")
	vel := strings.TrimPrefix(split[1],"v=")
	
	posCoords := strings.Split(pos,",")
	velCoords := strings.Split(vel,",")
	
	posX := 0
	posY := 0
	velX := 0 
	velY := 0
	
	fmt.Sscanf(posCoords[0], "%d", &posX)
	fmt.Sscanf(posCoords[1], "%d", &posY)
	fmt.Sscanf(velCoords[0], "%d", &velX)
	fmt.Sscanf(velCoords[1], "%d", &velY)
	
	return robot{
		location: vector{x: posX, y: posY},
		velocity: vector{x: velX, y: velY},
	}
}

func moveRobot(robot robot, boardWidth int, boardHeight int) robot{
	robot.location.x += robot.velocity.x
	robot.location.y += robot.velocity.y
	if(robot.location.x<0){
		robot.location.x = boardWidth + (robot.location.x%boardWidth)
	}else{
		robot.location.x = robot.location.x%boardWidth
	}
	if(robot.location.y<0){
		robot.location.y = boardHeight + (robot.location.y%boardHeight)
	}else{
		robot.location.y = robot.location.y%boardHeight
	}
	return robot
}

func printBoard(robots []robot, boardWidth int, boardHeight int, logFile *os.File){
	positions := make(map[vector]int)
	for _,robot := range robots{
		positions[robot.location]++
	}
	for i := 0; i < boardHeight; i++ {
		for j := 0; j < boardWidth; j++ {
			pos := vector{x: j, y: i}
			if _, ok := positions[pos]; ok {
				fmt.Fprint(logFile,positions[pos])
			} else {
				fmt.Fprint(logFile,".")
			}
		}
		fmt.Fprintln(logFile)
	}
}

func calculateSafetyFactor(robots []robot, boardWidth int, boardHeight int) int{
	firstQuadCount := 0
	secondQuadCount := 0
	thirdQuadCount := 0
	fourthQuadCount := 0
	for _,robot := range robots{
		if(robot.location.x<boardWidth/2 && robot.location.y<boardHeight/2){
			firstQuadCount++
		}else if(robot.location.x>boardWidth/2 && robot.location.y<boardHeight/2){
			secondQuadCount++
		}else if(robot.location.x<boardWidth/2 && robot.location.y>boardHeight/2){
			thirdQuadCount++
		}else if(robot.location.x>boardWidth/2 && robot.location.y>boardHeight/2){
			fourthQuadCount++
		}
	}
	return firstQuadCount*secondQuadCount*thirdQuadCount*fourthQuadCount
}

func calculateAdjacent(robots []robot, boardWidth int, boardHeight int) int{
	positions := make(map[vector]int)
	for _,robot := range robots{
		positions[robot.location]++
	}
	adjacent := 0
	for _,robot := range robots{
		for _,direction := range directions{
			newLocation := vector{x: robot.location.x+direction.x, y: robot.location.y+direction.y}
			if(newLocation.x>=0 && newLocation.x<boardWidth && newLocation.y>=0 && newLocation.y<boardHeight){
				_,ok := positions[newLocation]
				if(ok){
					adjacent++
				}
			}
		}
	}
	return adjacent
}

func main(){
	dat, err := os.ReadFile("../inputs/day14.txt")
	logFile, err := os.Create("log.txt")
	check(err)
	check(err)
	lines := strings.Split(string(dat),"\n")
	robots := make([]robot,len(lines))
	for i,line := range lines{
		robots[i] = parseRobot(line)
	}
	
	boardWidth:=101
	boardHeight:=103
	turns := boardWidth*boardHeight
	maxAdjacent :=0
	frameOfMaxAdjacent :=0
	for i:=0;i<turns;i++{
		for i,robot := range robots{
			robots[i] = moveRobot(robot,boardWidth,boardHeight)
		}
		//printBoard(robots,boardWidth,boardHeight,logFile)
		adjacent := calculateAdjacent(robots,boardWidth,boardHeight)
		if(adjacent>maxAdjacent){
			maxAdjacent = adjacent
			frameOfMaxAdjacent = i+1
			printBoard(robots,boardWidth,boardHeight,logFile)//capture the board when we have a new max
		}
		// if(i%100==0){
		// 	fmt.Println("on iteration",i,"max adjacent is",maxAdjacent,"at frame",frameOfMaxAdjacent)
		// }
		if(i==99){
			//answer to part 1
			fmt.Println("Safety factor:",calculateSafetyFactor(robots,boardWidth,boardHeight))
		}
	}
	//answer to part 2
	fmt.Println("Max adjacent:",maxAdjacent,"at frame",frameOfMaxAdjacent)
	
}