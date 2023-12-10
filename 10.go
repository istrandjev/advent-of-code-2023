package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode/utf8"
)

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

var move4 = [][2]int{ { -1, 0 }, { 0, 1 }, { 1, 0 }, { 0, -1 } }
var runeMoves = map[rune][]int{
    '|': []int{0, 2},
    '-': []int{1, 3},
    'L': []int{0, 1},
    'J': []int{0, 3},
    '7': []int{3, 2},
    'F': []int{2, 1},
}

func GetRuneAtIndex(s string, i int) rune {
    r, _ := utf8.DecodeRuneInString(s[i:])
    return r
}

func GetNeighbours(p [2]int, fieldMap []string) [][2]int{
    result := [][2]int{}
    currentRune := GetRuneAtIndex(fieldMap[p[0]], p[1])

    r, exists := runeMoves[currentRune]
    if !exists {
        return result
    }
    
    for _, v := range r {
        candidate := [2]int{p[0] + move4[v][0], p[1] + move4[v][1]}
        if candidate[0] >= 0 && candidate[0] < len(fieldMap) && candidate[1] >= 0 && candidate[1] < len(fieldMap[0]) {
            result = append(result, candidate)
        }
    }

    return result
}

func GetValidNeighbours(p [2]int, fieldMap []string) [][2]int {
    result := GetNeighbours(p, fieldMap)
    validatedResult := [][2]int{}
    for _, r := range result {
        temp := GetNeighbours(r, fieldMap)
        good := false
        for _, t := range temp {
            if t == p {
                good = true
                break
            }
        }
        if good {
            validatedResult = append(validatedResult, r)
        }
    }
    return validatedResult
}
func FindConnected(start [2]int, fieldMap []string) map[[2]int]int {
    result := map[[2]int]int{}

    result[start] = 0
    q := Queue{}
    q.Enqueue(start)
    for !q.IsEmpty() {
        cur := q.Deque()

        neighbours := GetValidNeighbours(cur, fieldMap)

        for _, ne := range neighbours {
            _, visited := result[ne]
            if !visited {
        		result[ne] = result[cur] + 1
        		q.Enqueue(ne)
            }
        }
    }
    return result
}

func getType(r rune, horizontal bool) int {
    if horizontal {
        if r == 'F' || r == '7' {
            return 0
        } else {
            return 1
        }
    } else {
        if r == 'L' || r == 'F' {
            return 0
        } else {
            return 1
        }
    }
}


func GetIntersectionsInDir(pos [2]int, groups [][]int, fieldMap []string, targetGroup int, dir int) int {
    result := 0
    prev := '*'
    for delta := 1; true; delta++ {
        temp := [2]int{
            pos[0] + move4[dir][0] * delta,
            pos[1] + move4[dir][1] * delta,
        }
        if temp[0] < 0 || temp[0] >= len(groups) || temp[1] < 0 || temp[1] >= len(groups[0]) {
            break
        }
        if groups[temp[0]][temp[1]] != targetGroup {
            continue
        }
        r := GetRuneAtIndex(fieldMap[temp[0]], temp[1])
        if dir % 2 == 0 && r == '|' {
            continue
        }
        if dir % 2 == 1 && r == '-' {
            continue
        }
        if r == '|' || r == '-' {
            result++
            continue
        }
        if prev == '*' {
            prev = r
            result++
            continue
        }

        prevType := getType(prev, dir % 2 == 1)
        curType := getType(r, dir % 2 == 1)
        prev = '*'
        if prevType == curType {
            result++
        }
    }    
    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fieldMap := []string{}
    for scanner.Scan() {
        fieldMap = append(fieldMap, scanner.Text())
    }

    var groups [][]int
    for range fieldMap {
        rowSlice := make([]int, len(fieldMap[0]))
        for i := range rowSlice {
            rowSlice[i] = -1
        }
        groups = append(groups, rowSlice)
    }

    start := [2]int{-1, -1}
    for i, row := range fieldMap {
        for j := range row {
            if row[j] == 'S' {
                start = [2]int{i, j}
                break
            }
        }
    }
    groupIndex := 0
    for i := range fieldMap {
        for j := range fieldMap[i] {
            currentRune := GetRuneAtIndex(fieldMap[i], j)
            if groups[i][j] != -1 || currentRune == '.' || currentRune == 'S' {
                continue
            }

            groupCells := FindConnected([2]int{i, j}, fieldMap)
            for k := range groupCells {
                groups[k[0]][k[1]] = groupIndex
            }
            groupIndex++;
        }
    }

    
    for k := range runeMoves {
        fieldMap[start[0]] = fmt.Sprintf("%s%c%s", fieldMap[start[0]][:start[1]], k, fieldMap[start[0]][start[1] + 1:])
        ne := GetValidNeighbours(start, fieldMap)
        if len(ne) == 2 && groups[ne[0][0]][ne[0][1]] == groups[ne[1][0]][ne[1][1]] {
            groups[start[0]][start[1]] = groups[ne[0][0]][ne[0][1]]
            break
        }
    }
    cycle := FindConnected(start, fieldMap)
    result1 := 0
    for _, v := range cycle {
        if v > result1 {
            result1 = v
        }
    }
    fmt.Println("Part 1", result1)
    targetGroup := groups[start[0]][start[1]]
    result2 := 0
    for i, row := range fieldMap {
        for j := range row {
            if groups[i][j] == targetGroup {
                continue
            }
            isIn := false
            for dir := 0; dir < 4; dir++ {
                intersections := GetIntersectionsInDir([2]int{i, j}, groups, fieldMap, targetGroup, dir) 
                if intersections % 2 == 1 {
                    isIn = true
                    break
                }
            }
            if isIn {
                result2++
            }
        }
    }
    fmt.Println("Part 2", result2)
}