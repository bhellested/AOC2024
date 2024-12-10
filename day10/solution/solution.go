package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
)

//left,right,down,up
var positions = [][]int{{0,-1},{0,1},{1,0},{-1,0}}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type pos struct {
	i int
	j int
}

func findTrailScore(lines []string, i int, j int, depth int,foundSummits map[pos]struct{}) int{
	//first check bounds:
	if i<0 || i >=len(lines) || j<0 || j>= len(lines[0]){
		return 0
	}
	//okay, now check if were at the right character
	val,_:=strconv.Atoi(string(lines[i][j]))
	if val == depth{
		if depth == 9 {
			//we've reached a summit
			foundSummits[pos{i:i,j:j}]=struct{}{}
			return 1
		}else{
			//tally how many paths reach a summit
			score :=0
			for _,pos := range positions {
				score+=findTrailScore(lines,i+pos[0],j+pos[1],depth+1,foundSummits)
			}
			return score
		}		
	}
	return 0
}

func main(){
	dat, err := os.ReadFile("../../inputs/day10.txt")
	check(err)
	lines := strings.Split(string(dat),"\n")
	score := 0
	totalPaths :=0
	for i,line := range lines{
		for j := range (line){
			summits := make(map[pos]struct{})
			totalPaths += findTrailScore(lines,i,j,0,summits)
			if len(summits) !=0 {
				score += len(summits)
			}
		}
	}
	fmt.Println("Sum of unique summits from each trailhead (part1):", score)
	fmt.Println("Sum of total paths from each trailhead (part2):", totalPaths)
}