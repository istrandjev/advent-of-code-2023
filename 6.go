package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func toInt64(number_str string) int64 {
    value, _ := strconv.ParseInt(number_str, 10, 64)
    return value
}

func parseTokens(scanner *bufio.Scanner) ([]int64, int64) {
    var fields []int64
    combined := ""
    if scanner.Scan() {
        line := scanner.Text()
        for _, value_str := range strings.Fields(line)[1:] {
            combined += value_str
            fields = append(fields, toInt64(value_str))
        }
    }
    return fields, toInt64(combined)
}

func getGood(maxTime int64, target int64) int64 {
    result := int64(0)
    for time := int64(1); time <= maxTime; time++ {
        dist := (maxTime - time) * time	
        if dist > target {
            result++
        }
    }
    return result
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    times, time2 := parseTokens(scanner)
    target, target2 := parseTokens(scanner)

    result1 := int64(1)
    for i := 0; i < len(times); i++ {
        current := getGood(times[i], target[i])
        result1 *= current
    }
    result2 := getGood(time2, target2)
    fmt.Println("Part 1", result1)
    fmt.Println("Part 2", result2)
}