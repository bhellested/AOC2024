package main

import (
	"fmt"
	"strconv"
	"os"
	"regexp"
	"strings"
)

var (
	mulRegex = regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	numRegex = regexp.MustCompile(`[0-9]+`)
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func calculateMultiplications(slice string) uint64 {
	res := mulRegex.FindAllString(slice,-1)
	var result uint64 = 0;
	for _, value := range res {
		nums := numRegex.FindAllString(value,2)
		first,_ :=strconv.ParseInt(nums[0],10,64)
		second,_ := strconv.ParseInt(nums[1],10,64)
		result+=uint64(first)*uint64(second)
	} 
	return result
}

func main() {
	dat, err := os.ReadFile("../../inputs/day3.txt")
	check(err)
	input := string(dat)
	result := uint64(0)
	isDo := true 

	for {
		if isDo {
			//find up until the next don't, add these to result
			index := strings.Index(input,`don't()`)
			if index != -1 {
				result += calculateMultiplications(input[0:index])
			}else{
				//we're done
				result += calculateMultiplications(input)
				break
			}
			input = input[index+4:]
			isDo = false
		}else{
			//find up until the next do, skip these
			index := strings.Index(input,`do()`)
			if index == -1 {
				break
			}
			input = input[index+2:]
			isDo = true
		}
	}

	fmt.Println("Result is: ", result)
}