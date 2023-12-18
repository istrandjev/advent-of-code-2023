package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)
var MOVE4 = [][2]int{ { 1, 0 }, { 0, -1 }, { -1, 0 }, { 0, 1 } }
var directions = map[string]int {
    "D": 0,
    "L": 1,
    "U": 2,
    "R": 3,
}
type Instruction struct {
    steps int
    direction string
}

func toInt(number_str string) int {
    value, _ := strconv.Atoi(number_str)
    return value
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

type Queue struct {
    items [][2]int
}
func (q *Queue) Enqueue(item [2]int) {
    q.items = append(q.items, item)
}
func (q *Queue) Deque() [2]int {
    result := q.items[0]
    q.items = q.items[1:]
    return result
}
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}
func GetCoordinates(instructions []Instruction) ([]int, []int) {
    cx, cy := 0, 0
    xs := []int{0}
    ys := []int{0}
    for _, inst := range instructions {
        move := MOVE4[directions[inst.direction]]
        dx, dy := inst.steps * move[0], inst.steps * move[1]

        cx += dx
        cy += dy
        xs = append(xs, cx)
        ys = append(ys, cy)
    }

    return xs, ys
}

func ProcessCoordinates(coordinates []int) ([]int, map[int]int) {
    unique := map[int]bool{}
    for _, x := range coordinates {
        unique[x] = true
        unique[x - 1] = true
        unique[x + 1] = true
    }
    uniqueCoordinates := []int{}
    for k := range unique {
        uniqueCoordinates = append(uniqueCoordinates, k)
    }
    sort.Ints(uniqueCoordinates)
    mapping := map[int]int{}
    for i, c := range uniqueCoordinates {
        mapping[c] = i
    }
    return uniqueCoordinates, mapping
}

func GetLen(x int, xs []int) int {
    if x + 1 == len(xs) {
        return 0
    }
    return xs[x + 1] - xs[x]
}
func GetRectArea(x int, y int, xs []int, ys []int) int {
    xDelta := GetLen(x, xs)
    yDelta := GetLen(y, ys)
    return xDelta * yDelta
}

func SolvePart1(instructions []Instruction) int {
    xsGiven, ysGiven := GetCoordinates(instructions)

    xs, xMapping := ProcessCoordinates(xsGiven)
    ys, yMapping := ProcessCoordinates(ysGiven)

    var field [][]bool
    for x := 0; x < len(xs); x++ {
        row := make([]bool, len(ys))
        field = append(field, row)
    }
    cx, cy := xMapping[0], yMapping[0]
    for _, inst := range instructions {
        move := MOVE4[directions[inst.direction]]
        if move[0] == 0 {
            endy := yMapping[ys[cy] + move[1] * inst.steps]
            for i := min(cy, endy); i <= max(cy, endy); i++ {
                field[cx][i] = true
            }
            cy = endy
        } else {
            endx := xMapping[xs[cx] + move[0] * inst.steps]
            for i := min(cx, endx); i <= max(cx, endx); i++ {
                field[i][cy] = true
            }
            cx = endx
        }
    }

    borders := [][2]int{}
    for i := 0; i < len(field); i++ {
        if !field[i][0] {
            borders = append(borders, [2]int{i, 0})
        }
        if !field[i][len(field[0]) - 1] {
            borders = append(borders, [2]int{i, len(field[0]) - 1})
        }
    }
    for j := 0; j < len(field[0]); j++ {
        if j == 0 || j + 1 == len(field[0]) {
            continue
        }
        if !field[0][j] {
            borders = append(borders, [2]int{0, j})
        }
        if !field[len(field) - 1][j] {
            borders = append(borders, [2]int{len(field) - 1, j})
        }
    }
    uncovered := 0
    visited := map[[2]int]bool{}
    q := Queue{}
    for _, b := range borders {
        q.Enqueue(b)
        visited[b] = true
        uncovered += GetRectArea(b[0], b[1], xs, ys)
    }
    for !q.IsEmpty() {
        cur := q.Deque()
        
        for _, move := range MOVE4 {
            tx := cur[0] + move[0]
            ty := cur[1] + move[1]
            if tx < 0 || ty < 0 || tx >= len(field) || ty >= len(field[0]) {
                continue
            }
            if field[tx][ty] {
                continue
            }
            next := [2]int{tx, ty}
            _, exists := visited[next]
            if !exists {
                q.Enqueue(next)
                visited[next] = true
                uncovered += GetRectArea(tx, ty, xs, ys)
            }
        }
    }

    xDelta := xs[len(xs) - 1] - xs[0]
    yDelta := ys[len(ys) - 1] - ys[0]
    return xDelta * yDelta - uncovered
}
func ParseInstruction(color string) Instruction {
    dir_map := "RDLU"
    result, _ := strconv.ParseInt(color[1:len(color)-1], 16, 64)
    dir := string(dir_map[toInt(string(color[len(color)-1]))])
    return Instruction{int(result), dir}
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    instructions1 := []Instruction{}
    instructions2 := []Instruction{}
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        instructions1 = append(instructions1, Instruction{toInt(fields[1]), fields[0]})
        instructions2 = append(instructions2, ParseInstruction(fields[2][1:len(fields[2]) -1]))
    }
    

    result1 := SolvePart1(instructions1)
    fmt.Println("Part 1", result1)
    result2 := SolvePart1(instructions2)
    fmt.Println("Part 2", result2)
}