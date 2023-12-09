package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func toInt(number_str string) int {
    value, _ := strconv.Atoi(number_str)
    return value
}

type Hand struct {
    handString string
    handType string
    amount int
    cardStrengths []int
}

func parseHandType(handString string) string {
    counter := map[rune]int{}

    for _, c := range handString {
        _, exists := counter[c]
        if exists {
            counter[c]++
        } else {
            counter[c] = 1
        }
    }
    jokers := counter['*']
    delete(counter, '*')
    if jokers == 5 {
        return "5"
    }

    vals := []int{}
    for _, v := range counter {
        vals = append(vals, v)
    }
    sort.Ints(vals)
    vals[len(vals) - 1] += jokers
    result := ""
    for _, v := range vals {
        result += strconv.Itoa(v)
    }
    return result
}

var HAND_STRENGTHS map[string]int
var CARD_STRENGTHS map[rune]int

func getCardStrengths(handString string) []int {
    res := []int{}
    for _, r := range handString {
        strength, _ := CARD_STRENGTHS[r]
        res = append(res, strength)
    }
    return res
}

type ByStrength []Hand
func (a ByStrength) Len() int           { return len(a) }
func (a ByStrength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStrength) Less(i, j int) bool { 
    strength1, _ := HAND_STRENGTHS[a[i].handType] 
    strength2, _ := HAND_STRENGTHS[a[j].handType]
    if strength1 != strength2 {
        return strength1 < strength2
    }

    for ix := 0; ix < len(a[i].cardStrengths); ix++ {
        if a[i].cardStrengths[ix] != a[j].cardStrengths[ix] {
            return a[i].cardStrengths[ix] < a[j].cardStrengths[ix]
        }
    }
    return false
 }

func withJokers(originalHand Hand) Hand {
    newHandStr := strings.Replace(originalHand.handString, "J", "*", -1)
    return Hand{newHandStr, parseHandType(newHandStr), originalHand.amount, getCardStrengths(newHandStr)}
}

func getHandsScore(hands []Hand) int {
    sort.Sort(ByStrength(hands))

    answer := 0
    for index, hand := range hands {
        answer += hand.amount * (index + 1)
    }
    return answer
}

func mapFn[V any](collection []V, fn func(V)V) []V {
	result := []V{}
	for _, val := range collection {
		result = append(result, fn(val))
	}
	return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    HAND_STRENGTHS = map[string]int{
        "11111":   1,
        "1112":   2,
        "122": 3,
        "113":  4,
        "23":  5,
        "14":   6,
        "5": 7,
    }

    CARD_STRENGTHS = map[rune]int{
        '*': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14,
    }


    var hands []Hand
    for scanner.Scan() {
        tokens := strings.Fields(scanner.Text())
        hands = append(hands, Hand{tokens[0], parseHandType(tokens[0]), toInt(tokens[1]), getCardStrengths(tokens[0])})
    }
    fmt.Println("Part 1", getHandsScore(hands))
	hands2 := append([]Hand{}, mapFn(hands, withJokers)...)
    fmt.Println("Part 2", getHandsScore(hands2))
}

