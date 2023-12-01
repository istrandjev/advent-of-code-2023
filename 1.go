package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

    spelled := map[string]int{
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    }

    total := 0
    total2 := 0
    re := regexp.MustCompile(`[a-z]+`)
	for scanner.Scan() {
		line := scanner.Text()
        res2 := ""
        for index := range line {
            for key, value := range spelled {
                if index + len(key) <= len(line) && line[index:index + len(key)] == key {
                    res2 += strconv.Itoa(value)
                    break
                }
            }
            res2 += string(line[index])
        }
        number := re.ReplaceAllString(line, "")
        number2 := re.ReplaceAllString(res2, "")
        
        first_part := string(number[0]) + string(number[len(number) - 1])
        i, _ := strconv.Atoi(first_part)
        total += i
        second_part := string(number2[0]) + string(number2[len(number2) - 1])
        j, _ := strconv.Atoi(second_part)
        total2 += j
        
	}
    fmt.Println("Part 1", total)
    fmt.Println("Part 2", total2)
}