package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
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

type Part struct {
    values [4]int
}
func (p *Part) SumOfRatings() int {
    result := 0
    for _, v := range p.values {
        result += v
    }
    return result
}
type Rule struct {
    variableCode int
    value int
    greater bool
    target string
}
var VARIABLE_CODES = map[string]int {
    "x": 0,
    "m": 1,
    "a": 2,
    "s": 3,
}
func (r *Rule) Matches(p Part) bool {
    value := p.values[r.variableCode]
    matches := true
    if r.greater {
        matches = value > r.value
    } else {
        matches = value < r.value
    }
    return matches
}
func (r *Rule) Apply(p Part) string {
    if r.Matches(p) {
        return r.target
    } else {
        return ""
    }
}

type Workflow struct {
    rules []Rule
    defaultTarget string
    name string
}
func (w *Workflow) Run(p Part) string {
    for _, r := range w.rules {
        result := r.Apply(p)
        if len(result) > 0 {
            return result
        }
    }
    return w.defaultTarget
}

func ParseRule(ruleString string) Rule {
    ruleRegex := regexp.MustCompile(`(\w+)(<|>)(\d+):(\w+)`)
    match := ruleRegex.FindStringSubmatch(ruleString)

    greater := match[2] == ">"
    return Rule{VARIABLE_CODES[match[1]], toInt(match[3]), greater, match[4]}
}

func IsWorkflowConst(workflow Workflow) bool {
    for _, rule := range workflow.rules {
        if rule.target != workflow.defaultTarget {
            return false
        }
    }
    return true
}
func ParseWorkflow(line string) Workflow {
    reWorkflow := regexp.MustCompile(`(\w+)\{(.+),(\w+)\}`)
    match := reWorkflow.FindStringSubmatch(line)
    wfName := match[1]
    rulesString := match[2]
    defaultTarget := match[3]
    rules := []Rule{}
    for _, ruleString := range strings.Split(rulesString, ",") {
        rules = append(rules, ParseRule(ruleString))
    }

    workflow := Workflow{rules, defaultTarget, wfName}
    if IsWorkflowConst(workflow) {
        workflow = Workflow{[]Rule{}, defaultTarget, wfName}
    }
    
    return workflow
}
func ParsePart(line string) Part {
    rePart := regexp.MustCompile(`x=(\d+),m=(\d+),a=(\d+),s=(\d+)`)
    match := rePart.FindStringSubmatch(line)
    values := [4]int{}
    for i := range "xmas" {
        values[i] = toInt(match[i + 1])
    }
    return Part{values}
}

type State struct {
    node string
    limits [4][2]int
}
var mem map[State]int
var globalWorkflows map[string]Workflow

func Intersect(value int, greater bool, limits [2]int) [2]int {
    if greater {
        if limits[0] > value {
            return [2]int{limits[0], limits[1]}
        } else if limits[1] <= value {
            return [2]int{-1, -1}
        } else {
            return [2]int{value + 1, limits[1]}
        }
    } else {
        if limits[1] < value {
            return [2]int{limits[0], limits[1]}
        } else if limits[0] >= value {
            return [2]int{-1, -1}
        } else {
            return [2]int{limits[0], value - 1}
        }
    }
}

func Exclude(value int, greater bool, limits [2]int) [2]int {
    if greater {
        return Intersect(value + 1, !greater, limits)
    } else {
        return Intersect(value - 1, !greater, limits)
    }
}

func CopyLimits(limits [4][2]int) [4][2]int {
    result := [4][2]int{}
    for i, r := range limits {
        for j, v := range r {
            result[i][j] = v
        }
    }
    return result
}

func Dp(state State) int {
    memoized, exists := mem[state]
    if exists {
        return memoized
    }
    if state.node == "R" {
        mem[state] = 0
        return 0
    }
    if state.node == "A" {
        res := 1
        for _, lim := range state.limits {
            res *= lim[1] - lim[0] + 1
        }
        mem[state] = res
        return res
    }
    workflow := globalWorkflows[state.node]
    result := 0
    localLimits := CopyLimits(state.limits)
    possible := true
    for _, rule := range workflow.rules {
        idx := rule.variableCode
        temp := Intersect(rule.value, rule.greater, localLimits[idx])
        if temp[0] != -1 {
            newLimits := CopyLimits(localLimits)
            newLimits[idx] = temp
            next := State{rule.target, newLimits}
            result += Dp(next)
        }
        temp = Exclude(rule.value, rule.greater, localLimits[idx])
        if temp[0] == -1 {
            possible = false
            break
        }
        localLimits[idx] = temp
    }
    if possible {
        result += Dp(State{workflow.defaultTarget, localLimits})
    }
    mem[state] = result
    return result
}
func SolvePart2(workflows map[string]Workflow) int {
    mem = map[State]int{}
    globalWorkflows = workflows
    limits := [4][2]int{
        {1, 4000},
        {1, 4000},
        {1, 4000},
        {1, 4000},
    }
    return Dp(State{"in", limits})
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    workflows := map[string]Workflow{}
    parts := []Part{}
    for scanner.Scan() {
        line := scanner.Text()
        if len(line) == 0 {
            continue
        }
        if strings.ContainsRune(line, '=') {
            parts = append(parts, ParsePart(line))
        } else {
            workflow := ParseWorkflow(line)
            workflows[workflow.name] = workflow
        }
    }
    result1 := 0
    for _, part := range parts {
        wf := "in"
        for wf != "R" && wf != "A" {
            currentWf := workflows[wf] 
            wf = currentWf.Run(part)
        }
        if wf == "A" {
            result1 += part.SumOfRatings()
        }
    }
    fmt.Println("Part 1", result1)
    result2 := SolvePart2(workflows)
    fmt.Println("Part 2", result2)
}