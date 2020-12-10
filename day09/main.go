package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func readInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	var (
		rounds        int
		inputFilename string
	)
	flag.IntVar(&rounds, "rounds", 1, "number of rounds for benchmark")
	flag.StringVar(&inputFilename, "input", "input", "input data filename")
	flag.Parse()

	input, err := readInput(inputFilename)
	if err != nil {
		panic(err)
	}
	// PART 1
	fmt.Println("=== Part 1 ===")
	doTimes := make([]time.Duration, rounds)
	var res int64
	for i := 0; i < rounds; i++ {
		res, doTimes[i] = doPart1(input)
	}
	var tot time.Duration
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))

	// PART 2
	fmt.Println("\n=== Part 2 ===")
	invalid := res
	for i := 0; i < rounds; i++ {
		res, doTimes[i] = doPart2(input, invalid)
	}
	tot = 0
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
}

func stringSliceToInt64Slice(lines []string) []int64 {
	res := make([]int64, len(lines))
	for i, line := range lines {
		res[i], _ = strconv.ParseInt(line, 10, 64)
	}
	return res
}

func doPart1(lines []string) (int64, time.Duration) {
	start := time.Now()
	input := stringSliceToInt64Slice(lines)
OUTER:
	for i := 25; i < len(input); i++ {
		n := input[i]
		for j := i - 1; j >= i-25; j-- {
			search := n - input[j]
			for k := j - 1; k >= i-25; k-- {
				if input[k] == search {
					// found two numbers in the last 25 lines which add to n
					continue OUTER
				}
			}
		}
		return n, time.Since(start)
	}
	return -1, time.Since(start)
}

func doPart2(lines []string, invalid int64) (int64, time.Duration) {
	start := time.Now()
	input := stringSliceToInt64Slice(lines)
	for i := 0; i < len(input); i++ {
		var (
			sum      int64
			smallest = int64(math.MaxInt64)
			largest  = int64(math.MinInt64)
		)
		for j := i; j < len(input); j++ {
			if input[j] < smallest {
				smallest = input[j]
			}
			if input[j] > largest {
				largest = input[j]
			}
			sum += input[j]
			if sum >= invalid {
				break
			}
		}
		if sum == invalid {
			return smallest + largest, time.Since(start)
		}
	}
	return -1, time.Since(start)
}
