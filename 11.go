package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

func CountSmaller(value int, collection []int) int {
    if collection[len(collection) - 1] < value {
        return len(collection)
    }
    if collection[0] > value {
        return 0
    }

    beg := 0
    end := len(collection)
    for end - beg > 1 {
        mid := (end + beg) / 2
        if collection[mid] < value {
            beg = mid
        } else {
            end = mid
        }
    }
    return end
}

func CountEmpties(from int, to int, empties []int, expansion int) int {
    a, b := from, to
    if a > b {
        a, b = b, a
    }
    return b - a + (CountSmaller(b, empties) - CountSmaller(a, empties)) * (expansion - 1)
}
func Solve1(galaxy1 [2]int, galaxy2 [2]int, emptyRows []int, emptyCols []int, expansion int) int {
    return CountEmpties(galaxy1[0], galaxy2[0], emptyRows, expansion) + CountEmpties(galaxy1[1], galaxy2[1], emptyCols, expansion)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sky := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        sky =  append(sky, line)
    }
    galaxiesList := [][2]int{}
    emptyRowsMap := map[int]bool{}
    emptyColsMap := map[int]bool{}

    for i := range sky {
        emptyRowsMap[i] = true
    }
    for j := range sky[0] {
        emptyColsMap[j] = true
    }
    for i, skyRow := range sky {
        for j := range skyRow {
            if string(skyRow[j]) == "#" {
                galaxiesList = append(galaxiesList, [2]int{i, j})
                delete(emptyRowsMap, i)
                delete(emptyColsMap, j)
            }
        }
    }
    emptyRows := []int{}
    for i := range emptyRowsMap {
        emptyRows = append(emptyRows, i)
    }
    sort.Ints(emptyRows)
    emptyCols := []int{}
    for i := range emptyColsMap {
        emptyCols = append(emptyCols, i)
    }
    sort.Ints(emptyCols)
    result1 := 0
    for i, g1 := range galaxiesList {
        for _, g2 := range galaxiesList[i + 1:] {
            result1 += Solve1(g1, g2, emptyRows, emptyCols, 2)
        }
    }
    fmt.Println("Part 1", result1)
    result2 := 0
    for i, g1 := range galaxiesList {
        for _, g2 := range galaxiesList[i + 1:] {
            result2 += Solve1(g1, g2, emptyRows, emptyCols, 1000000)
        }
    }
    fmt.Println("Part 2", result2)
}