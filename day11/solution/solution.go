package main

import (
	"fmt"
	"strconv"
	"os"
	"strings"
	"container/list"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//brute force method, 25 iterations works OK but crawls to a halt around the 40 mark.
func iterateRules(stones *list.List){
	/*
If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
If the stone is engraved with a number that has an even number of digits, it is replaced by two stones. The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

	*/

	for e := stones.Front(); e != nil; e = e.Next() {
		if(e.Value==0){
			e.Value=1
		} else {
			strVal := strconv.Itoa(e.Value.(int))
			if(len(strVal)%2==0){
				first,_:= strconv.Atoi(strVal[0:len(strVal)/2])
				second,_ := strconv.Atoi(strVal[(len(strVal)/2):])
				e.Value=first
				stones.InsertAfter(second,e)
				e=e.Next()
			}else{
				e.Value = e.Value.(int)*2024
			}
		}
	}
}
//memoize the results so we're only calculating each stone at a particular depth once.
func calcDP(stone int, depth int, dp map[int]map[int]int) int {
	if(depth==0){
		return 1
	}
	mapVal,numExists := dp[stone]
	if numExists {
		mapMapVal, existsAtDepth := mapVal[depth]
		if existsAtDepth{
			return mapMapVal
		}
	}else{
		dp[stone] = make(map[int]int)
	}
	if(stone==0){
		dp[stone][depth] = calcDP(1,depth-1,dp)
		return dp[stone][depth]
	} else {
		strVal := strconv.Itoa(stone)
		if(len(strVal)%2==0){
			first,_:= strconv.Atoi(strVal[0:len(strVal)/2])
			second,_ := strconv.Atoi(strVal[(len(strVal)/2):])
			dp[stone][depth] = calcDP(first,depth-1,dp)+calcDP(second,depth-1,dp)
			return dp[stone][depth]
		}else{
			dp[stone][depth] = calcDP(stone*2024,depth-1,dp)
			return dp[stone][depth]
		}
	}
}

func main(){
	dat, err := os.ReadFile("../../inputs/day11.txt")
	check(err)
	stonesStr := strings.Split(string(dat)," ")
	stones := list.New()
	totalAt75:=0
	totalAt25:=0
	for _,val := range stonesStr{
		num,err := strconv.Atoi(val)
		check(err)
		stones.PushBack(num)
		totalAt75+=calcDP(num,75,make(map[int]map[int]int))
		totalAt25+=calcDP(num,25,make(map[int]map[int]int))
	}

	fmt.Println("Total at 75 (part 2):",totalAt75)
	fmt.Println("Total at 25 (part 1):",totalAt25)
	
	//old method, actually splits the stones and tracks the entire list.
	// for i:=0;i<75;i++{
	// 	iterateRules(stones)
	// 	uniqueNumbers := make(map[int]bool)
	// 	for e := stones.Front(); e != nil; e = e.Next() {
	// 		uniqueNumbers[e.Value.(int)] = true
	// 	}
	// 	fmt.Println(i,",",len(uniqueNumbers))
	// 	fmt.Println(stones.Len())
	// }
}