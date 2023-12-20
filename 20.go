package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
type Pulse struct {
    source string
    destination string
    isHigh bool
}

type Module interface {
    Process(Pulse) []Pulse
    AddInput(string)
    GetDestinations() []string
}

func GetPulses(isHigh bool, source string, destinations []string) []Pulse {
    result := []Pulse{}
    for _, d := range destinations {
        result = append(result, Pulse{source, d, isHigh})
    }
    return result
}

type FlipFlop struct {
    name string
    onOff bool
    destinations []string
}

func (f *FlipFlop)Process(p Pulse) []Pulse {
    result := []Pulse{}
    if p.isHigh != true {
        f.onOff = !f.onOff
        result = GetPulses(f.onOff, f.name, f.GetDestinations())
    }
    return result
}
func (f *FlipFlop)AddInput(_ string) {
}
func (f *FlipFlop)GetDestinations() []string {
    return f.destinations
}

type Conjunction struct {
    name string
    inputs map[string]bool
    destinations []string
}

func (c *Conjunction)Process(p Pulse) []Pulse {
    c.inputs[p.source] = p.isHigh
    allHigh := true
    for _, v :=  range c.inputs {
        if !v {
            allHigh = false
            break
        }
    }
    return GetPulses(!allHigh, c.name, c.GetDestinations())
}
func (c *Conjunction)AddInput(inp string) {
    c.inputs[inp] = false
}
func (c *Conjunction)GetDestinations() []string  {
    return c.destinations
}


type Broadcast struct {
    name string
    destinations []string
}
func (b *Broadcast)Process(p Pulse) []Pulse {
    return GetPulses(p.isHigh, b.name, b.GetDestinations())
}
func (b *Broadcast)AddInput(_ string) {
}
func (b *Broadcast)GetDestinations() []string  {
    return b.destinations
}


var modules = map[string]Module{}

func ParseAndAddModule(line string) {
    parts := strings.Split(line, " -> ")
    destinations := strings.Split(parts[1], ", ")
    if parts[0][0] == '&' {
        name := parts[0][1:]
        modules[name] = &Conjunction{name: name, inputs: map[string]bool{}, destinations: destinations}
    } else if parts[0][0] == '%' {
        name := parts[0][1:]
        modules[name] = &FlipFlop{name: name, destinations: destinations}
    } else {
        name := parts[0]
        if name != "broadcaster" {
            fmt.Println("Not broadcast", name)
        }
        modules[name] = &Broadcast{name: name, destinations: destinations}
    }
}

var interesting = map[string]int{"vq": 0, "rf": 0, "sr": 0, "sn": 0}

func ProcessPulse(p Pulse, pressNumber int) (int, int) {
    toProcess := []Pulse{p}
    high := 0
    low := 1
    current := 0
    for current < len(toProcess) {
        p := toProcess[current]
        current++
        module, exists := modules[p.destination]


        if !exists {
            continue
        }
        pulses := module.Process(p)
        if len(pulses) > 0 {
            value, exists := interesting[pulses[0].source]
            if exists && pulses[0].isHigh && value == 0 {
                interesting[pulses[0].source] = pressNumber
            }
            if pulses[0].isHigh {
                high += len(pulses)
            } else {
                low += len(pulses)
            }
        }
        toProcess = append(toProcess, pulses...)
    }
    return low, high
}
func Gcd(a int, b int) int {
    if a < b {
        return Gcd(b, a)
    }
    for a % b != 0 {
        r := a % b
        a = b
        b = r
    }
    return b
}
func Lcm(a int, b int) int {
    d := Gcd(a, b)
    return (a / d) * b
}
func GetAnswer2() int {
    res := 1
    for _, v := range interesting {
        res = Lcm(res, v)
    }
    return res
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        line := scanner.Text()
        ParseAndAddModule(line)
    }

    for name, module := range modules {
        destinations := module.GetDestinations()
        for _, d := range destinations {
            module, exists := modules[d]
            if exists {
                module.AddInput(name)
            }
            
        }
    }
    low1, high1 := 0, 0
    result1 := 0
    for i := 0; i < 20000; i++ {
        l, h := ProcessPulse(Pulse{"button", "broadcaster", false}, i + 1)
        low1 += l
        high1 += h
        if i == 999 {
            result1 = low1 * high1
        }
        allHit := true
        for _, v := range interesting {
            if v == 0 {
                allHit = false
            }
        }
        if allHit && i > 1000{
            break
        }
    }
    fmt.Println("Part 1", result1)
    result2 := GetAnswer2()
    fmt.Println("Part 2", result2)
}