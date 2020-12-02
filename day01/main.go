package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
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

func parseInput(lines []string) []int {
	res := make([]int, len(lines))
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		res[i] = n
	}
	return res
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
	input := parseInput(lines)
	sort.IntSlice(input).Sort()
	for i := len(input) - 1; i >= 0; i-- {
		m := sort.SearchInts(input, 2020-input[i])
		if m == len(input) {
			continue
		}
		if input[m]+input[i] == 2020 {
			return input[m] * input[i], time.Since(start)
		}
	}
	return -1, time.Since(start)
}

func trimImpossible(input []int) []int {
	smallestAddend := input[0] + input[1]
	i := len(input) - 1
	for ; i > 1 && input[i]+smallestAddend > 2020; i-- {
	}
	return input[0:i]
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	input := parseInput(lines)
	sort.IntSlice(input).Sort()
	trimmed := trimImpossible(input)
	for i := 0; i < len(trimmed); i++ {
		for j := i + 1; j < len(trimmed); j++ {
			for k := j + 1; k < len(trimmed); k++ {
				if trimmed[i]+trimmed[j]+trimmed[k] == 2020 {
					return trimmed[i] * trimmed[j] * trimmed[k], time.Since(start)
				}
			}
		}
	}
	return -1, time.Since(start)
}
