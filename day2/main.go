package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1 := solvePart1()
	part2 := solvePart2()
	fmt.Println("Sum of invalid ID:", part1)
	fmt.Println("Sum of invalid ID:", part2)
}

func solvePart1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// Split by commas → each comma-separated entry is a range
		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			// split "start-end"
			parts := strings.Split(strings.TrimSpace(r), "-")
			if len(parts) != 2 {
				continue
			}

			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			// iterate through range
			for id := start; id <= end; id++ {
				s := strconv.Itoa(id)

				if len(s)%2 != 0 {
					continue
				}

				half := len(s) / 2
				left := s[:half]
				right := s[half:]

				if left == right {
					total += id
				}
			}
		}
	}
	return total
}

func isInvalid(id int) bool {
	s := strconv.Itoa(id)
	for N := 2; N <= 12; N++ {
		if len(s)%N != 0 {
			continue
		}
		chunkSize := len(s) / N
		first := s[:chunkSize]
		allSame := true
		for i := 1; i < N; i++ {
			if s[i*chunkSize:(i+1)*chunkSize] != first {
				allSame = false
				break
			}
		}
		if allSame {
			return true
		}
	}

	return false
}

func solvePart2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			parts := strings.Split(strings.TrimSpace(r), "-")
			if len(parts) != 2 {
				continue
			}
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			for id := start; id <= end; id++ {
				if isInvalid(id) {
					total += id
				}
			}
		}
	}
	return total
}
