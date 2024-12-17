package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type button struct{
	right int64
	up int64
}
type game struct{
	a button//costs 3 to press
	b button//costs 1 to press
	goalX int64
	goalY int64
}

func parseButton(line string) button{
	split := strings.Split(line,":")
	coords := strings.Split(split[1],",")
	x := strings.TrimSpace(coords[0])
	y := strings.TrimSpace(coords[1])
	x = strings.TrimPrefix(x,"X+")
	y = strings.TrimPrefix(y,"Y+")
	right,err := strconv.ParseInt(x,10,64)
	check(err)
	up,err := strconv.ParseInt(y,10,64)
	check(err)
	return button{
		right: right,
		up: up,
	}
}

func parseGame(line string) game{
	game := game{}
	split := strings.Split(line,"\n")
	buttonA := parseButton(split[0])
	buttonB := parseButton(split[1])
	game.a = buttonA
	game.b = buttonB
	split2 := strings.Split(split[2]," ")
	x := strings.TrimPrefix(split2[1],"X=")
	y := strings.TrimPrefix(split2[2],"Y=")
	x = strings.TrimSuffix(x," ")
	y = strings.TrimSuffix(y," ")
	xInt,err := strconv.Atoi(x[0:len(x)-1])
	check(err)
	yInt,err := strconv.Atoi(y)
	check(err)
	game.goalX = int64(xInt)
	game.goalY = int64(yInt)
	return game
}

func areSimilar(a,b button) bool{
	return a.right/a.up == b.right/b.up &&
		(a.right%a.up)*b.up == (b.right%b.up)*a.up
}
func solve(game game) (int64,int64,error){
	//first check the similarity of the goal and the buttons
	if areSimilar(button{game.goalX,game.goalY}, game.a) {
		if areSimilar(game.a, game.b) {
			return 0,0,errors.New("A B and Game are all similar. A solution may exist but is not yet implemented.")
		}
	} else if areSimilar(button{game.goalX,game.goalY}, game.b) {
		if game.goalX%game.b.right == 0 {
			return game.goalX / game.b.right,0,nil
		} else {
			return 0,0,errors.New("B and Game similar but A is not, B cannot reach the goal")
		}
	} else if areSimilar(game.a, game.b) {
		return 0,0,errors.New("A and B similar but Game is not, no solutions possible.")
	}

	//solve for A presses 
	numerator := float64(game.goalY) - (float64(game.goalX)*float64(game.b.up)/float64(game.b.right))
	denominator := float64(game.a.up) - ((float64(game.a.right)*float64(game.b.up))/float64(game.b.right))
	numA := numerator/denominator

	numB := (float64(game.goalX)-(numA*float64(game.a.right)))/float64(game.b.right)
	bRounded := math.Round(numB)
	aRounded := math.Round(numA)
	actualX := int64(aRounded)*game.a.right + int64(bRounded)*game.b.right
	actualY := int64(aRounded)*game.a.up + int64(bRounded)*game.b.up
	if actualX == game.goalX && actualY == game.goalY{
		return int64(aRounded),int64(bRounded),nil
	}else{
		//fmt.Println("\tGame goals:",game.goalX,game.goalY,"\n\tactual:",actualX,actualY)
		return 0,0,errors.New("goal not reached")
	}
}

func main(){
	dat, err := os.ReadFile("../inputs/day13.txt")
	check(err)
	games := strings.Split(string(dat),"\n\n")
	totalCost := int64(0)
	totalCostPart2 := int64(0)
	for _,game := range games{
		game := parseGame(game)
		aPresses,bPressed,err := solve(game)
		if err == nil{
			totalCost += aPresses*3 + bPressed*1
		}
		game.goalX += 10000000000000
		game.goalY += 10000000000000
		aPresses,bPressed,err = solve(game)
		if err == nil{
			totalCostPart2 += aPresses*3 + bPressed*1
		}
	}
	fmt.Println(totalCost)
	fmt.Println(totalCostPart2)
}