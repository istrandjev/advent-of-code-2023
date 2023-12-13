package main

import (
	"bufio"
	"fmt"
	"os"
)

func RowsMatch(i int,  j int, pattern []string) bool {
	for l := 0; l < len(pattern[0]); l++ {
		if pattern[i][l] != pattern[j][l] {
			return false
		}
	}
	return true
}

func ColsMatch(i int, j int, pattern []string) bool {
	for l := 0; l < len(pattern); l++ {
		if pattern[l][i] != pattern[l][j] {
			return false
		}
	}
	return true
}

func HorizontalAfter(row int, pattern []string, ignore int) bool {
	if row == len(pattern) - 1 || row == ignore{
		return false
	 }
	for i := 0; row - i >= 0 && row + i + 1 < len(pattern); i++ {
		if !RowsMatch(row - i, row + i + 1, pattern) {
			return false
		}
	}
	return true
}

func VerticalAfter(col int, pattern []string, ignore int) bool {
	if col == len(pattern[0]) - 1 || col == ignore{
		return false
	}
	for i := 0; col - i >= 0 && col + i + 1 < len(pattern[0]); i ++ {
		if !ColsMatch(col - i, col + i + 1, pattern) {
			return false
		}
	}
	return true
}

func ProcessPattern(pattern []string, ignoreRow int, ignoreCol int) int {
	for row := 0; row < len(pattern); row++ {
		if HorizontalAfter(row, pattern, ignoreRow) {
			return (row + 1) * 100
		}
	}
	for col := 0 ; col < len(pattern[0]); col++ {
		if VerticalAfter(col, pattern, ignoreCol) {
			return col + 1
		}
	}
	return 0
}

func SwapSymbol(s string, idx int) string {
	newSymbol := "#"
	if string(s[idx]) == "#" {
		newSymbol = "."
	}
	return fmt.Sprintf("%s%s%s", s[:idx], newSymbol, s[idx+1:])
}
func ProcessSmudge(pattern []string) int {
	oldValue := ProcessPattern(pattern, -1, -1)
	ignoreRow := -1
	ignoreCol := -1
	if oldValue < 100 {
		ignoreCol = oldValue - 1
	} else {
		ignoreRow = oldValue / 100 - 1
	}
	for i := range pattern {
		for j  := range pattern[i] {
			pattern[i] = SwapSymbol(pattern[i], j)
			temp := ProcessPattern(pattern, ignoreRow, ignoreCol)
			pattern[i] = SwapSymbol(pattern[i], j)
			if temp != 0 {
				return temp
			}
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var patterns [][]string
	current := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			patterns = append(patterns, current)
			current = []string{}
		} else {
			current = append(current, line)
		}
	}

	patterns = append(patterns, current)
	result1 := 0
	for _, pattern := range patterns {
		temp := ProcessPattern(pattern, -1, -1)
		result1 += temp
	}
	fmt.Println("Part 1", result1)

	result2 := 0
	for _, pattern := range patterns {
		temp := ProcessSmudge(pattern)
		result2 += temp
	}
	fmt.Println("Part 2", result2)
}