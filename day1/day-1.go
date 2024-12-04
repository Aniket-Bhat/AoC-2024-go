package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := getInput("input.txt")
	l1 := input[0]
	l2 := input[1]
	l1 = quickSort(l1)
	l2 = quickSort(l2)
	// part a solution
	answerA := 0
	for i := 0; i < len(l1); i++ {
		answerA += absInt(l1[i] - l2[i])
	}

	l1 = input[0]
	l2 = input[1]

	// part b solution
	freq := make(map[int]int)
	answerB := 0
	for _, value := range l2 {
		freq[value]++
	}

	for _, value := range l1 {
		answerB += value * freq[value]
	}

	fmt.Printf("Answer to part A is %d ", answerA)
	fmt.Println("")
	fmt.Println("----------------------")
	fmt.Printf("Answer to part B is %d", answerB)
}

func getInput(filename string) [2][]int {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic("error reading file")
	}
	str := string(content)
	strArr := strings.Split(str, "\n")

	var l1 []int
	var l2 []int
	for _, value := range strArr {
		valueArr := strings.Split(value, "   ")
		n1, err := strconv.Atoi(valueArr[0])
		if err != nil {
			fmt.Println("error reading idx 0")
		}

		n2, err := strconv.Atoi(valueArr[1])
		if err != nil {
			fmt.Println("error reading idx 1")
		}

		l1 = append(l1, n1)
		l2 = append(l2, n2)
	}
	result := [...][]int{l1, l2}
	return result
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func quickSort(arr []int) []int {
	if len(arr) < 1 {
		return arr
	}
	var left []int
	var right []int
	pivot := arr[0]
	for _, value := range arr[1:] {
		if value >= pivot {
			right = append(right, value)
		} else {
			left = append(left, value)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
