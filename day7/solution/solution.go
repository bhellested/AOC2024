package main 

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	
)

type Equation struct {
	solution uint64
	nums []uint64
	canBeSolved bool
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func parseEquation(line string) Equation {
	solutionStr := line[0:strings.Index(line,":")]
	solutionVal,_ := strconv.ParseUint(solutionStr, 10, 64)
	operators := line[strings.Index(line,":")+2:]
	operatorsSplit := strings.Split(operators, " ")
	var nums []uint64
	for _,num := range operatorsSplit {
		val,_ := strconv.Atoi(num)
		nums = append(nums, uint64(val))
	}
	return Equation{solution: solutionVal, nums: nums, canBeSolved: false}
}

func recurse(equation Equation, itemsLeft []uint64,checkConcat bool) bool {
	if len(itemsLeft) == 1 {
		return itemsLeft[0] == equation.solution
	}

	ifAdded := itemsLeft[0] +itemsLeft[1]
	deepCopyAdded := make([]uint64, len(itemsLeft)-1)
	copy(deepCopyAdded, itemsLeft[1:])
	deepCopyAdded[0] = ifAdded

	ifMultiplied := itemsLeft[0] * itemsLeft[1]
	deepCopyMultiplied := make([]uint64, len(itemsLeft)-1)
	copy(deepCopyMultiplied, itemsLeft[1:])
	deepCopyMultiplied[0] = ifMultiplied

	if !checkConcat {
		return recurse(equation, deepCopyAdded,false) || recurse(equation, deepCopyMultiplied,false)
	}
	ifConcatinated := strconv.FormatUint(itemsLeft[0], 10) + strconv.FormatUint(itemsLeft[1], 10)
	ifConcatinatedVal,_ := strconv.ParseUint(ifConcatinated, 10, 64)
	deepCopyConcatinated := make([]uint64, len(itemsLeft)-1)
	copy(deepCopyConcatinated, itemsLeft[1:])
	deepCopyConcatinated[0] = ifConcatinatedVal

	return recurse(equation, deepCopyAdded,true) || recurse(equation, deepCopyMultiplied,true) || recurse(equation, deepCopyConcatinated,true)
}

func main() {
	dat, err := os.ReadFile("../../inputs/day7.txt")
	check(err)
	lines:=strings.Split(string(dat), "\n")
	equations := make([]Equation, len(lines))
	var sumPart1 uint64 = 0
	var sumPart2 uint64 = 0
	for i,line := range lines {
		equations[i] = parseEquation(line)
		//maxNum = max(maxNum, len(equations[i].nums))
		if recurse(equations[i], equations[i].nums,false) {
			sumPart1 += equations[i].solution
		}
		if recurse(equations[i], equations[i].nums,true) {
			sumPart2 += equations[i].solution
		}
	}

	fmt.Println("Part 1:", sumPart1)
	fmt.Println("Part 2:", sumPart2)
}