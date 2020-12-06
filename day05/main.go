package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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
	var res int
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
	for i := 0; i < rounds; i++ {
		res, doTimes[i] = doPart2(input)
	}
	tot = 0
	for i := 0; i < rounds; i++ {
		tot += doTimes[i]
	}
	fmt.Printf("result:\t %v\n", res)
	fmt.Printf("elapsed:\t%v\n", tot/time.Duration(rounds))
}

func doPart1(lines []string) (int, time.Duration) {
	start := time.Now()
	max := 0
	for _, line := range lines {
		value := 0
		for _, c := range line {
			if c == 'B' || c == 'R' {
				value = (value << 1) | 1
			} else {
				value <<= 1
			}
		}
		if value > max {
			max = value
		}
	}
	return max, time.Since(start)
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	seats := make([]int, len(lines))
	for i, line := range lines {
		value := 0
		for _, c := range line {
			if c == 'B' || c == 'R' {
				value = (value << 1) | 1
			} else {
				value <<= 1
			}
		}
		seats[i] = value
	}
	sort.Ints(seats)
	last := -1
	for _, seat := range seats {
		if seat-last == 2 {
			return seat - 1, time.Since(start)
		}
		last = seat
	}
	return -1, time.Since(start)
}
