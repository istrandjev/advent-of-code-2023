package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
)

type NodeKids struct {
	left int
	leftStr string
	right int
	rightStr string
}

func parseNode(nodeLine string) (string, string, string) {
    re := regexp.MustCompile(`([a-zA-Z]+) = \(([a-zA-Z]+), ([a-zA-Z]+)\)`)
	match := re.FindStringSubmatch(nodeLine)
    return match[1], match[2], match[3]
}

var nodeCodes map[string]int
func getNodeCode(nodeString string) int {
	code, exists := nodeCodes[nodeString]
	if exists {
		return code
	}
	result := len(nodeCodes)
	nodeCodes[nodeString] = result
	return result
}

func solvePart1(graph map[int]NodeKids, current int, targetNodes map[int]bool, instructions string) int {
	currentInstruction := 0
	steps1 := 0
	for true {
		_, isTarget := targetNodes[current]
		if isTarget {
			break
		}
		r := instructions[currentInstruction]
		currentInstruction++
		currentInstruction %= len(instructions)

		currentNode, _ := graph[current]
		if r == 'R' {
			current = currentNode.right
		} else {
			current = currentNode.left
		}
		steps1++
	}
	return steps1
}

func gcd(a int, b int) int {
	if a < b {
		return gcd(b, a)
	}
	for a % b != 0 {
		c := a % b
		a = b
		b = c
	}
	return b
}
func lcm(a int, b int) int {
	d := gcd(a, b)
	return (a / d) * b
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)
	instructions := ""
	if scanner.Scan() {
		instructions = scanner.Text()
	}
	scanner.Scan()  // skip the empty line
	nodeCodes = map[string]int{}
	graph := map[int]NodeKids{}
	startNodes := []int{}
	endNodes := map[int]bool{}
    for scanner.Scan() {
		line := scanner.Text()
		nodeName, leftNode, rightNode := parseNode(line)
		if nodeName[2] == 'A' {
			startNodes = append(startNodes, getNodeCode(nodeName))
		}
		if nodeName[2] == 'Z' {
			endNodes[getNodeCode(nodeName)] = true
		}
		graph[getNodeCode(nodeName)] = NodeKids{getNodeCode(leftNode), leftNode, getNodeCode(rightNode), rightNode}
    }
	current := getNodeCode("AAA")
	target := getNodeCode("ZZZ")
	targetNodes1 := map[int]bool{target: true}
	steps1 := solvePart1(graph, current, targetNodes1, instructions)
    fmt.Println("Part 1", steps1)
	result2 := 1
	for _, sn := range startNodes {
		steps := solvePart1(graph, sn, endNodes, instructions)
		result2 = lcm(result2, steps)
	}
    fmt.Println("Part 2", result2)
}


