package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strconv"
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

type Point struct {
    x int
    y int
    z int
}


type Brick struct {
    beg Point
    end Point
    originalIndex int
}

type ByZ []Brick
func (a ByZ) Len() int           { return len(a) }
func (a ByZ) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByZ) Less(i, j int) bool { return a[i].beg.z < a[j].beg.z }

func ParseBrick(line string, brickIndex int) Brick {
    reBrick := regexp.MustCompile(`(\d+),(\d+),(\d+)~(\d+),(\d+),(\d+)`)
    match := reBrick.FindStringSubmatch(line)
    coords := mapFn(match[1:], toInt)

    beg := Point{
        min(coords[0], coords[3]),
        min(coords[1], coords[4]),
        min(coords[2], coords[5]),
    }
    end := Point{
        max(coords[0], coords[3]),
        max(coords[1], coords[4]),
        max(coords[2], coords[5]),
    }
    return Brick{beg, end, brickIndex}
}
type Cell struct {
    height int
    brickNumber int
}

type Set struct {
    values map[int]bool
}
func ConstructSet() Set {
    return Set{map[int]bool{}}
}
func (s *Set)Add(v int) {
    s.values[v] = true
}
func (s *Set)Contains(v int) bool {
    _, exists := s.values[v]
    return exists
}
func (s *Set)Delete(v int) {
    delete(s.values, v)
}
func (s *Set)Len() int {
    return len(s.values)
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

func GetField(bricks []Brick) (Set, [][]int){
    minx, maxx := 1000000, 0
    miny, maxy := 1000000, 0

    for _, brick := range bricks {
        minx = min(brick.beg.x, minx)
        miny = min(brick.beg.y, miny)
        maxx = max(brick.beg.x, maxx)
        maxy = max(brick.beg.y, maxy)
    }

    field := [][]Cell{}
    for x := 0; x <= maxx; x++ {
        row := []Cell{}
        for y := 0; y <= maxy; y++ {
            row = append(row, Cell{0, -1})
        }
        field = append(field, row)
    }
    brickToSupports := map[int]int{}
    brickToSupported := map[int]Set{}
    for bi, brick := range bricks {
        supports := ConstructSet()
        maxHeight := 0
        for x := brick.beg.x; x <= brick.end.x; x++ {
            for y := brick.beg.y; y <= brick.end.y; y++ {
                height := field[x][y].height
                if maxHeight < height {
                    maxHeight = height
                    supports = ConstructSet()
                    if field[x][y].brickNumber != - 1 {
                        supports.Add(field[x][y].brickNumber)
                    }
                } else if maxHeight == height {
                    if field[x][y].brickNumber != - 1 {
                        supports.Add(field[x][y].brickNumber)
                    }
                } 
            }
        }
        brickToSupports[bi] = supports.Len()
        newHeight := maxHeight + brick.end.z - brick.beg.z + 1
        for x := brick.beg.x; x <= brick.end.x; x++ {
            for y := brick.beg.y; y <= brick.end.y; y++ {
                if field[x][y].height == maxHeight {
                    if field[x][y].brickNumber != - 1 {
                        supported, exists := brickToSupported[field[x][y].brickNumber]
                        if !exists {
                            supported = ConstructSet()
                            brickToSupported[field[x][y].brickNumber] = supported
                        } 
                        supported.Add(bi)
                    }
                    
                } 
                field[x][y] = Cell{newHeight, bi}
            }
        }
    }

    ne := [][]int{}
    for bi := range bricks {
        ne = append(ne, []int{})
        supported, exists := brickToSupported[bi]
        if !exists {
            continue
        }
        for s := range supported.values {
            ne[bi] = append(ne[bi], s)
        }
    }

    supportedByOne := ConstructSet()
    for bi, entry := range brickToSupports {
        if entry == 1 {
            supportedByOne.Add(bi)
        }
    }
    return supportedByOne, ne
}

var bricks = []Brick{}
func Dfs(bi int, ne [][]int) int {
    counts := make([]int, len(bricks))
    for _, x := range ne {
        for _, nxt := range x {
            counts[nxt]++
        }
    }
    toProcess := Queue{}
    toProcess.Enqueue(bi)
    result := 0
    for !toProcess.IsEmpty() {
        cur := toProcess.Deque()
        for _, nxt := range ne[cur] {
            counts[nxt]--
            if counts[nxt] == 0 {
                result++
                toProcess.Enqueue(nxt)
            }
        }
    }
    return result
}
func SolvePart2(supportedByOne Set, ne [][]int) int {
    starting := ConstructSet()
    for bi := range bricks {
        for _, s := range ne[bi] {
            if supportedByOne.Contains(s) {
                starting.Add(bi)
            }
        }
    }
    result := 0
    for bi := range starting.values {
        result += Dfs(bi, ne)
    }
    return result
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    brickIndex := 1
    for scanner.Scan() {
        line := scanner.Text()
        bricks = append(bricks, ParseBrick(line, brickIndex))
        brickIndex++
    }
    sort.Sort(ByZ(bricks))

    supportedByOne, ne := GetField(bricks)
    
    result1 := 0
    for bi := range bricks {
        required := false
        for _, s := range ne[bi] {
            if supportedByOne.Contains(s) {
                required = true
            }
        }
        if !required {
            result1++
        }
    }
    fmt.Println("Part 1", result1)
    result2 := SolvePart2(supportedByOne, ne)
    fmt.Println("Part 2", result2)
}