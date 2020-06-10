package main

import (
	"fmt"
)

func main() {
	fmt.Println("Factorial With Recurrsion")
	for i := 0; i <= 5; i++ {
		fmt.Println(fact(i))
	}
	fmt.Println("Factorial with Recurrsion and Memoization")
	memo := map[int]int{}
	for i := 0; i <= 5; i++ {
		fmt.Println(factMemoized(i, memo))
	}

}

func fact(n int) int {
	fmt.Println("Calling fact for", n)
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func factMemoized(n int, memo map[int]int) int {
	fmt.Println("Calling fact for", n)
	if n == 0 {
		memo[0] = 1
		return 1
	}
	_, status := memo[n-1]
	if !status {
		memo[n-1] = factMemoized(n-1, memo)
	}
	memo[n] = n * memo[n-1]
	return memo[n]
}
