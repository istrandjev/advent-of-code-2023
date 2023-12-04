package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
	"strings"
)

func getNumbers(input string) []int {
	var result []int
	for _, number_str := range strings.Split(strings.Trim(input, " "), " ") {
		number_str = strings.Trim(number_str, " ")
		if len(number_str) == 0 {
			continue
		}
		value, _ := strconv.Atoi(number_str)
		result = append(result, value)
	}
	return result
}

func getIntersection(a []int, b []int) int {
	helper := map[int]bool{}
	for _, v := range a {
		helper[v] = true
	}
	res := 0
	for _, x := range b {
		if _, ok := helper[x]; ok {
			res += 1
		}
	}
	return res
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	result1 := 0
	
	var goods []int
	for scanner.Scan() {
		line := scanner.Text()
		card_values := strings.Split(line, ":")[1]
		parts := strings.Split(card_values, "|")
		winning := getNumbers(parts[0])
		mine := getNumbers(parts[1])
		good := getIntersection(winning, mine)
		goods = append(goods, good)
		if good > 0 {
			result1 += 1 << (good - 1)
		}
	}
	var copies []int
	for range goods {
		copies = append(copies, 1)
	}
	
	for i, good := range goods {
		for j := i + 1; j <= i + good && j < len(copies); j++ {
			copies[j] += copies[i]
		}
	}
	result2 := 0
	for _, c := range copies {
		result2 += c
	}
	fmt.Println("Part 1", result1)
	fmt.Println("Part 2", result2)
}