package main

import (
    "bufio"
    "fmt"
    "math/big"
    "os"
    "strconv"
    "strings"
)

func toInt64(number_str string) int64 {
    value, _ := strconv.ParseInt(number_str, 10, 64)
    return value
}

func Mul(a *big.Rat, b *big.Rat) *big.Rat {
    res := big.NewRat(0, 1)
    res.Mul(a, b)
    return res
}

func Sub(a *big.Rat, b *big.Rat) *big.Rat {
    res := big.NewRat(0, 1)
    res.Sub(a, b)
    return res
}

func Add(a *big.Rat, b *big.Rat) *big.Rat {
    res := big.NewRat(0, 1)
    res.Add(a, b)
    return res
}

type Line struct {
    A *big.Rat
    B *big.Rat
    C *big.Rat
}
func GetLine(p []int64, v []int64) Line {
    x0, x1 := big.NewRat(p[0], 1), big.NewRat(p[0] + v[0], 1)
    y0, y1 := big.NewRat(p[1], 1), big.NewRat(p[1] + v[1], 1)

    return Line{
        Sub(y0, y1),
        Sub(x1, x0),
        Sub(Mul(x0, y1), Mul(y0, x1)),
    }
}

func Intersect(l1 Line, l2 Line) (bool, *big.Rat, *big.Rat) {
    denom := Sub(Mul(l1.B, l2.A), Mul(l1.A, l2.B))
    numY := Sub(Mul(l1.A, l2.C), Mul(l1.C, l2.A))
    if denom.Cmp(big.NewRat(0, 1)) == 0 {
        return false, big.NewRat(0, 1), big.NewRat(0, 1)
    }
    denom = denom.Inv(denom)
    y := Mul(numY, denom)
    numX := Sub(Mul(l1.C, l2.B), Mul(l1.B, l2.C))
    x := Mul(numX, denom)
    return true, x, y
}

type Stone struct {
    p [3]int64
    v [3]int64
}
func ParseTriple(part string) [3]int64 {
    result := [3]int64{}
    for i, x := range strings.Split(part, ", ") {
        x = strings.Trim(x, " ")
        result[i] = toInt64(x)
    }
    return result
}
func ParseStone(line string) Stone {
    parts := strings.Split(line, "@")
    coords := ParseTriple(parts[0])
    vs := ParseTriple(parts[1])
    return Stone{coords, vs}
}

func GetTAlongAxis(start int64, vInt64 int64, value *big.Rat) (bool, *big.Rat) {
    if vInt64 == 0 {
        return false, new(big.Rat).SetInt64(0)
    }
    v := new(big.Rat).SetInt64(vInt64)
    v = v.Inv(v)
    s := new(big.Rat).SetInt64(start)
    delta := Sub(value, s)
    return true, Mul(delta, v)
}
func GetT(s Stone, x *big.Rat, y *big.Rat) *big.Rat {
    byX, res := GetTAlongAxis(s.p[0], s.v[0], x)
    if !byX {
        _, res = GetTAlongAxis(s.p[1], s.v[1], y)
    }
    return res
}
func IsGood(r *big.Rat, minv int64, maxv int64) bool {
    lowerLimit := new(big.Rat).SetInt64(minv)
    upperLimit := new(big.Rat).SetInt64(maxv)

    return r.Cmp(lowerLimit) >= 0 && r.Cmp(upperLimit) <= 0
}
func CountIntersections(stones []Stone, minv int64, maxv int64) int {
    result := 0
    zero := new(big.Rat).SetInt64(0)
    for i, s1 := range stones {
        for j := i + 1; j < len(stones); j++ {
            s2 := stones[j]
            l1 := GetLine(s1.p[:2], s1.v[:2])
            l2 := GetLine(s2.p[:2], s2.v[:2])
            exists, x, y := Intersect(l1, l2)
            t1 := GetT(s1, x, y)
            t2 := GetT(s2, x, y)
            if exists && IsGood(x, minv, maxv) && IsGood(y, minv, maxv) && t1.Cmp(zero) >= 0 && t2.Cmp(zero) >= 0 {
                result++
            }
        }
    }
    return result
}

func GetT1AndT0(s0 Stone, s1 Stone, vx int64, vy int64) (bool, int64, int64) {
    deltaX := s0.p[0] - s1.p[0]
    deltaY := s0.p[1] - s1.p[1]
    deltaVx0 := s0.v[0] - vx
    deltaVx1 := s1.v[0] - vx
    deltaVy0 := s0.v[1] - vy
    deltaVy1 := s1.v[1] - vy
    denom := deltaVy1 * deltaVx0 - deltaVy0 * deltaVx1
    num := deltaY * deltaVx0 - deltaVy0 * deltaX
    if denom == 0 {
        if num != 0 {
            return false, 0, 0
        } else {
            return true, 0, 0
        }
    }
    if num % denom != 0 {
        return false, 0, 0
    }
    t1 := num / denom
    t0Num := deltaVx1 * t1 - deltaX
    t0Denom := deltaVx0
    if t0Num % t0Denom != 0 {
        return false, 0, 0
    }
    t0 := t0Num / t0Denom
    return true, t0, t1
}
func GetPossibleTimes(s0 Stone, s1 Stone, vx int64, vy int64) map[int64]bool {
    good, t0, _ :=  GetT1AndT0(s0, s1, vx, vy)
    if !good {
        return map[int64]bool{}
    }
    return map[int64]bool{t0: true}
}
func GetVz(s0 Stone, s1 Stone, vx int64, vy int64, t0 int64) int64 {
    _, t0, t1 := GetT1AndT0(s0, s1, vx, vy)
    deltaZ := s0.p[2] - s1.p[2]
    denomZ := t1 - t0
    numZ := -deltaZ + t1 * s1.v[2] - t0 * s0.v[2]
    return numZ / denomZ
}
func SolvePart2(stones []Stone) int64 {
    limit := int64(512)
    common := map[[2]int64]bool{}
    for i := 0; i < 5; i++ {
        newCommon := map[[2]int64]bool{}
        for vx := -limit; vx <= limit; vx++ {
            for vy := -limit; vy <= limit; vy++ {
                possible := GetPossibleTimes(stones[i], stones[i + 1], vx, vy)

                if len(possible) > 0 {
                    newCommon[[2]int64{vx, vy}] = true
                }
            }
        }
        if i == 0 {
            common = newCommon
        } else {
            for k := range common {
                if _, exists := newCommon[k]; !exists {
                    delete(common, k)
                }
            }
        }
    }
    vx, vy := int64(0), int64(0)
    for k := range common {
        vx, vy = k[0], k[1]
    }
    possible := GetPossibleTimes(stones[0], stones[1], vx, vy)
    t0 := int64(0)
    for k := range possible {
        t0 = k
    }
    vz := GetVz(stones[0], stones[1], vx, vy, t0)
    x := stones[0].p[0] + t0 * (stones[0].v[0] - vx)
    y := stones[0].p[1] + t0 * (stones[0].v[1] - vy)
    z := stones[0].p[2] + t0 * (stones[0].v[2] - vz)

    return x + y + z
}
func main() {
    scanner := bufio.NewScanner(os.Stdin)

    stones := []Stone{}
    for scanner.Scan() {
        line := scanner.Text()
        stones = append(stones, ParseStone(line))
    }

    result1 := CountIntersections(stones, 200000000000000, 400000000000000)
    fmt.Println("Part 1", result1)
    result2 := SolvePart2(stones)
    fmt.Println("Part 2", result2)
}