package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Failed to open!")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}

func solutionOne(text []string) int {
	var counter, last int = 0, 0
	for _, each_ln := range text {
		val, _ := strconv.Atoi(each_ln)
		if val > last {
			counter++
		}
		last = val
	}
	return counter - 1
}

func solutionTwo(text []string) int {
	var counter, lastSum int = 0, 0
	var inputSize = len(text)

	for i := 0; i < inputSize-2; i++ {
		sum := 0
		for j := 0; j < 3; j++ {
			index := i + j
			val, _ := strconv.Atoi(text[index])
			sum += val
		}
		if sum > lastSum {
			counter++
		}
		lastSum = sum
	}

	return counter - 1
}

func main() {
	text := getInput()
	fmt.Println("Answer 1:", solutionOne(text))
	fmt.Println("Answer 2:", solutionTwo(text))
}
