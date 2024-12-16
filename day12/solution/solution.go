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

//up,right,down,left
var positions = [][]int{{-1,0},{0,1},{1,0},{0,-1}}

type pos struct{
	x int
	y int
}
//returns (area,perimeter)
func recurse(lines []string, i int,j int, val rune, seen map[pos]struct{},fencePositions map[pos]map[int]struct{},directionTraveled int) (int,int){
	//first check bounds:
	if i<0 || i >=len(lines) || j<0 || j>= len(lines[0]){
		_,exists := fencePositions[pos{x:i,y:j}]
		if !exists{
			fencePositions[pos{x:i,y:j}] = make(map[int]struct{})
		}
		fencePositions[pos{x:i,y:j}][(directionTraveled+1)%4] = struct{}{}
		return 0,1
	}
	_,exists := seen[pos{x:i,y:j}]
	if exists && rune(lines[i][j]) != val{
		_,exists = fencePositions[pos{x:i,y:j}]
		if !exists{
			fencePositions[pos{x:i,y:j}] = make(map[int]struct{})
		}
		fencePositions[pos{x:i,y:j}][(directionTraveled+1)%4] = struct{}{}
		return 0,1
	}else if exists{
		return 0,0
	}
	areaRet:=0
	perimRet:=0
	if rune(lines[i][j]) == val{
		areaRet+=1
		seen[pos{x:i,y:j}] = struct{}{}
		for direction,pos := range positions {
			area,perim := recurse(lines,i+pos[0],j+pos[1],val,seen,fencePositions,direction)
			areaRet += area
			perimRet += perim
		}
		return areaRet,perimRet
	} else{
		_,exists = fencePositions[pos{x:i,y:j}]
		if !exists{
			fencePositions[pos{x:i,y:j}] = make(map[int]struct{})
		}
		fencePositions[pos{x:i,y:j}][(directionTraveled+1)%4] = struct{}{}
		return 0,1
	}
}

//invert means we are working with internal fences
func iterateFencePositions(currentDirection int,i int,j int,fencePositions map[pos]map[int]struct{}) int{
	sides :=0
	//go as far as possible in the current direction
	curPos:=pos{x:i,y:j}
	posMap,posExists := fencePositions[curPos]
	if !posExists{
		return 0
	}
	_,dirExists := posMap[currentDirection]
	if !dirExists{
		return 0
	}
	sides+=1
	delete(posMap,currentDirection)
	if len(fencePositions[curPos]) == 0{
		delete(fencePositions,curPos)
	}
	for {
		curDirectionPos := 0;
		curDirectionPos = currentDirection
		curPos.x += positions[curDirectionPos][0]
		curPos.y += positions[curDirectionPos][1]
		posMap,posExists := fencePositions[curPos]
		if !posExists{
			//move back
			curPos.x -= positions[curDirectionPos][0]
			curPos.y -= positions[curDirectionPos][1]
			break
		}
		_,dirExists := posMap[currentDirection]
		if !dirExists {
			//move back
			curPos.x -= positions[curDirectionPos][0]
			curPos.y -= positions[curDirectionPos][1]
			break
		}
		delete(posMap,currentDirection)
		if len(posMap) == 0{
			delete(fencePositions,curPos)
		}
	}
	switch currentDirection{
	case 0://up
		sides+=iterateFencePositions(1,curPos.x-1,curPos.y+1,fencePositions)
		sides+=iterateFencePositions(3,curPos.x,curPos.y,fencePositions)
	case 1://right
		sides+=iterateFencePositions(0,curPos.x,curPos.y,fencePositions)
		sides+=iterateFencePositions(2,curPos.x+1,curPos.y+1,fencePositions)
	case 2://down
		sides+=iterateFencePositions(3,curPos.x+1,curPos.y-1,fencePositions)
		sides+=iterateFencePositions(1,curPos.x,curPos.y,fencePositions)
	case 3://left
		sides+=iterateFencePositions(0,curPos.x-1,curPos.y-1,fencePositions)
		sides+=iterateFencePositions(2,curPos.x,curPos.y,fencePositions)
	}
	return sides
}

func main(){
	dat, err := os.ReadFile("../../inputs/day12.txt")
	check(err)
	lines := strings.Split(string(dat),"\n")

	seen := make(map[pos]struct{})
	totalPrice:=0
	totalPrice2:=0
	for i,line := range lines{
		for j,val := range line{
			_,exists := seen[pos{x:i,y:j}]
			if !exists {
				fencePositions := make(map[pos]map[int]struct{})
				area,perimeter := recurse(lines,i,j,val,seen,fencePositions,-1)
				totalPrice+=area*perimeter
				sides := 0
				//we'll always be going right first since we know there is a fence right above us
				sides+=iterateFencePositions(1,i-1,j,fencePositions)
				for len(fencePositions) > 0 {
					for position,_ := range fencePositions{
						seen2 := make(map[pos]struct{})
						fencePositions2 := make(map[pos]map[int]struct{})
						recurse(lines,position.x,position.y,rune(lines[position.x][position.y]),seen2,fencePositions2,-1)
						for pos,posMap := range fencePositions2{
							_,dirExists := posMap[1]
							if dirExists{
								//now go to the left until we don't see a fence going to the right
								for {
									pos.y -= 1
									posMap,posExists := fencePositions2[pos]
									if !posExists{
										pos.y += 1
										break
									}
									_,dirExists = posMap[1]
									if !dirExists{
										pos.y += 1
										break
									}
								}
								sides+=iterateFencePositions(1,pos.x,pos.y,fencePositions2)
								//delete everything seen from fencePositions
								for pos,_ := range seen2{
									delete(fencePositions,pos)
								}
							}
						}
					}
				}
				totalPrice2+=sides*area
			}
		}
	}
	fmt.Println("Total price: ",totalPrice)
	fmt.Println("Total price2: ",totalPrice2)
}