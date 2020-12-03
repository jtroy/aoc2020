package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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
	return checkSlope(lines, 3, 1), time.Since(start)
}

func checkSlope(input []string, right, down int) int {
	trees := 0
	for x, y := 0, 0; y < len(input); x, y = x+right, y+down {
		line := input[y]
		idx := x % len(line)
		if line[idx] == '#' {
			trees++
		}
	}
	return trees
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	return checkSlope(lines, 1, 1) * checkSlope(lines, 3, 1) * checkSlope(lines, 5, 1) * checkSlope(lines, 7, 1) * checkSlope(lines, 1, 2), time.Since(start)
}
