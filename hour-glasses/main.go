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

// Complete the hourglassSum function below.
// https://www.hackerrank.com/challenges/2d-array/problem
func hourglassSum(arr [][]int32) int32 {

	rows := len(arr)
	columns := len(arr[0])

	var maxSum int32 = math.MinInt32
	for i := 0; i < rows-2; i++ {
		for j := 0; j < columns-2; j++ {
			a1 := arr[i][j]
			a2 := arr[i][j+1]
			a3 := arr[i][j+2]

			b2 := arr[i+1][j+1]

			c1 := arr[i+2][j]
			c2 := arr[i+2][j+1]
			c3 := arr[i+2][j+2]

			sum := a1 + a2 + a3 + b2 + c1 + c2 + c3

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum

}

func main() {
	file, _ := os.Open("hour-glasses/input.txt")
	reader := bufio.NewReaderSize(file, 1024*1024)

	stdout, err := os.Create("hour-glasses/output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
