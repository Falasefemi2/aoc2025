package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
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
}
