package main

import (
	//"bufio"
	"fmt"
	"strconv"
	//"io"
	"os"
	"regexp"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	dat, err := os.ReadFile("../../inputs/day3.txt")
    check(err)

	input := string(dat)
	re := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
	re2 := regexp.MustCompile("[0-9]+")
    //fmt.Print(string(dat))
	res := re.FindAllString(input,-1)
	fmt.Print(res)
	var result uint64 =0
	for _, value := range res {
		nums := re2.FindAllString(value,2)
		first,_ :=strconv.ParseInt(nums[0],10,64)
		second,_ := strconv.ParseInt(nums[1],10,64)
		result+=uint64(first)*uint64(second)
	} 
	fmt.Println("Result is: ", result)
}