package main

import "fmt"

func main() {
	memo := map[int]int{}
	fmt.Println("Recurrsion without Memoization")
	fmt.Println(fibo(8))
	fmt.Println("Recurrsion with Memoization")
	fmt.Println(fiboMemoized(8, memo))

	for i := 0; i <= 8; i++ {
		fmt.Println(fiboMemoized(i, memo))
	}

}

// fibo(n) = fibo(n-1) + fibo(n-2)
// 0, 1, 1, 2, 3, 5, 8, 13 ....
func fiboMemoized(n int, memo map[int]int) int {
	fmt.Println("Calling fibo for ", n)
	if n == 0 {
		memo[0] = 0
		return 0
	}
	if n == 1 {
		memo[1] = 1
		return 1
	}

	_, status := memo[n-1]
	if !status {
		memo[n-1] = fiboMemoized(n-1, memo)
	} else {
		fmt.Println("Using Memoized value for n-1=", n-1)
	}
	_, status2 := memo[n-2]
	if !status2 {
		memo[n-2] = fiboMemoized(n-2, memo)
	} else {
		fmt.Println("Using Memoized value for n-2=", n-2)
	}
	memo[n] = memo[n-1] + memo[n-2]
	return memo[n]
}

func fibo(n int) int {
	fmt.Println("Calling fibo for ", n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}
