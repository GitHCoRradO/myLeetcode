package main

import (
	"fmt"
	"math"
	"os"
)

/*
AA
1 * 26^1 + 1 * 26^0
AZ
1 * 26^1 + 26 * 26^0
 */

func main() {
	fmt.Println(titleToNumber(os.Args[1]))
}

func titleToNumber(columnTitle string) int {
	var sum int64 = 0
	for i, c := range columnTitle {
		fmt.Println(c)
		fmt.Println(int64(math.Pow(26.0, float64(len(columnTitle)- i - 1))))
		sum += int64(c - 64) * int64(math.Pow(26.0, float64(len(columnTitle)- i - 1)))
	}
	return int(sum)
}

