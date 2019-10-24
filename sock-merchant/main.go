package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/sock-merchant/

type sock struct {
	number      int32
	occurrences int32
}

func (pair *sock) increment() {
	pair.occurrences = pair.occurrences + 1
}

func (pair *sock) getAmountOfPairs() int32 {
	return pair.occurrences / 2
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {

	pairs := make([]*sock, 0, 10)

	for _, e := range ar {

		found := false
		for _, pair := range pairs {
			if pair.number == e {
				pair.increment()
				found = true
			}
		}

		if !found {
			pairs = append(pairs, &sock{
				number:      e,
				occurrences: 1,
			})
		}
	}

	var amountOfPairs int32 = 0
	for _, pair := range pairs {
		amountOfPairs += pair.getAmountOfPairs()
	}

	return amountOfPairs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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
