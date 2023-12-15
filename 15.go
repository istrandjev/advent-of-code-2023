package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

type Lens struct {
    label string
    value int
}

func toInt(number_str string) int {
    value, _ := strconv.Atoi(number_str)
    return value
}

func Hash(s string) int {
    result := 0
    for _, c := range s {
        result = ((result + int(c)) * 17) % 256
    }
    return result
}

func ParseCommand(command string) (string, int) {
    reEquals := regexp.MustCompile(`([a-zA-Z]+)=(\d+)`)
    match := reEquals.FindStringSubmatch(command)
    if len(match) != 0 {
        return match[1], toInt(match[2])
    }
    reDash := regexp.MustCompile(`([a-zA-Z]+)-`)
    match = reDash.FindStringSubmatch(command)
    return match[1], -1
}

func ProcessCommand(boxes [][]Lens, label string, value int) {
    boxNumber := Hash(label)

    idx := -1
    for i, box := range boxes[boxNumber] {
        if box.label == label {
            idx = i
        }
    }

    if value == -1 {
        if idx != -1 {
            boxes[boxNumber] = append(boxes[boxNumber][:idx], boxes[boxNumber][idx + 1:]...)
        }
    } else {
        if idx == -1 {
            boxes[boxNumber] = append(boxes[boxNumber], Lens{label, value})
        } else {
            boxes[boxNumber][idx] = Lens{label, value}
        }
    }
}

func ComputeScore2(boxes [][]Lens) int {
    result := 0
    for i, box := range boxes {
        for j, lens := range box {
            result += (i + 1) * (j + 1) * lens.value
        }
    }
    return result
 }
 
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    result1 := 0
    var boxes [][]Lens
    for i := 0; i < 256; i++ {
        boxes = append(boxes, []Lens{})
    }

    for scanner.Scan() {
        line := scanner.Text()
        tokens := strings.Split(line, ",")
        for _, t := range tokens {
            result1 += Hash(t)
            label, value := ParseCommand(t)
            ProcessCommand(boxes, label, value)
        }
        
    }

    fmt.Println("Part 1", result1)
    result2 := ComputeScore2(boxes)
    fmt.Println("Part 2", result2)
}