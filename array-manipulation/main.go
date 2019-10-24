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

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {

	result := make([]int64, n, n)

	for _, q := range queries {
		plus := int64(q[2])

		for i := q[0] - 1; i <= q[1]-1; i++ {
			result[i] += plus
		}
	}

	var max int64 = math.MinInt64
	for _, r := range result {
		if r > max {
			max = r
		}
	}

	return max
}

func main() {

	file, err := os.Open("array-manipulation/input.txt")
	checkError(err)
	reader := bufio.NewReaderSize(file, 1024*1024)

	stdout, err := os.Create("array-manipulation/output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

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
