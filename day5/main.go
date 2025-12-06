package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	part1 := part1()
	part2 := part2()
	fmt.Println("fresh", part1)
	fmt.Println("fresh", part2)
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var ranges [][2]int
	var ingredientIDs []int
	scanner := bufio.NewScanner(file)
	parsingRanges := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			parsingRanges = false
			continue
		}
		if parsingRanges {
			parts := strings.Split(line, "-")
			if len(parts) == 2 {
				min, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
				max, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
				ranges = append(ranges, [2]int{min, max})
			}
		} else {
			id, _ := strconv.Atoi(line)
			ingredientIDs = append(ingredientIDs, id)
		}
	}
	freshCount := 0
	for _, id := range ingredientIDs {
		for _, r := range ranges {
			if id >= r[0] && id <= r[1] {
				freshCount++
				break
			}
		}
	}
	return freshCount
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var ranges [][2]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			min, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			max, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			ranges = append(ranges, [2]int{min, max})
		}
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})
	var merged [][2]int
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
		} else {
			last := &merged[len(merged)-1]
			if r[0] <= last[1]+1 {
				if r[1] > last[1] {
					last[1] = r[1]
				}
			} else {
				merged = append(merged, r)
			}
		}
	}
	total := 0
	for _, r := range merged {
		total += r[1] - r[0] + 1
	}

	return total
}
