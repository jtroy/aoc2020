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

type containedBag struct {
	containedByColor string
	numContained     int
}

func doPart1(lines []string) (int, time.Duration) {
	start := time.Now()
	bags := make(map[string][]containedBag)
	for _, line := range lines {
		fields := strings.Fields(line)
		containedByColor := fields[0] + " " + fields[1]
		if fields[4] == "no" {
			containers := bags["none"]
			bags["none"] = append(containers, containedBag{containedByColor, 0})
			continue
		}
		fields = fields[4:]
		for i := 0; i < len(fields)/4; i++ {
			bagColor := fields[4*i+1] + " " + fields[4*i+2]
			numContained, err := strconv.Atoi(fields[4*i])
			if err != nil {
				panic(err)
			}
			containers := bags[bagColor]
			bags[bagColor] = append(containers, containedBag{containedByColor, numContained})
		}
	}
	visited := visitUniqContainers(bags, "shiny gold")
	return len(visited), time.Since(start)
}

func visitUniqContainers(bags map[string][]containedBag, bag string) []string {
	uniq := make(map[string]struct{})
	all := visit(bags, bag)
	for _, bag := range all {
		uniq[bag] = struct{}{}
	}
	delete(uniq, bag)
	var res []string
	for bag := range uniq {
		res = append(res, bag)
	}
	return res
}

func visit(bags map[string][]containedBag, bag string) []string {
	containers, ok := bags[bag]
	if !ok {
		return []string{bag}
	}
	var visited []string
	for _, c := range containers {
		thisVisited := append(visit(bags, c.containedByColor), bag)
		visited = append(visited, thisVisited...)
	}
	return visited
}

type bag struct {
	color string
	count int
}

func doPart2(lines []string) (int, time.Duration) {
	start := time.Now()
	bags := make(map[string][]bag)
	for _, line := range lines {
		fields := strings.Fields(line)
		containingBagColor := fields[0] + " " + fields[1]
		var containedBags []bag
		if fields[4] == "no" {
			continue
		}
		fields = fields[4:]
		for i := 0; i < len(fields)/4; i++ {
			containedBagColor := fields[4*i+1] + " " + fields[4*i+2]
			containedBagCount, err := strconv.Atoi(fields[4*i])
			if err != nil {
				panic(err)
			}
			containedBags = append(containedBags, bag{containedBagColor, containedBagCount})
		}
		bags[containingBagColor] = containedBags
	}
	return countContainedBags(bags, "shiny gold") - 1, time.Since(start)
}

func countContainedBags(bags map[string][]bag, bagColor string) int {
	total := 1
	for _, b := range bags[bagColor] {
		total += b.count * countContainedBags(bags, b.color)
	}
	return total
}
