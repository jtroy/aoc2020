package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type password struct {
	n1  int
	n2  int
	let byte
	str string
}

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

func parseInput(lines []string) []password {
	res := make([]password, len(lines))
	for i, line := range lines {
		dash := strings.Index(line, "-")
		n1, _ := strconv.Atoi(line[0:dash])
		space := strings.Index(line[dash:], " ")
		n2, _ := strconv.Atoi(line[dash+1 : dash+space])
		res[i] = password{
			n1, n2,
			line[dash+space+1],
			line[dash+space+4:],
		}
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
	valid := 0
	for _, password := range input {
		c := strings.Count(password.str, string(password.let))
		if c >= password.n1 && c <= password.n2 {
			valid++
		}
	}
	return valid, time.Since(start)
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	input := parseInput(lines)
	valid := 0
	for _, password := range input {
		// Poor man's XOR
		if (password.str[password.n1-1] != password.let) != (password.str[password.n2-1] != password.let) {
			valid++
		}
	}
	return valid, time.Since(start)
}
