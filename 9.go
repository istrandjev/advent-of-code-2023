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

func parseSeq(line string) []int {
	tokens := strings.Fields(line)
	result := []int{}
	for _, token := range tokens {
		result = append(result, toInt(token))
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
		sequence := parseSeq(scanner.Text())
		sequences = append(sequences, sequence)
		sequenceRight := interpolateSequence(sequence, false)
		result1 += sequenceRight[len(sequenceRight) - 1]
		sequenceLeft := interpolateSequence(sequence, true)
		result2 += sequenceLeft[0]
    }

    fmt.Println("Part 1", result1)
    fmt.Println("Part 2", result2)
}

