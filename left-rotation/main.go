package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Complete the rotLeft function below.
// https://www.hackerrank.com/challenges/ctci-array-left-rotation
func rotLeft(a []int32, d int32) []int32 {

	size := int32(len(a))

	if size > d {
		size = d % size
	}

	log.Println("Initial array [", len(a), "]: ", a)
	log.Println("Shifting array", d, "symbols to the right")
	first := a[size:]
	log.Println("First part: ", first)
	second := a[:size]
	log.Println("Second part: ", second)
	result := append(first, second...)
	log.Println("Result: ", result)

	return result
}

func main() {
	file, _ := os.Open("left-rotation/input.txt")
	reader := bufio.NewReaderSize(file, 1024*1024)

	stdout, err := os.Create("left-rotation/output.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
