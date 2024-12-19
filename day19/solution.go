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

func canMakeTowel(towel string,available map[string]struct{},cache map[string]int) int{
	if len(towel) == 0{
		return 1
	}
	count := 0
	if cached,exists := cache[towel];exists{
		return cached
	}
	for i:=1; i<=len(towel);i++{
		currentSubstr := towel[:i]
		_,exists := available[currentSubstr]
		if(exists){
			count+=canMakeTowel(towel[i:],available,cache)
		}
	}
	cache[towel] = count
	return count
}

func main() {
	input, err := os.ReadFile("../inputs/day19.txt")
	check(err)
	lines := strings.Split(string(input), "\n\n")
	availableTowels := strings.Split(lines[0], ", ")
	availableTowelsMap := make(map[string]struct{})
	cache := make(map[string]int)
	for _, towel := range availableTowels {
		availableTowelsMap[towel] = struct{}{}
	}
	towelOrders := strings.Split(lines[1], "\n")
	totalCount := 0
	numCanMake := 0
	for _, towelOrder := range towelOrders {
		res:=canMakeTowel(towelOrder,availableTowelsMap,cache)
		if res >0 {
			numCanMake++
			totalCount+=res
		}
	}
	fmt.Println("we can make:",numCanMake,"of the desired combinations")
	fmt.Println("We can make", totalCount, " unique towel combinations")
}