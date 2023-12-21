package main

import (
    "bufio"
    "fmt"
    "os"
)
var MOVE4 = [][2]int{ { 1, 0 }, { 0, -1 }, { -1, 0 }, { 0, 1 } }
var field []string
type State struct {
    steps int
    start [2]int
}
var mem = map[State]int{}

func SolvePart1(s State) int {
    if value, exists := mem[s]; exists {
        return value
    }
    steps := s.steps
    start := s.start
    n := len(field)
    m := len(field[0])
    ne := [][]int{}
    for i := 0; i < n * m; i++ {
        ne = append(ne, []int{})
    }
    current := map[int]bool{}
    for i, row := range field {
        for j, c := range row {
            if i == start[0] && j == start[1] {
                current[i * m + j] = true
            }
            if c == '#' {
                continue
            }

            for _, delta := range MOVE4 {
                tx := i + delta[0]
                ty := j + delta[1]
                if tx < 0 || ty < 0 || tx >= n || ty >= m {
                    continue
                }
                if field[tx][ty] == '#' {
                    continue
                }
                ne[i * m + j] = append(ne[i * m + j], tx * m + ty)
            }
        }
    }
    
    for si := 0; si < steps; si++ {
        next := map[int]bool{}
        for k := range current {
            for _, v := range ne[k] {
                next[v] = true
            }
        }
        current = next
    }
    mem[s] = len(current)
    return len(current)
}

var borders = [][2]int{}
func PopulateBorders() {
    n := len(field)
    m := len(field[0])
    interestingx := []int{0, n / 2, n - 1}
    interestingy := []int{0, m / 2, m - 1}

    for i, ix := range interestingx {
        for j, iy := range interestingy {
            if i == 1 && j ==  1 { 
                continue
            }
            borders = append(borders, [2]int{ix, iy})
        }
    }
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func ManhattanDistance(a [2]int, b [2]int) int {
    result := 0
    for i := range a {
        result += abs(a[i] - b[i])
    }
    return result
}
func Process(x int, y int, steps int, start [2]int) int {
    n := len(field)
    m := len(field[0])    

    offset := [2]int{x * n, y * m}
    best := -1
    bestPoint := [2]int{}
    for _, border := range borders {
        point := [2]int{}
        for i, os := range offset {
            point[i] = os + border[i]
        }
        cd := ManhattanDistance(point, start)
        if best == -1 || cd < best {
            best = cd
            bestPoint = border
        }
    }
    if steps - best < 0 {
        return 0
    }
    toDo := steps - best 
    if toDo > 300 {
        toDo = 300 + toDo % 2
    }
    return SolvePart1(State{toDo, bestPoint})
}

func SolvePart2(steps int, start [2] int) int {
    n := len(field)
    m := len(field[0])
    PopulateBorders()

    deltax := steps / n + 2
    
    result := 0
    odds := 0
    evens := 0

    limit := 2

    for x := -deltax; x <= deltax; x++ {
        rem := steps - abs(x * n)
        deltay := rem / m + 2

        if deltay <= limit {
            for y := -deltay; y <= deltay; y++ {
                result += Process(x, y, steps, start)
            }
            continue
        }
        
        for y := -deltay; y < -deltay + limit; y++ {
            result += Process(x, y, steps, start)
        }
        for y := deltay - (limit - 1); y <= deltay; y++ {
            result += Process(x, y, steps, start)
        }

        
        remaining := (deltay - limit) * 2 + 1
        even := remaining / 2
        odd := remaining - even
        if (x + deltay - limit) % 2 != 0 {
            even, odd = odd, even
        }
        evens += even
        odds += odd
    }
    evenAnswer := SolvePart1(State{300, start})
    oddAnswer := SolvePart1(State{301, start})
    return result + odds * oddAnswer + evens * evenAnswer
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    field = []string{}
    for scanner.Scan() {
        line := scanner.Text()
        field = append(field, line)
    }

    start := [2]int{}
    for i, row := range field {
        for j, c := range row {
            if c == 'S' {
                start = [2]int{i, j}
            }
        }
    }
    
    fmt.Println("Part 1", SolvePart1(State{64, start}))
    fmt.Println("Part 2", SolvePart2(26501365, start))
}