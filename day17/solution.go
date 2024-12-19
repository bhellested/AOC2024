package main

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func getOutput(registerA int, instructions []int) []int {
	registerB := 0
	registerC := 0
	programCounter := 0
	output := []int{}
	for programCounter < len(instructions) {
		instruction := instructions[programCounter]
		operand := instructions[programCounter+1]
		comboOperand := operand
		//grab the right combo operand
		switch operand {
		case 4:
			comboOperand = registerA
		case 5:
			comboOperand = registerB
		case 6:
			comboOperand = registerC
		}

		switch instruction {
		case 0:
			result := int(float64(registerA) / math.Pow(2, float64(comboOperand)))
			registerA = result
			programCounter += 2
		case 1:
			result := registerB ^ operand
			registerB = result
			programCounter += 2
		case 2:
			registerB = comboOperand % 8
			programCounter += 2
		case 3:
			if registerA != 0 {
				programCounter = operand
			} else {
				programCounter += 2
			}
		case 4:
			registerB = registerB ^ registerC
			programCounter += 2
		case 5:
			result := comboOperand % 8
			output = append(output, result)
			programCounter += 2
		case 6:
			result := int(float64(registerA) / math.Pow(2, float64(comboOperand)))
			registerB = result
			programCounter += 2
		case 7:
			result := int(float64(registerA) / math.Pow(2, float64(comboOperand)))
			registerC = result
			programCounter += 2
		}
	}

	return output
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Execution time: %d microseconds\n", time.Since(start).Microseconds())
	}()

	input, err := os.ReadFile("../inputs/day17.txt")
	check(err)
	lines := strings.Split(string(input), "\n")
	registerA, err := strconv.Atoi(strings.Split(lines[0], ":")[1][1:])
	check(err)
	programString := strings.Split(lines[4], ":")[1][1:]
	program := strings.Split(programString, ",")
	instructions := make([]int, len(program))
	for i, instruction := range program {
		instructions[i], err = strconv.Atoi(instruction)
		check(err)
	}
	output := getOutput(registerA, instructions)
	fmt.Println("Part 1 (make sure to add the ',' inbetween the numbers): ", output[:len(output)-1])
	possibleResults := []int{0}
	for i := len(instructions) - 1; i >= 0; i-- {
		mustMatch := instructions[i:]
		newPossibleResults := []int{}
		for j := 0; j < 8; j++ {
			for _, result := range possibleResults {
				output := getOutput(result<<3+j, instructions)
				if reflect.DeepEqual(output, mustMatch) {
					newPossibleResults = append(newPossibleResults, result<<3+j)
				}
			}
		}
		possibleResults = newPossibleResults
	}

	smallest := math.MaxInt
	for _, result := range possibleResults {
		if result < smallest {
			smallest = result
		}
	}
	fmt.Println("Smallest (part2): ", smallest)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
