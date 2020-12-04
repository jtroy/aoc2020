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
	var (
		fields int
		hasCid bool
		count  int
	)
	for _, line := range lines {
		if len(line) == 0 {
			if fields == 8 || (fields == 7 && !hasCid) {
				count++
			}
			fields = 0
			hasCid = false
			continue
		}
		fields += strings.Count(line, ":")
		hasCid = hasCid || strings.Contains(line, "cid")
	}
	if fields == 8 || (fields == 7 && !hasCid) {
		count++
	}
	return count, time.Since(start)
}

func validLine(line string) bool {
	for _, field := range strings.Fields(line) {
		col := strings.Index(field, ":")
		val := field[col+1:]
		switch field[0:col] {
		case "byr":
			if byr, err := strconv.Atoi(val); err != nil || !(byr >= 1920 && byr <= 2002) {
				return false
			}
		case "iyr":
			if iyr, err := strconv.Atoi(val); err != nil || !(iyr >= 2010 && iyr <= 2020) {
				return false
			}
		case "eyr":
			if eyr, err := strconv.Atoi(val); err != nil || !(eyr >= 2020 && eyr <= 2030) {
				return false
			}
		case "hgt":
			if strings.HasSuffix(val, "cm") {
				if hgt, err := strconv.Atoi(val[:len(val)-2]); err != nil || !(hgt >= 150 && hgt <= 193) {
					return false
				}
			} else if strings.HasSuffix(val, "in") {
				if hgt, err := strconv.Atoi(val[:len(val)-2]); err != nil || !(hgt >= 59 && hgt <= 76) {
					return false
				}
			} else {
				return false
			}
		case "hcl":
			if val[0] != '#' {
				return false
			}
			if len(val) != 7 {
				return false
			}
			if _, err := strconv.ParseUint(val[1:], 16, 24); err != nil {
				return false
			}
		case "ecl":
			if len(val) != 3 || !strings.Contains("amb blu brn gry grn hzl oth", val) {
				return false
			}
		case "pid":
			if len(val) != 9 {
				return false
			}
			if _, err := strconv.Atoi(val); err != nil {
				return false
			}
			//case "cid":
		}
	}
	return true
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	var (
		fields  int
		hasCid  bool
		isValid = true
		count   int
	)
	for _, line := range lines {
		if len(line) == 0 {
			if isValid && (fields == 8 || (fields == 7 && !hasCid)) {
				count++
			}
			fields = 0
			hasCid = false
			isValid = true
			continue
		}
		if !isValid {
			continue
		}
		if !validLine(line) {
			isValid = false
			//continue
		}
		fields += strings.Count(line, ":")
		hasCid = hasCid || strings.Contains(line, "cid")
	}
	if isValid && (fields == 8 || (fields == 7 && !hasCid)) {
		count++
	}
	return count, time.Since(start)
}
