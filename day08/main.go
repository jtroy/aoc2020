package main

import (
	"bufio"
	"flag"
	"fmt"
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

type machine struct {
	pc           int   // program counter
	a            int   // accumulator
	cycle        int   // starts at 1 to make visitedCycle test easier
	visitedCycle []int // cycle at which instruction was visited -- 0 means not visited
}

// runMachine runs over the input until a loop occurs or pc is out of bounds.
// It returns the accumulator and program counter at the time the program
// terminated.
func runMachine(lines []string) (int, int) {
	m := machine{0, 0, 1, make([]int, len(lines))}
	for m.pc < len(lines) && m.pc >= 0 {
		m.visitedCycle[m.pc] = m.cycle
		m.cycle++
		line := lines[m.pc]
		op, arg := line[0:3], line[4:]
		switch op {
		case "nop":
		case "acc":
			argtoi, _ := strconv.Atoi(arg)
			m.a += argtoi
		case "jmp":
			addr, _ := strconv.Atoi(arg)
			m.pc += addr
			if m.pc == len(lines) || m.visitedCycle[m.pc] > 0 {
				return m.a, m.pc
			}
			continue
		}
		m.pc++
	}
	return m.a, m.pc
}

func doPart1(lines []string) (int, time.Duration) {
	start := time.Now()
	a, _ := runMachine(lines)
	return a, time.Since(start)
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	// loop through the input, swapping nop and jmp at each relevant line
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		op := line[0:3]
		if op == "acc" {
			continue
		}
		linesCopy := make([]string, len(lines))
		copy(linesCopy, lines[0:i])
		switch op {
		case "nop":
			linesCopy[i] = "jmp" + line[3:]
		case "jmp":
			linesCopy[i] = "nop" + line[3:]
		}
		copy(linesCopy[i+1:], lines[i+1:])
		a, pc := runMachine(linesCopy)
		if pc == len(lines) {
			return a, time.Since(start)
		}
	}
	return -1, time.Since(start)
}
