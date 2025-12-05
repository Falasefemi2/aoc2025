package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1 := part1()
	part2 := part2()
	fmt.Println("Sum of battries:", part1)
	fmt.Println("Sum of battries:", part2)
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		maxPair := 0
		for i := 0; i < len(line)-1; i++ {
			first := int(line[i] - '0')
			for j := i + 1; j < len(line); j++ {
				second := int(line[j] - '0')
				value := first*10 + second
				if value > maxPair {
					maxPair = value
				}
			}
		}
		total += maxPair
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return total
}

func part2() int64 {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int64 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		const selectCount = 12
		result := make([]byte, 0, selectCount)
		remaining := len(line)
		needed := selectCount
		pos := 0
		for needed > 0 {
			canSkip := remaining - needed
			maxIdx := pos
			maxDigit := line[pos]
			for i := pos; i <= pos+canSkip; i++ {
				if line[i] > maxDigit {
					maxDigit = line[i]
					maxIdx = i
				}
			}
			result = append(result, maxDigit)
			pos = maxIdx + 1
			remaining = len(line) - pos
			needed--
		}
		joltage, err := strconv.ParseInt(string(result), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		total += joltage
	}
	return total
}
