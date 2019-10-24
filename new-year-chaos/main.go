package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type person struct {
	order   int32
	swapped uint8
}

func (p *person) canSwap() bool {
	return p.swapped != 2
}

func (p *person) swap() {
	if p.canSwap() {
		p.swapped++
	} else {
		panic("Cannot swap. Person already swapped 2 times.")
	}
}

// Complete the minimumBribes function below.
// https://www.hackerrank.com/challenges/new-year-chaos
func minimumBribes(q []int32) {

	length := len(q)

	people := make([]person, 0, length)
	for _, p := range q {
		people = append(people, person{p, 0})
	}

	keepGoing := true
	for keepGoing {
		for i := 0; i < length-1; i++ {

			if q[i]-int32(i+1) > 2 {
				fmt.Println("Too chaotic")
				return
			}

			if q[i] > q[i+1] && people[i].canSwap() {
				people[i].swap()
				q[i], q[i+1] = q[i+1], q[i]
			}
		}

		keepGoing = false
		for i := 0; i < length-1; i++ {
			needSwap := q[i] > q[i+1]
			canSwap := people[i].canSwap()

			if needSwap && canSwap {
				keepGoing = true
				break
			}
		}
	}

	var swaps int32 = 0
	for _, p := range people {
		swaps += int32(p.swapped)
	}

	fmt.Println(swaps)
}

func main() {

	file, err := os.Open("new-year-chaos/input.txt")
	checkError(err)
	reader := bufio.NewReaderSize(file, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
