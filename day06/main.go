package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strings"
	"time"
)

func scanRecords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}

func readInput(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(scanRecords)
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

func doPart1(groups []string) (int, time.Duration) {
	start := time.Now()
	var (
		total int
	)
	for _, group := range groups {
		var answers uint32
		for _, c := range group {
			if c == '\n' {
				continue
			}
			answers |= 1 << uint32(c-'a')
		}
		total += bits.OnesCount32(answers)
	}
	return total, time.Since(start)
}

func doPart2(groups []string) (int, time.Duration) {
	start := time.Now()
	total := 0
	for _, group := range groups {
		answers := uint32(math.MaxUint32)
		for _, decl := range strings.Fields(group) {
			var answer uint32
			for _, c := range decl {
				answer |= 1 << uint32(c-'a')
			}
			answers &= answer
		}
		total += bits.OnesCount32(answers)
	}
	return total, time.Since(start)
}
