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
var maxDistance = 0
var graph = map[[2]int][]Edge{}

func Dfs(v [2]int, shouldForce bool) {
    if v == target {
        if visited[v[0]][v[1]] > maxDistance {
            maxDistance = visited[v[0]][v[1]]
        }
        return
    }
    for _, edge := range graph[v] {
        if shouldForce && edge.isReversed {
            continue
        }
        if visited[edge.to[0]][edge.to[1]] != -1 {
            continue
        }
        visited[edge.to[0]][edge.to[1]] = visited[v[0]][v[1]] + edge.length
        Dfs(edge.to, shouldForce)
        visited[edge.to[0]][edge.to[1]] = -1
    }
}
type Edge struct {
    from [2]int
    to [2]int
    isReversed bool
    length int
}

func IsNode(v [2]int) bool {
    n := len(field)
    m := len(field[0])
    if v == start || v == target {
        return true
    }
    br := 0
    for _, move := range MOVE4 {
        next := [2]int{v[0] + move[0], v[1] + move[1]}
        if next[0] < 0 || next[1] < 0 || next[0] >= n || next[1] >= m {
            continue
        }
        if field[next[0]][next[1]] == '#' {
            continue
        }
        br++
    }
    return br > 2
}
func GetEdge(v [2]int, processed[][]bool) Edge {
    n := len(field)
    m := len(field[0])
    length := 2
    toCheck := [][2]int{}
    toCheck = append(toCheck, v)
    processed[v[0]][v[1]] = true
    ends := [][2]int{}
    isReversed := false
    for len(toCheck) > 0 {
        cur := toCheck[len(toCheck) - 1]
        toCheck = toCheck[:len(toCheck) - 1]

        for l, move := range MOVE4 {
            next := [2]int{cur[0] + move[0], cur[1] + move[1]}
            if next[0] < 0 || next[1] < 0 || next[0] >= n || next[1] >= m {
                continue
            }
            if field[next[0]][next[1]] == '#' || processed[next[0]][next[1]]{
                continue
            }
            if IsNode(next) {
                ends = append(ends, next)
                continue
            }
            if fd, exists := forced[string(field[next[0]][next[1]])]; exists {
                if (fd + 2) % 4 == l {
                    isReversed = true
                }
            }
            length++
            toCheck = append(toCheck, next)
            processed[next[0]][next[1]] = true
        }
    }

    return Edge{ends[0], ends[1], isReversed, length}
}

func RevereEdge(e Edge) Edge {
    return Edge{e.to, e.from, !e.isReversed, e.length}
}
func AddEdgeToNode(v [2]int, e Edge, graph map[[2]int][]Edge) {
    edges, exists := graph[v]
    if !exists {
        edges = []Edge{}
        graph[v] = edges
    }
    graph[v] = append(edges, e)
}

func AddEdge(e Edge, graph map[[2]int][]Edge) {
    AddEdgeToNode(e.from, e, graph)
    AddEdgeToNode(e.to, RevereEdge(e), graph)
}

func ConstructGraph() map[[2]int][]Edge {
    n := len(field)
    m := len(field[0])
    graph := map[[2]int][]Edge{}
    processed := [][]bool{}
    for range field {
        processed = append(processed, make([]bool, len(field[0])))
    }

    for i, row := range field {
        for j := range row {
            if processed[i][j] || field[i][j] == '#' {
                continue
            }
            v := [2]int{i, j}
            if IsNode(v) {
                for _, move := range MOVE4 {
                    next := [2]int{v[0] + move[0], v[1] + move[1]}
                    if next[0] < 0 || next[1] < 0 || next[0] >= n || next[1] >= m {
                        continue
                    }
                    if field[next[0]][next[1]] == '#' || processed[next[0]][next[1]]{
                        continue
                    }
                    e := GetEdge(next, processed)
                    AddEdge(e, graph)
                }
            }
        }
    }
    return graph
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
    maxDistance = 0
    visited[start[0]][start[1]] = 0
    graph = ConstructGraph()
    Dfs(start, true)
    fmt.Println("Part 1", maxDistance)
    maxDistance = 0
    visited[start[0]][start[1]] = 0
    Dfs(start, false)
    fmt.Println("Part 2", maxDistance)
}