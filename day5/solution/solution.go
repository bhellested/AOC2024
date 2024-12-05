package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)	

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func CheckForValidity(update []int, rulesMap map[int]map[int]struct{}) int64 {
	//first make a set of all ints in update, mapping value to index in array
	updateSet := make(map[int]int)
	for i,num := range update {
		updateSet[num] = i
	}
	
	for key,valueMap := range rulesMap {
		//key must come before all items in valuemap iff both exist
		keyPosition,ok := updateSet[key]
		if ok {
			//key exists
			for value := range valueMap {
				valuePosition,ok := updateSet[value]
				if(ok){
					//now make sure that the key comes before the value. If not, we early exit
					if keyPosition > valuePosition {
						return 0
					}
				}
			}
		}
	}
	//if we've made it here we have passed, return value in the center of update
	return int64(update[len(update)/2])
}

func ReorderAndFindMedian(update []int, rulesMap map[int]map[int]struct{}) int64{
	//first make a set of all ints in update, mapping value to index in array
	updateSet := make(map[int]int)
	for i,num := range update {
		updateSet[num] = i
	}
	//basic swap every time we find something wrong, until we don't find anything wrong
	for{
		foundError := false
		for key,valueMap := range rulesMap {
			//key must come before all items in valuemap iff both exist
			keyPosition,ok := updateSet[key]
			if ok {
				//key exists
				for value := range valueMap {
					valuePosition,ok := updateSet[value]
					if(ok){
						//now make sure that the key comes before the value. If not, we early exit
						if keyPosition > valuePosition {
							foundError=true
							//swap these items in the update list
							temp:=update[keyPosition]
							update[keyPosition] = update[valuePosition]
							update[valuePosition] = temp
							//update values in the set that tracks indices
							updateSet[key] = valuePosition
							updateSet[value] = keyPosition
							break;
						}
					}
				}
			}
		}
		if !foundError{
			break
		}
	}
	return int64(update[len(update)/2])
}

func main() {
	dat, err := os.ReadFile("../../inputs/day5.txt")
    check(err)

	lines:=strings.Split(string(dat), "\n\n")
	rules, updates := lines[0],lines[1]
	rulesMap := make(map[int]map[int]struct{})

	for _,line := range strings.Split(rules, "\n"){
		parts := strings.Split(line, "|")
		before, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		after, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		if rulesMap[before] == nil {
			rulesMap[before] = make(map[int]struct{})
		}
		rulesMap[before][after]=struct{}{}
	}
	var sumOfValidUpdates int64 = 0
	var sumOfInvalidUpdates int64 = 0
	for _,updateString := range strings.Split(updates,"\n") {
		splitUpdate := strings.Split(updateString, ",")
		update :=make([]int,len(splitUpdate))
		for i,str := range splitUpdate {
			update[i],_ = strconv.Atoi(str)
		}
		val := CheckForValidity(update,rulesMap)
		if val != 0 {
			sumOfValidUpdates += val
		} else {
			sumOfInvalidUpdates += ReorderAndFindMedian(update,rulesMap)

		}
	}
	fmt.Println("Sum of Valid (part1): ",sumOfValidUpdates)
	fmt.Println("Sum of reordered Invalid (part2)",sumOfInvalidUpdates )
}