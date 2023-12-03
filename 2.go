package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
	"strings"
)

type Draw struct {
    number int
    colour string
}

func parseDraw(draw string) []Draw {
	draws := []Draw{}
	re := regexp.MustCompile(`(\d+) ([a-z]+)`)
	for _, match := range re.FindAllStringSubmatch(draw, -1) {
		cnt, _ := strconv.Atoi(match[1])
		draws = append(draws, Draw{cnt, match[2]})
	}
	return draws
}

func parseGroup(groupLine string) [][]Draw {
	var group_draws [][]Draw

	groups := strings.Split(groupLine, ":")[1]

	draw_strings := strings.Split(groups, ";")
	for _, draw_string := range draw_strings {
		group_draws = append(group_draws, parseDraw(draw_string))
	}
	return group_draws
}

func groupIsGoodPart1(group [][]Draw) bool {
	bag := map[string]int{
        "red":   12,
        "green": 13,
        "blue":  14,
    }

	for _, draw_data := range group {
		for _, ball_draw := range draw_data {
			if cube_limit, ok := bag[ball_draw.colour]; ok {
				if cube_limit < ball_draw.number {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
func getGroupPower(group [][]Draw) int {
	bag := map[string]int{}

	for _, draw_data := range group {
		for _, ball_draw := range draw_data {
			if cube_limit, ok := bag[ball_draw.colour]; ok {
				if cube_limit < ball_draw.number {
					bag[ball_draw.colour] = ball_draw.number
				}
			} else {
				bag[ball_draw.colour] = ball_draw.number
			}
		}
	}

	result := 1
	for _, value := range bag {
		result *= value
	}
	return result
	
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	result1 := 0
	result2 := 0
	group_index := 1
	for scanner.Scan() {
		line := scanner.Text()
        group_draws := parseGroup(line)
		if groupIsGoodPart1(group_draws) {
			result1 += group_index
		}
		result2 += getGroupPower(group_draws)
		group_index += 1
	}
	fmt.Println("Part 1", result1)
	fmt.Println("Part 2", result2)
}