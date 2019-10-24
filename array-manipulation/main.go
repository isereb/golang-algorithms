package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
// https://www.hackerrank.com/challenges/crush/problem
func arrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n+1)

	for _, e := range queries {
		a := int(e[0])
		b := e[1]
		k := int64(e[2])
		arr[a] += k
		if (b + 1) <= n {
			arr[b+1] -= k
		}
	}

	var x, max int64 = 0, 0
	for i := 1; i <= int(n); i++ {
		x += arr[i]
		if max < x {
			max = x
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
