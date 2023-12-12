package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func toInt(number_str string) int {
    value, _ := strconv.Atoi(number_str)
    return value
}

func mapFn[V any, T any](collection []V, fn func(V)T) []T {
    result := []T{}
    for _, val := range collection {
        result = append(result, fn(val))
    }
    return result
}

var mem [][]int
var s string
var groups []int
func solve(idx int, gi int) int {
    if mem[idx][gi] != -1 {
        return mem[idx][gi]
    }
    if idx == len(s) {
        res := 0
        if gi == len(groups) {
            res = 1
        }
        mem[idx][gi] = res
        return res
    }
    if gi == len(groups) {
        
        good := true
        for _, c := range s[idx:] {
            if c == '#' {
                good = false
            }
        }
        mem[idx][gi] = 0
        if good {
            mem[idx][gi] = 1
        }
        return mem[idx][gi]
    }
    if s[idx] == '.' {
        mem[idx][gi] = solve(idx + 1, gi)
        return mem[idx][gi]
    }
    if s[idx] == '#' {
        if idx + groups[gi] > len(s) {
            mem[idx][gi] = 0
            return 0
        }
        good := true
        for j := idx +  1; j < idx + groups[gi]; j++ {
            if s[j] == '.' {
                good = false
                break
            }
        }
        if idx + groups[gi] < len(s) && s[idx + groups[gi]] == '#' {
            good = false
        }
        if !good {
            mem[idx][gi] = 0
            return 0
        }
        result := 1
        if idx + groups[gi] + 1 < len(s) {
            result *= solve(idx + groups[gi] + 1, gi + 1)
        } else if gi + 1 < len(groups) {
            result = 0
        }
        mem[idx][gi] = result
        return result
    }
    result := solve(idx + 1, gi)
    s = fmt.Sprintf("%s%c%s", s[:idx], '#', s[idx+1:])
    temp := solve(idx, gi)
    s = fmt.Sprintf("%s%c%s", s[:idx], '?', s[idx+1:])
    result += temp
    mem[idx][gi] = result
    return result
}

func solveLine(lineString string, groupsGiven []int) int {
    s = lineString
    groups = groupsGiven
    mem = [][]int{}
    for i := 0; i <= len(s); i++ {
        rowMem := []int{}
        for j := 0; j <= len(groups); j++ {
            rowMem = append(rowMem, -1)
        }
        mem = append(mem, rowMem)
    }
    return solve(0, 0)
}


func parseLine(line string) (string, []int) {
    parts := strings.Fields(line)
    lineString := parts[0]
    groups := mapFn(strings.Split(parts[1], ","), toInt)
    return lineString, groups
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    result1 := 0
    result2 := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineString, groupsGiven := parseLine(line)
        current := solveLine(lineString, groupsGiven)
        result1 += current
        groups2 := []int{}
        lineString2 := lineString
        for i := 0; i < 5; i++ {
            groups2 = append(groups2, groupsGiven...)
        }
        for i := 0 ; i < 4; i ++ {
            lineString2 = lineString2 + "?" + lineString
        }
        result2 += solveLine(lineString2, groups2)
    }

    fmt.Println("Part 1", result1)
    fmt.Println("Part 2", result2)
}