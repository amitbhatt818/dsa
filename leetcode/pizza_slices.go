package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'pizzaSlices' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func pizzaSlices(arr []int32) int32 {
	l := int32(len(arr))
	helper := func(start, end int32) int32 {
		dp := [2][]int32{}
		var i, j, k int32
		for k = 0; k < 2; k++ {
			dp[k] = make([]int32, l/3+1)
		}
		for i = start; i < end; i++ {
			for j = int32(math.Min(float64(l/3), float64(i+2-start)/2)); j >= 1; j-- {
				dp[(i+1)%2][j] = int32(math.Max(float64(dp[i%2][j]), float64(dp[(i+1)%2][j-1]+arr[i])))
			}
		}
		return int32(math.Max(float64(dp[0][l/3]), float64(dp[1][l/3])))
	}
	return int32(math.Max(float64(helper(0, l-1)), float64(helper(1, l))))
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := pizzaSlices(arr)

	fmt.Printf("%d\n", result)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
