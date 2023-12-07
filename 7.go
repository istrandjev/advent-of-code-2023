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

    jokers := 0
    for _, c := range handString {
        if c == '*' {
            jokers++
            continue
        }
        _, exists := counter[c]
        if exists {
            counter[c]++
        } else {
            counter[c] = 1
        }
    }

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
    newHandStr := ""
    for _, r := range originalHand.handString {
        toAdd := r
        if r == 'J' {
            toAdd = '*'
        } 
        newHandStr = fmt.Sprintf("%s%c", newHandStr, toAdd)
    }
    return Hand{newHandStr, parseHandType(newHandStr), originalHand.amount, getCardStrengths(newHandStr)}
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
    sort.Sort(ByStrength(hands))

    answer1 := 0
    for index, hand := range hands {
        answer1 += hand.amount * (index + 1)
    }
    fmt.Println("Part 1", answer1)

    var hands2[] Hand
    for _, hand := range hands {
        hands2 = append(hands2, withJokers(hand))
    }
    sort.Sort(ByStrength(hands2))
    answer2 := 0
    for index, hand := range hands2 {
        answer2 += hand.amount * (index + 1)
    }
    fmt.Println("Part 2", answer2)
}

