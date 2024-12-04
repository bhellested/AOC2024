package main

import (
	"fmt"
	"os"
	"strings"
)	
var XMAS = []rune("XMAS")
var directions = [][]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{1,0},{1,1}}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkForXMAS(count *uint32, letters [][]rune, i int, j int) {
	for _,direction := range directions {
		for x:=0;x<len(XMAS);x++{
			newI:=i+direction[0]*x;
			newJ:=j+direction[1]*x;
			if newI<0 || newI >=len(letters) || newJ<0 || 
			newJ>=len(letters[0]) || letters[newI][newJ]!=XMAS[x]{
				break
			}
			if(x==3){
				*count++;
			}
		}
	}
}

func main() {
	dat, err := os.ReadFile("../../inputs/day4.txt")
    check(err)
	lines:=strings.Split(string(dat), "\n")
	letters := make([][]rune, len(lines))
	for i := range letters {
		letters[i] = make([]rune, len(lines[0]))
		for j := range letters[i]{
			letters[i][j]=rune(lines[i][j])
		}
	}
	var count uint32 = 0
	for i:=0; i< len(letters); i++ {
		for j :=0; j<len(letters[0]);j++{
			checkForXMAS(&count,letters,i,j)
		}
	}
	fmt.Println("Count: ",count)
}