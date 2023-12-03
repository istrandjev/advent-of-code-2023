package main

import (
    "bufio"
    "fmt"
    "os"
)

func parseField(field []string) [][]int {
	var result [][]int
	for _, line := range field {
		line_codes := []int{}
		for j, c := range line {
			if c >= '0' && c <= '9' {
				value := 0
				if j > 0 && line_codes[j - 1] >= 0 {
					value = line_codes[j - 1] * 10 
				}
				value += (int(c) - int('0'))
				line_codes = append(line_codes, value)
			} else if c == '.' {
				line_codes = append(line_codes, -3)
			} else if c == '*' {
				line_codes = append(line_codes, -4)
			} else {
				line_codes = append(line_codes, -2)
			}
		}
		result = append(result, line_codes)
	}
	for i, row := range result {
		for j := len(row) - 1; j >= 0; j-- {
			if result[i][j] < 0 || j == len(row) - 1 || row[j + 1] < 0 {
				continue
			} else {
				result[i][j] = result[i][j + 1]
			}

		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var field []string
	for scanner.Scan() {
		line := scanner.Text()
		field = append(field, line)
	}
	
	parsed_field := parseField(field)

	result1 := 0
	result2 := 0
	
    var numbers [][]int
    for i := 0; i < len(parsed_field); i++ {
        newRow := make([]int, len(parsed_field[0]))
        numbers = append(numbers, newRow)
    }

	moves := [8][2]int{{-1, 0}, {-1, -1}, {0, -1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {-1, 1}} 
	for i, row := range parsed_field {
		for j := range row {
			if parsed_field[i][j] != -4 {
				continue
			}
			count := 0
			mul := 1
			for _, delta := range moves {
				tx := i + delta[0]
				ty := j + delta[1]
				if tx < 0 || tx >= len(parsed_field) || ty < 0 || ty >= len(row) {
					continue
				}
				if parsed_field[tx][ty] >= 0 && numbers[tx][ty] == 0 {
					count += 1

					for i := 1; ty - i >= 0 && parsed_field[tx][ty - i] == parsed_field[tx][ty]; i++ {
						numbers[tx][ty - i] = 1
					}

					for i := 1; ty + i < len(row) && parsed_field[tx][ty + i] == parsed_field[tx][ty]; i++ {
						numbers[tx][ty + i] = 1
					}
					mul *= parsed_field[tx][ty]
				}
			}

			for _, delta := range moves {
				tx := i + delta[0]
				ty := j + delta[1]
				if tx < 0 || tx >= len(parsed_field) || ty < 0 || ty >= len(row) {
					continue
				}
				if parsed_field[tx][ty] >= 0 && numbers[tx][ty] == 0 {
					for i := 1; ty - i >= 0 && parsed_field[tx][ty - i] == parsed_field[tx][ty]; i++ {
						numbers[tx][ty - i] = 0
					}
					for i := 1; ty + i < len(row) && parsed_field[tx][ty + i] == parsed_field[tx][ty]; i++ {
						numbers[tx][ty + i] = 0
					}
				}
			}

			if count == 2 {
				result2 += mul
			}
		}
	}


	
	for i, row := range parsed_field {
		for j := range row {
			if parsed_field[i][j] >= 0 || parsed_field[i][j] == -3 {
				continue
			}
			for _, delta := range moves {
				tx := i + delta[0]
				ty := j + delta[1]
				if tx < 0 || tx >= len(parsed_field) || ty < 0 || ty >= len(row) {
					continue
				}
				if parsed_field[tx][ty] > 0 {
					result1 += parsed_field[tx][ty]

					for i := 1; ty - i >= 0 && parsed_field[tx][ty - i] == parsed_field[tx][ty]; i++ {
						parsed_field[tx][ty - i] = 0
					}

					for i := 1; ty + i < len(row) && parsed_field[tx][ty + i] == parsed_field[tx][ty]; i++ {
						parsed_field[tx][ty + i] = 0
					}
					parsed_field[tx][ty] = 0
				}
			}
		}
	}
	
	fmt.Println("Part 1", result1)
	fmt.Println("Part 2", result2)
}