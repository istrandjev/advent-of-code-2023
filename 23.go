package main

import (
    "bufio"
    "fmt"
    "os"
)

var field = []string{}
var visited = [][]int{}
var MOVE4 = [][2]int{ { 1, 0 }, { 0, -1 }, { -1, 0 }, { 0, 1 } }
var forced = map[string]int{"^": 2, ">": 3, "v": 0, "<": 1}
var start = [2]int{}
var target = [2]int{}
var maxDistance1 = 0
func DfsPart1(v [2]int, shouldForce bool) {
    n := len(field)
    m := len(field[0])
    forcedMove := -1
    if v == target {
        if visited[v[0]][v[1]] > maxDistance1 {
            maxDistance1 = visited[v[0]][v[1]]
        }
        return
    }
    if shouldForce {
        if val, exists := forced[string(field[v[0]][v[1]])]; exists {
            forcedMove = val
        }
    }
    for l, move := range MOVE4 {
        if forcedMove != -1 && l != forcedMove {
            continue
        }
        next := [2]int{v[0] + move[0], v[1] + move[1]}
        if next[0] < 0 || next[1] < 0 || next[0] >= n || next[1] >= m {
            continue
        }
        if field[next[0]][next[1]] == '#' || visited[next[0]][next[1]] != -1 {
            continue
        }
        visited[next[0]][next[1]] = visited[v[0]][v[1]] + 1
        DfsPart1(next, shouldForce)
        visited[next[0]][next[1]] = -1
    }
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        field = append(field, line)
    }
    for _, field_row := range field {
        row := []int{}
        for range field_row {
            row = append(row, -1)
        }
        visited = append(visited, row)
    }
    
    for i := range field[0] {
        if field[0][i] != '#' {
            start = [2]int{0, i}
        }
        if field[len(field) - 1][i] != '#' {
            target = [2]int{len(field) - 1, i}
        }
    }
    maxDistance1 = 0
    visited[start[0]][start[1]] = 0
    DfsPart1(start, true)
    fmt.Println("Part 1", maxDistance1)
    maxDistance1 = 0
    visited[start[0]][start[1]] = 0
    DfsPart1(start, false)
    fmt.Println("Part 2", maxDistance1)
}