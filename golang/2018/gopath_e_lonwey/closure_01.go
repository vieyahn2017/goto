package main

import (
	"fmt"
	)

func main() {
	s := make([]int ,10)
	fmt.Println(s)

	for i, _ := range s {
		s[i] = i
	}
	fmt.Println(s)
	doubleSum := sum(isDouble)
	singleSum := sum(isSingle)
	fmt.Println(doubleSum(s))
	fmt.Println(singleSum(s))
}

func isDouble(a int) bool {
	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func isSingle(a int) bool {
	if a%2 == 0 {
		return false
	} else {
		return true
	}
}

func sum(f func(int) bool) func([] int) int {
	return func(a [] int) int {
		sum := 0
		for _, v := range a {
			if f(v) {
				sum += v
			}
		}
		return sum
	}
}