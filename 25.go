package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
func AddNone(node string, mapping map[string]int, ne [][]int) [][]int {
    nodeCode, exists := mapping[node]
    if exists {
        return ne
    }
    nodeCode = len(mapping)
    mapping[node] = nodeCode
    ne = append(ne, []int{})
    return ne
}
func AddEdge(from string, to string, mapping map[string]int, ne [][]int) [][]int{
    ne = AddNone(from, mapping, ne)
    ne = AddNone(to, mapping, ne)
    fromCode, toCode := mapping[from], mapping[to]
    ne[fromCode] = append(ne[fromCode], toCode)
    ne[toCode] = append(ne[toCode], fromCode)
    return ne
}
func ParseLine(line string, mapping map[string]int, ne [][]int) [][]int{
    parts := strings.Fields(line)

    result := ne
    for _, to := range parts[1:] {
        result = AddEdge(parts[0][0:len(parts[0]) - 1], to, mapping, result)
    }
    return result
}

type Queue struct {
    items []int
}
func (q *Queue) Enqueue(item int) {
    q.items = append(q.items, item)
}
func (q *Queue) Deque() int {
    result := q.items[0]
    q.items = q.items[1:]
    return result
}
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

func min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}
func max(x int, y int) int {
    if x > y {
        return x
    }
    return y
}
var pre = []int{}
var low = []int{}
var cnt = 0
var ans = [][2]int{}
var exluded = map[[2]int]bool{}
func IsExcluded(a int, b int) bool {
    _, exists := exluded[[2]int{min(a, b), max(a, b)}]
    return exists
}
func Dfs(ne [][]int, dad int, v int) {
    pre[v] = cnt
    cnt++
    low[v] = pre[v]
    for i := range ne[v] {
        if IsExcluded(v, ne[v][i]) {
            continue
        }
        if pre[ne[v][i]] == -1 {
            Dfs(ne, v, ne[v][i]);
            low[v] = min(low[v], low[ne[v][i]])
            if low[ne[v][i]] > pre[v] {
                ans = append(ans, [2]int{min(v, ne[v][i]), max(v, ne[v][i])})
            }
        } else if ne[v][i] != dad {
            low[v] = min(low[v], low[ne[v][i]])
        }
    }
}

func GetBridges(ne [][]int) [][2]int {
    ans = [][2]int{}
    cnt = 0
    if len(pre) < len(ne) {
        pre = make([]int, len(ne))
        low = make([]int, len(ne))
    }
    for i := range ne {
        pre[i] = -1
        low[i] = -1
    }
    
    for i := range ne {
        if pre[i] == -1 {
            Dfs(ne, i, i)
        }
    }
    return ans
}

func RemovedEdges(ne [][]int) [][2]int {
    allEdges := [][2]int{}
    for from, v := range ne {
        for _, to := range v {
            allEdges = append(allEdges, [2]int{min(from, to), max(from, to)})
        }
    }
    for i, edge1 := range allEdges {
        for _, edge2 := range allEdges[i + 1:] {
            exluded = map[[2]int]bool{edge1: true, edge2: true}
            res := GetBridges(ne)
            if len(res) > 0 {
                if len(res) > 1 {
                    fmt.Println("Problem!", len(res))
                }
                return [][2]int{edge1, edge2, res[0]}
            }
        }
    }
    fmt.Println("Not found!")
    return [][2]int{}
}

func GetComponentSizes(ne [][]int) int {
    visited := make([]bool, len(ne))
    result := 1
    for v := range ne {
        if visited[v] {
            continue
        }	
        componentSize := 1
        visited[v] = true
        q := Queue{}
        q.Enqueue(v)
        for !q.IsEmpty() {
            cur := q.Deque()
            for _, next := range ne[cur] {
                if IsExcluded(cur, next) {
                    continue
                }
                if visited[next] {
                    continue
                }
                visited[next] = true
                componentSize++
                q.Enqueue(next)
            }
        }
        result *= componentSize
    }
    return result
}
func SolvePart1(ne [][]int) int {
    t := 0
    for _, v := range ne {
        t += len(v)
    }
    res := RemovedEdges(ne)
    exluded = map[[2]int]bool{}
    for _, e := range res {
        exluded[e] = true
    }
    return GetComponentSizes(ne)
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    mapping := map[string]int{}
    ne := [][]int{}
    for scanner.Scan() {
        line := scanner.Text()
        ne = ParseLine(line, mapping, ne)
    }

    result1 := SolvePart1(ne)
    fmt.Println("Part 1", result1)
}