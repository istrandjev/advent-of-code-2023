package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func toInt(number_str string) int {
    value, _ := strconv.Atoi(number_str)
    return value
}

type Node struct {
    x int
    y int
    dir int
    steps int
}
type HeapItem struct {
    value int
    node Node
}
func EmptyNode() Node {
    return Node{-1, -1, -1, -1}
}
func EmptyHeapItem() HeapItem {
    return HeapItem{-1, EmptyNode()}
}
type Heap struct {
    items []HeapItem
    size int
    capacity int
}
func (h *Heap) _ConstructEmptyHeap() {
    h.items = append(h.items, EmptyHeapItem())
    h.capacity = 0
    h.size = 0
}
func (h *Heap) _EnsureItemCanBeAdded() {
    if len(h.items) == 0 {
        h._ConstructEmptyHeap()
    }
    for h.capacity < h.size + 1 {
        toAdd := len(h.items)
        for i := 0; i < toAdd; i++ {
            h.items = append(h.items, EmptyHeapItem())
            h.capacity++
        }
    }
}
func (h *Heap) Push(value int, n Node) {
    h._EnsureItemCanBeAdded()
    h.size++
    h.items[h.size] = HeapItem{value, n}
    h.SiftUp(h.size)
}
func (h *Heap) SiftUp(idx int) {
    for idx > 1 {
        parent := idx / 2
        if h.items[parent].value > h.items[idx].value {
            h.items[parent], h.items[idx] = h.items[idx], h.items[parent]
        }
        idx /= 2
    }
}
func (h *Heap) IsEmpty() bool {
    return h.size == 0
}
func (h *Heap) Pop() HeapItem {
    result := h.items[1]
    h.items[1] = h.items[h.size]
    h.size--
    h.SinkDown(1)
    return result
}
func (h *Heap) SinkDown(idx int) {
    for idx * 2 <= h.size {
        child := idx * 2
        if child + 1 <= h.size && h.items[child + 1].value < h.items[child].value {
            child++
        }
        if h.items[child].value < h.items[idx].value {
            h.items[child], h.items[idx] = h.items[idx], h.items[child]
        }
        idx = child
    }
}
func ConstructHeap() Heap {
    result := Heap{}
    result._ConstructEmptyHeap()
    return result
}

var field [][]int
var MOVE4 = [][2]int{ { 1, 0 }, { 0, -1 }, { -1, 0 }, { 0, 1 } }

func Dijkstra(field [][]int, maxSteps int, minSteps int) int {
    visited := map[Node]int{}
    dist := map[Node]int{}
    h := ConstructHeap()
    n := len(field)
    m := len(field[0])
    h.Push(0, Node{0, 0, 0, 0})
    for !h.IsEmpty() {
        cur := h.Pop()
        _, isVisited := visited[cur.node]
        if isVisited {
            continue
        }
        d := cur.value
        visited[cur.node] = d
        if cur.node.x + 1 ==  n && cur.node.y + 1 == m {
            return d
        }
        for i := 3; i <= 5; i++ {
            newDir := (cur.node.dir + i) % 4
            newSteps := 1
            if newDir == cur.node.dir {
                newSteps = cur.node.steps + 1
            } 
            tx := cur.node.x + MOVE4[newDir][0]
            ty := cur.node.y + MOVE4[newDir][1]
            if tx < 0 || ty < 0 || tx >= n || ty >= m || newSteps > maxSteps {
                continue
            }
            if cur.node.steps > 0 && cur.node.steps < minSteps && newDir != cur.node.dir {
                continue
            }
            nextNode := Node{tx, ty, newDir, newSteps}
            td, exists := dist[nextNode]
            if !exists || td > d + field[tx][ty] {
                dist[nextNode] = d + field[tx][ty]
                h.Push(d + field[tx][ty], nextNode)
            }
        }
    }
    return -1
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        row := []int{}
        for _, c := range line {
            row = append(row, toInt(string(c)))
        }
        field = append(field, row)
    }
    result1 := Dijkstra(field, 3, 0)	
    fmt.Println("Part 1", result1)
    result2 := Dijkstra(field, 10, 4)	
    fmt.Println("Part 2", result2)
}