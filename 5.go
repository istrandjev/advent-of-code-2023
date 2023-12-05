package main

import (
    "bufio"
    "fmt"
    "os"
	"sort"
    "strconv"
	"strings"
)

type RangeMapping struct {
    targetFrom int64
	sourceFrom int64
	length int64
}

type BySourceFrom []RangeMapping
func (a BySourceFrom) Len() int           { return len(a) }
func (a BySourceFrom) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a BySourceFrom) Less(i, j int) bool { return a[i].sourceFrom < a[j].sourceFrom }


func toInt(number_str string) int64 {
	value, _ := strconv.ParseInt(number_str, 10, 64)
	return value
}

func parseRangeMapping(line string) RangeMapping {
	tokens := strings.Split(line, " ")
	return RangeMapping{toInt(tokens[0]), toInt(tokens[1]), toInt(tokens[2])}
}

func parseMapping(scanner *bufio.Scanner) ([]RangeMapping, bool) {
	parsed := false
	var mapping []RangeMapping
	for scanner.Scan() {
		line := scanner.Text()
		line = scanner.Text()
		if len(line) == 0 {
			break
		}
		
		if line[len(line) - len("map:"):] == "map:" {
			parsed = true
			continue
		}
		if parsed {
			mapping = append(mapping, parseRangeMapping(line))
		}
	}
	return mapping, parsed
}

func getTarget(seed int64, mapping []RangeMapping) int64 {
	if seed < mapping[0].sourceFrom {
		return seed
	}
	beg := 0
	end := len(mapping)
	for end - beg > 1 {
		mid := (end + beg) / 2
		if seed >= mapping[mid].sourceFrom {
			beg = mid
		} else {
			end = mid
		}
	}
	if seed < mapping[beg].sourceFrom + mapping[beg].length {
		return seed - mapping[beg].sourceFrom + mapping[beg].targetFrom
	} else {
		return seed
	}
}

func mapSeed(seed int64, mappings [][]RangeMapping) int64 {
	for _, mapping := range mappings {
		seed = getTarget(seed, mapping)
	}
	return seed
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	seeds := []int64{}
	if scanner.Scan() {
		line := scanner.Text()
		for _, val_str := range strings.Split(line, " ")[1:] {
			seeds = append(seeds, toInt(val_str))
		}
	}
	scanner.Scan()  // Skip the empty line
	var mappings [][]RangeMapping
	for true {
		new_mapping, hasMore := parseMapping(scanner)
		if !hasMore {
			break
		}
		sort.Sort(BySourceFrom(new_mapping))
		mappings = append(mappings, new_mapping)
	}
	var answer1 int64 = -1
	for _, seed := range seeds {
		mapped_seed := mapSeed(seed, mappings)
		if answer1 == -1 || mapped_seed < answer1 {
			answer1 = mapped_seed
		}
	}	
	fmt.Println("Part 1", answer1)
	var answer2 int64 = -1
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i] + seeds[i + 1]; j++ {
			mapped_seed := mapSeed(j, mappings)
			if answer2 == -1 || mapped_seed < answer2 {
				answer2 = mapped_seed
			}
		}
	}
	fmt.Println("Part 2", answer2)
}