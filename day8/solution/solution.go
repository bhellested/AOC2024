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

type pos struct {
	x int
	y int
}

func CheckBounds(p pos, uniqueAntinodes map[pos]struct{},maxX int, maxY int) bool {
	if(p.x<0 || p.y<0 || p.x>maxX-1 || p.y > maxY-1){
		return false
	}
	uniqueAntinodes[p] = struct{}{}
	return true
}

func Distance(p1 pos, p2 pos) (int,int) {
	return p2.x - p1.x, p2.y - p1.y
}

func CheckAllDistances(p1 pos, p2 pos, uniqueAntinodes map[pos]struct{},maxX int, maxY int) {
	distX,distY := Distance(p1,p2)
	iteration:=0
	for {
		newPos := pos{x:p2.x+(distX*iteration),y:p2.y+(distY*iteration)}
		if(CheckBounds(newPos,uniqueAntinodes,maxX,maxY)){
			iteration++
		}else{
			iteration =0
			break;
		}
	}

	for {
		newPos:= pos{x:p1.x-(distX*iteration),y:p1.y-(distY*iteration)}
		if(CheckBounds(newPos,uniqueAntinodes,maxX,maxY)){
			iteration++
		}else{
			iteration =0
			break;
		}
	}
}

func main(){
	dat, err := os.ReadFile("../../inputs/day8.txt")
	check(err)
	lines:=strings.Split(string(dat), "\n")
	maxX:=len(lines)
	maxY:=len(lines[0])
	antennaMap := make(map[rune][]pos)

	for i,line := range lines {
		for j,runeVal := range line {
			if runeVal != '.'{
				val,exists := antennaMap[runeVal]
				if(exists){
					antennaMap[runeVal] = append(val,pos{x:i,y:j})
				} else {
					antennaMap[runeVal] = []pos{{x: i, y: j}}
				}
			}
		}
	}
	uniqueAntinodes := make(map[pos]struct{})
	uniqueAntinodesPart2 := make(map[pos]struct{})
	for _,antennaType := range antennaMap{
		for i:=0; i<len(antennaType)-1;i++{
			for j:=i+1;j<len(antennaType);j++{
				//find the distance from first to second:
				distX,distY := Distance(antennaType[i],antennaType[j])
				firstCandidate := pos{x:antennaType[i].x-distX, y:antennaType[i].y-distY}
				CheckBounds(firstCandidate,uniqueAntinodes,maxX,maxY)
				secondCandidate := pos{x:antennaType[j].x+distX, y:antennaType[j].y+distY}
				CheckBounds(secondCandidate,uniqueAntinodes,maxX,maxY)

				//solve part2 for this pair
				CheckAllDistances(antennaType[i],antennaType[j],uniqueAntinodesPart2,maxX,maxY)
			}
		}
	}
	fmt.Println("Part1: ",len(uniqueAntinodes))
	fmt.Println("Part2: ",len(uniqueAntinodesPart2))
}