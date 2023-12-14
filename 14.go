package main

import (
    "bufio"
    "fmt"
    "os"
)
var MOVE4 = [][2]int{ { -1, 0 }, { 0, -1 }, { 1, 0 }, { 0, 1 } }

func GetArray(rawField []string) [][]rune {
    var field [][]rune
    for _, line := range rawField {
        row := []rune{}
        for _, c := range line {
            row = append(row, c)
        }
        
        field = append(field, row)
    }
    return field
}

func GetLast(dirStep int, maxValue int) int {
    if dirStep < 0 {
        return -1
    } else {
        return maxValue
    }
}

func MoveVertical(result [][]rune, dir int) [][]rune {
    for y := 0; y < len(result[0]); y++ {
        lastx := GetLast(MOVE4[dir][0], len(result))
        for x := lastx - MOVE4[dir][0]; x >= 0 && x < len(result); x -= MOVE4[dir][0] {
            if result[x][y] == '.' {
                continue
            }
            if result[x][y] == '#' {
                lastx = x 
                continue
            }
            result[x][y] = '.'
            lastx = lastx - MOVE4[dir][0]
            result[lastx][y] = 'O'
        }
    }
    return result
}

func MoveHorizontal(result [][]rune, dir int) [][]rune {
    for x := 0; x < len(result); x++ {
        lasty := GetLast(MOVE4[dir][1], len(result[0]))
        for y := lasty - MOVE4[dir][1]; y >= 0 && y < len(result[0]); y -= MOVE4[dir][1] {
            if result[x][y] == '.' {
                continue
            }
            if result[x][y] == '#' {
                lasty = y 
                continue
            }
            result[x][y] = '.'
            lasty = lasty - MOVE4[dir][1]
            result[x][lasty] = 'O'
        }
    }
    return result
}

func MoveInDirection(field [][]rune, dir int) [][]rune {
    var result [][]rune
    for i := range field {
        row := []rune {}
        for _, c := range field[i] {
            row = append(row, c)
        }
        result = append(result, row)
    }

    if dir % 2 == 0 {
        MoveVertical(result, dir)
    } else {
        MoveHorizontal(result, dir)		
    }
    
    return result
}

func GetScore(field [][]rune) int {
    result := 0
    for i, row := range field {
        for _, c := range row {
            if c == 'O' {
                result += len(field) - i
            }
        }
    }
    return result
}

func SolvePart1(field [][]rune) int {
    moved := MoveInDirection(field, 0)
    return GetScore(moved)
}
func GetString(field [][]rune) string {
    result := ""
    for i := range field {
        for _, c := range field[i] {
            result += string(c)
        }
    }
    return result
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    rawField := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        rawField = append(rawField, line)
    }
    field := GetArray(rawField)
    result1 := SolvePart1(field)
    fmt.Println("Part 1", result1)

    field = GetArray(rawField)
    cache := map[string]int{}
    answers := []int{}
    result2 := 0
    answers = append(answers, GetScore(field))
    for i := 1; i <= 1000; i++ {
        for dir := 0; dir < 4; dir++ {
            field = MoveInDirection(field, dir)
        }
        answers = append(answers, GetScore(field))
        s := GetString(field)
        value, exists := cache[s]
        if exists {
            cycleLength := i - value
            cycleStart := value
            target := (1000000000 - cycleStart) % cycleLength + cycleStart
            result2 = answers[target]
            break
        } else {
            cache[s] = i
        }
    }
    fmt.Println("Part 2", result2)
}