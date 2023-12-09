package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
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

func interpolateSequence(sequence []int, onLeft bool) []int {
    allZeroes := true
    for _, val := range sequence {
        if val != 0 {
            allZeroes = false
            break
        }
    }
    if allZeroes {
        return append(sequence, 0)
    }

    differences := []int{}
    for i := 0; i < len(sequence) - 1; i++ {
        differences = append(differences, sequence[i + 1] - sequence[i])
    }

    differences = interpolateSequence(differences, onLeft)
    if onLeft {
        newFirst := sequence[0] - differences[0]
        sequence = append([]int{newFirst}, sequence...)
    } else {
        newLast := sequence[len(sequence) - 1] + differences[len(differences) - 1]
        sequence = append(sequence, newLast)
    }
    return sequence
}


func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var sequences [][]int
    result1 := 0
    result2 := 0
    for scanner.Scan() {
        sequence := mapFn(strings.Fields(scanner.Text()), toInt)
        sequences = append(sequences, sequence)
        sequenceRight := interpolateSequence(sequence, false)
        result1 += sequenceRight[len(sequenceRight) - 1]
        sequenceLeft := interpolateSequence(sequence, true)
        result2 += sequenceLeft[0]
    }

    fmt.Println("Part 1", result1)
    fmt.Println("Part 2", result2)
}

