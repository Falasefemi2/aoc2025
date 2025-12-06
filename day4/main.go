package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	part1 := part1()
	part2 := part2()
	fmt.Println("count", part1)
	fmt.Println("count", part2)
}

func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			grid = append(grid, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	accessibleCount := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' {
				neighborCount := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						nr := r + dr
						nc := c + dc
						if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
							if grid[nr][nc] == '@' {
								neighborCount++
							}
						}
					}
				}
				if neighborCount < 4 {
					accessibleCount++
				}
			}
		}
	}
	return accessibleCount
}

func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			grid = append(grid, []byte(line))
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	totalRemoved := 0
	for {
		accessiblePositions := make([]struct{ r, c int }, 0)
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				if grid[r][c] == '@' {
					neighborCount := 0
					for dr := -1; dr <= 1; dr++ {
						for dc := -1; dc <= 1; dc++ {
							if dr == 0 && dc == 0 {
								continue
							}
							nr := r + dr
							nc := c + dc
							if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
								if grid[nr][nc] == '@' {
									neighborCount++
								}
							}
						}
					}
					if neighborCount < 4 {
						accessiblePositions = append(accessiblePositions, struct{ r, c int }{r, c})
					}
				}
			}
		}
		if len(accessiblePositions) == 0 {
			break
		}
		for _, pos := range accessiblePositions {
			grid[pos.r][pos.c] = '.'
		}

		totalRemoved += len(accessiblePositions)
	}
	return totalRemoved
}
