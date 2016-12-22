package main

import (
    "fmt"
    "io/ioutil"
    "sort"
    "strconv"
    "strings"
)

type Range struct {
    min, max int
}

var blockedRanges []Range
type ByRange []Range

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    ipranges, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day20/data")
    check(err)

    arr := strings.Split(strings.TrimSpace(string(ipranges)), "\n")
    ranges := parseRanges(arr)
    sort.Sort(ByRange(ranges))

    for _,r := range ranges {
        insertRange(r)
    }

    fmt.Printf("Lowest-valued IP is: %d\n", (blockedRanges[0]).max + 1)
    fmt.Printf("Number of IPs in the blacklist: %d\n", len(blockedRanges) - 1)
}

func parseRanges(arr []string) (ranges []Range) {
    for _,a := range arr {
        r := strings.Split(a, "-")
        min, err := strconv.Atoi(r[0])
        check(err)
        max, err := strconv.Atoi(r[1])
        check(err)
        ranges = append(ranges, Range{min, max})
    }

    return
}

func insertRange(r Range) {
    if len(blockedRanges) == 0 {
        blockedRanges = append(blockedRanges, r)
    } else {
        shouldMerge := search(r)
        last := len(blockedRanges) - 1
        if shouldMerge {
            blockedRanges[last] = mergeRange(r, blockedRanges[last])
        } else {
            blockedRanges = append(blockedRanges, r)
        }
    }
}

func search(r Range) bool {
    b := blockedRanges[len(blockedRanges) - 1]
    if r.min <= b.max + 1 {
        return true
    }

    return false
}

func mergeRange(r1 Range, r2 Range) Range {
    var min int
    if r1.min < r2.min {
        min = r1.min
    } else {
        min = r2.min
    }

    var max int
    if r1.max < r2.max {
        max = r2.max
    } else {
        max = r1.max
    }

    return Range{min, max}
}

func (s ByRange) Less(a, b int) bool {
    b1 := s[a]
    b2 := s[b]

    if b1.min == b2.min {
        return b1.max < b2.max
    }

    return b1.min < b2.min
}

func (s ByRange) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByRange) Len() int {
    return len(s)
}

