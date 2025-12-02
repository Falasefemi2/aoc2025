package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	part1 := solvePart1()
	part2 := solvePart2()
	fmt.Println("Part 1 Password:", part1)
	fmt.Println("Part 2 Password:", part2)
}

func solvePart1() int {
	s := 50
	count := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		switch dir {
		case 'L':
			s = (s - dist + 100) % 100
		case 'R':
			s = (s + dist) % 100
		}
		if s == 0 {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password:", count)
	return count
}

func solvePart2() int {
	pos := 50
	count := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		step := 1
		if dir == 'L' {
			step = -1
		}
		for i := 0; i < dist; i++ {
			pos += step
			if pos < 0 {
				pos = 99
			} else if pos > 99 {
				pos = 0
			}
			if pos == 0 {
				count++
			}
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	return count
}
