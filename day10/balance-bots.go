package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
    "strconv"
)

type Bot struct {
    LowPath, HighPath int
    LowTarget, HighTarget string
    Values []int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, err := os.Open("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day10/data")
    check(err)

    botMap := make(map[int]Bot)
    var leaf []int
    outputs := make(map[int]int)
    var partOne int

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), " ")
        if line[0] == "value" {
            value, _ := strconv.Atoi(line[1])
            bot, _ := strconv.Atoi(line[5])
            if val, ok := botMap[bot]; ok {
                val.Values = append(val.Values, value)
                botMap[bot] = val
                if len(val.Values) == 2 {
                    leaf = append(leaf, bot)
                }
            } else {
                botMap[bot] = Bot{
                    Values: []int{value},
                }
            }
        } else {
            bot, _ := strconv.Atoi(line[1])
            low, _ := strconv.Atoi(line[6])
            high, _ := strconv.Atoi(line[11])
            lowTarget := line[5]
            highTarget := line[10]

            if val, ok := botMap[bot]; ok {
                val.LowPath = low
                val.HighPath = high
                val.LowTarget = lowTarget
                val.HighTarget = highTarget
                botMap[bot] = val
            } else {
                botMap[bot] = Bot{
                    low, high, lowTarget, highTarget, []int{},
                }
            }
        }
    }

    for {
        if val, ok := botMap[leaf[0]]; (ok && len(val.Values) == 2) {
            sort.Ints(val.Values)
            if (val.Values[0] == 17) && (val.Values[1] == 61) {
                partOne = leaf[0]
            }

            if val.LowTarget == "bot" {
                low := botMap[val.LowPath]
                low.Values = append(low.Values, val.Values[0])

                botMap[val.LowPath] = low

                if len(low.Values) == 2 {
                    leaf = append(leaf, val.LowPath)
                }
            } else {
                outputs[val.LowPath] = val.Values[0]
            }

            if val.HighTarget == "bot" {
                high := botMap[val.HighPath]
                high.Values = append(high.Values, val.Values[1])

                botMap[val.HighPath] = high

                if len(high.Values) == 2 {
                    leaf = append(leaf, val.HighPath)
                }
            } else {
                outputs[val.HighPath] = val.Values[1]
            }

            leaf = append(leaf[:0], leaf[1:]...)
        }

        if len(leaf) == 0 {
            break
        }
    }

    fmt.Printf("Bot %d compares 61 with 17\n", partOne)
    fmt.Printf("Multiply 0, 1, and 2 is %d\n", outputs[0] * outputs[1] * outputs[2])
}
