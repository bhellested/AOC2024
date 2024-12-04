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

func checkForXMAS(count *uint32, letters [][]rune, i int, j int) {
	if ((letters[i-1][j-1]=='M' && letters[i+1][j+1]=='S') || (letters[i-1][j-1]=='S' && letters[i+1][j+1]=='M')) && 
	   ((letters[i-1][j+1]=='M' && letters[i+1][j-1]=='S') || (letters[i-1][j+1]=='S' && letters[i+1][j-1]=='M')) {
		*count++
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
	for i:=1; i < len(letters)-1; i++ {
		for j:=1; j<len(letters[0])-1;j++{
			if(letters[i][j]=='A'){
				checkForXMAS(&count,letters,i,j)
			}
		}
	}
	fmt.Println("Count: ",count)
}
