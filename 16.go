package main

import (
    "bufio"
    "fmt"
    "os"
)

type Node struct {
    dir int
    p [2]int
}
type Queue struct {
    items []Node
}
func (q *Queue) Enqueue(item Node) {
    q.items = append(q.items, item)
}
func (q *Queue) Deque() Node {
    result := q.items[0]
    q.items = q.items[1:]
    return result
}
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}
var MOVE4 = [][2]int{ { -1, 0 }, { 0, -1 }, { 1, 0 }, { 0, 1 } }
var reflections = map[string][4]int{
    "/": {3, 2, 1, 0},
    "\\": {1, 0, 3, 2},
}

func valid(tx int, ty int, field []string) bool {
    return tx >= 0 && ty >= 0 && tx < len(field) && ty < len(field[0])
}

func processNextNode(dir int, cur Node, visited map[Node]int, q *Queue, field []string) {
    tx := cur.p[0] + MOVE4[dir][0]
    ty := cur.p[1] + MOVE4[dir][1]

    if !valid(tx, ty, field) {
        return
    }
    next := Node{dir, [2]int{tx, ty}}
    _, exists := visited[next]
    if !exists {
        visited[next] = visited[cur] + 1
        (*q).Enqueue(next)
    }
}
func CountEnergized(field []string, startNode Node) int {
    q := Queue{}
    q.Enqueue(startNode)
    visited := map[Node]int{}
    visited[startNode] = 0

    for !q.IsEmpty() {
        cur := q.Deque()
        c := string(field[cur.p[0]][cur.p[1]])
        if c == "." || (c == "|" && cur.dir % 2 == 0) || (c == "-" && cur.dir % 2 == 1){
            processNextNode(cur.dir, cur, visited, &q, field)
        } else if c == "|" || c == "-" {
            for dd := 1; dd <= 3; dd += 2 {
                nextDir := (cur.dir + dd) % 4
                processNextNode(nextDir, cur, visited, &q, field)
            }
        } else {
            mappedDir := reflections[c][cur.dir]
            processNextNode(mappedDir, cur, visited, &q, field)
        }
    }
    visitedPositions := map[[2]int]bool{}
    for k := range visited {
        visitedPositions[k.p] = true
    }

    return len(visitedPositions)
}

func CountEdges(i int, j int, field []string) int {
    result := 0
    if i == 0 || i + 1 == len(field) {
        result += 1
    }
    if j == 0 || j + 1 == len(field[i]) {
        result += 1
    }
    return result
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    field := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        field = append(field, line)
    }

    result1 := CountEnergized(field, Node{3, [2]int{0, 0}})
    fmt.Println("Part 1", result1)

    result2 := 0
    for i := 0; i < len(field); i++ {
        for j := 0; j < len(field[i]); j++ {
            edges := CountEdges(i, j, field)
            if edges == 0 {
                continue
            }
            for dir := 0; dir < 4; dir++ {
                tx := i + MOVE4[dir][0]
                ty := j + MOVE4[dir][1]
                if !valid(tx, ty, field) || CountEdges(tx, ty, field) == edges {
                    continue
                }
                current := CountEnergized(field, Node{dir, [2]int{i, j}})
                if current > result2 {
                    result2 = current
                }
            }
        }
    }
    fmt.Println("Part 2", result2)
}