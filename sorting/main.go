package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
}

/*
	4 8 2 =  14
	4 5 7 =  16
	6 1 6 =  13
 16 14 14 15 15

 	4 8 2 =  14
	4 5 7 =  16
	6 1 6 =  13
 16 14 14 15 15

*/
/*
 * Complete the towerBreakers function below.
 */
func towerBreakers(arr []int32) int32 {
	/*
	 * Write your code here.
	 */

	var count int64 = 0
	for i := 0; i < len(arr); i++ {
		// Count the number of times the numbers in the above array can be divided by 2
		// If the count is even then player 2 wins
		// if the count is odd then the player 1 wins
		if arr[i]%2 != 0 {
			continue
		}
		log2 := int64(math.Floor(math.Log2(float64(arr[i]))))
		fmt.Println(arr[i], log2)
		count += log2
	}
	if count%2 == 0 {
		return 2
	} else {
		return 1
	}

	// 1 4 6 8 7 5

}

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {
	sort.SliceStable(ar, func(i, j int) bool { return ar[i] < ar[j] })
	tallest := ar[len(ar)-1]
	var count int32 = 0
	for i := len(ar) - 1; i >= 0; i-- {
		if tallest == ar[i] {
			count++
			continue
		}
		break
	}
	return count
}

func minMax(num []int32) {
	sort.SliceStable(num, func(i, j int) bool { return num[i] < num[j] })
	var sum int64 = 0
	for i := 0; i < len(num); i++ {
		sum = sum + int64(num[i])
	}
	max := sum - int64(num[0])
	min := sum - int64(num[len(num)-1])
	fmt.Println(min, max)
}
