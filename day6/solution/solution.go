package main

import (
	"fmt"
	"os"
	"strings"
)	

type Pos struct {
	x,y int
}
var Directions = [][]int{{-1,0},{0,1},{1,0},{0,-1}}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func IsGoodObstruction(linesOriginal []string,visitedOriginal map[Pos]map[int]struct{},curDirectionIndex int, position Pos,obstructions map[string]struct{}) {
	//make a copy of the lines
	lines := make([][]rune,len(linesOriginal))
	for i,line := range linesOriginal {
		lines[i] = make([]rune,len(line))
		copy(lines[i],[]rune(line))
	}
	//add an obstruction in front of us
	positionForObstruction := Pos{position.x+Directions[curDirectionIndex][0],position.y+Directions[curDirectionIndex][1]}
	lines[positionForObstruction.x][positionForObstruction.y] = '$'
	visitedCopy := make(map[Pos]map[int]struct{})
	
	
	for {
		nextPos := Pos{position.x+Directions[curDirectionIndex][0],position.y+Directions[curDirectionIndex][1]}
		if nextPos.x < 0 || nextPos.x >= len(lines) || nextPos.y < 0 || nextPos.y >= len(lines[0]) {
			return
		}else if lines[nextPos.x][nextPos.y] == '#' || lines[nextPos.x][nextPos.y] == '$' {
			//turn
			curDirectionIndex = (curDirectionIndex+1) % 4
		}else {
			val,exists := visitedCopy[nextPos]
			if !exists {
				visitedCopy[nextPos] = make(map[int]struct{})
			} else{
				_, dirExists := val[curDirectionIndex]
                if dirExists {
					obstructions[fmt.Sprintf("%d,%d",positionForObstruction.x,positionForObstruction.y)] = struct{}{}
					return
                }
			}
			visitedCopy[nextPos][curDirectionIndex]=struct{}{}
			position = nextPos
		}
	}
}


func main() {
	dat, err := os.ReadFile("../../inputs/day6.txt")
	check(err)
	var position Pos
	lines:=strings.Split(string(dat), "\n")
	visited := make(map[Pos]map[int]struct{})//length of visited cells holds the answer for part 1
	curDirectionIndex :=0
	for i,line := range lines {
		index := strings.Index(line,"^")
		if index != -1 {
			position = Pos{i,index}
			visited[position] = make(map[int]struct{})
			visited[position][curDirectionIndex] = struct{}{}
		}
	}
	
	obstructions:=make(map[string]struct{})
	//this is to make sure we only check for a loop the first time we see a position, otherwise we wouldn't already passed it
	alreadyCheckedForLoopPositions := make(map[string]struct{})

	for {
		nextPos := Pos{position.x+Directions[curDirectionIndex][0],position.y+Directions[curDirectionIndex][1]}
		if nextPos.x < 0 || nextPos.x >= len(lines) || nextPos.y < 0 || nextPos.y >= len(lines[0]) {
			//we've exited the map
			break 
		}else if lines[nextPos.x][nextPos.y] == '#' {
			//turn
			curDirectionIndex = (curDirectionIndex+1) % 4
		}else {
			//see if this is a good place for an obstruction,
			if _, ok := alreadyCheckedForLoopPositions[fmt.Sprintf("%d,%d",nextPos.x,nextPos.y)]; !ok {
				IsGoodObstruction(lines,visited,curDirectionIndex,position,obstructions)
				alreadyCheckedForLoopPositions[fmt.Sprintf("%d,%d",nextPos.x,nextPos.y)] = struct{}{}
			}
			val := visited[nextPos]
			if val == nil {
				visited[nextPos] = make(map[int]struct{})
			}
			visited[nextPos][curDirectionIndex] = struct{}{}
			position = nextPos
		}
	}
	
	fmt.Println("Number of cells visited: (part1)",len(visited))
	fmt.Println("Number of good obstructions: (part2)", len(obstructions))
}