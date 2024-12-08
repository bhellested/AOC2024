package main 

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"time"
	
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
	}else if itemsLeft[0] > equation.solution {
		return false
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
	}

	start := time.Now()
	for _,equation := range equations {
		if recurse(equation, equation.nums,false) {
			sumPart1 += equation.solution
		}
	}
	elapsedPart1 := time.Since(start)
	
	start = time.Now()
	for _,equation := range equations {
		if recurse(equation, equation.nums,true) {
			sumPart2 += equation.solution
		}
	}
	elapsedPart2 := time.Since(start)
	fmt.Println("Part 1:", sumPart1, "(took", elapsedPart1, ")")
	fmt.Println("Part 2:", sumPart2, "(took", elapsedPart2, ")")
}