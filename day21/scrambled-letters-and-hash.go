package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    operations, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day21/data")
    check(err)

    input := "abcdefgh"
    arr := strings.Split(strings.TrimSpace(string(operations)), "\n")
    for _,op := range arr {
        i := strings.Split(op, " ")
        input = scramble(input, i)
    }

    fmt.Printf("Scrambled password is %q\n", input)

    input = "fbgdceah"
    for i := len(arr) - 1; i >= 0; i-- {
        op := strings.Split(arr[i], " ")
        input = unscramble(input, op)
    }

    fmt.Printf("Unscrambled password is %q\n", input)
}

func scramble(input string, i []string) (output string) {
    switch i[0] {
        case "swap":
            if i[1] == "position" {
                p1,_ := strconv.Atoi(i[2])
                p2,_ := strconv.Atoi(i[5])
                output = swapPosition(input, p1, p2)
            } else if i[1] == "letter" {
                output = swapLetter(input, strings.Index(input, i[2]), strings.Index(input, i[5]))
            }
        case "rotate":
            if len(i) == 4 {
                steps, err := strconv.Atoi(i[2])
                check(err)

                output = rotateDirection(input, i[1] == "right", steps % len(input))
            } else {
                output = rotateByIndex(input, strings.Index(input, i[6]), false)
            }
        case "reverse":
            start,_ := strconv.Atoi(i[2])
            end,_ := strconv.Atoi(i[4])
            output = reverse(input, start, end)
        case "move":
            from,_ := strconv.Atoi(i[2])
            to,_ := strconv.Atoi(i[5])
            output = move(input, from, to)
        default:
    }

    return
}

func unscramble(input string, i []string) (output string) {
    switch i[0] {
        case "swap":
            if i[1] == "position" {
                p1,_ := strconv.Atoi(i[2])
                p2,_ := strconv.Atoi(i[5])
                output = swapPosition(input, p1, p2)
            } else if i[1] == "letter" {
                output = swapLetter(input, strings.Index(input, i[2]), strings.Index(input, i[5]))
            }
        case "rotate":
            if len(i) == 4 {
                steps, err := strconv.Atoi(i[2])
                check(err)

                output = rotateDirection(input, i[1] != "right", steps % len(input))
            } else {
                output = rotateByIndex(input, strings.Index(input, i[6]), true)
            }
        case "reverse":
            start,_ := strconv.Atoi(i[2])
            end,_ := strconv.Atoi(i[4])
            output = reverse(input, start, end)
        case "move":
            from,_ := strconv.Atoi(i[2])
            to,_ := strconv.Atoi(i[5])
            output = move(input, to, from)
        default:
    }

    return
}

func swapPosition(s string, p1, p2 int) string {
    runes := []rune(s)
    runes[p1], runes[p2] = runes[p2], runes[p1]

    return string(runes)
}

func swapLetter(s string, i1, i2 int) string {
    return swapPosition(s, i1, i2)
}

func rotateDirection(s string, rotateR bool, steps int) string {
    if steps == len(s) {
        return s
    }

    runes := []rune(s)
    if rotateR {
        runes = append(runes[len(s) - steps:], runes[:len(s) - steps]...)
    } else {
        runes = append(runes[steps:], runes[0:steps]...)
    }
    return string(runes)
}

func rotateByIndex(s string, index int, reverse bool) string {
    // pos shift newpos
    //   0     1      1
    //   1     2      3
    //   2     3      5
    //   3     4      7
    //   4     6      2
    //   5     7      4
    //   6     8      6
    //   7     9      0
    if reverse {
        if (index % 2 == 1) || (index == 0) {
            index = (index / 2) + 1
        } else {
            index = (index / 2) + 5
        }
    } else {
        if index >= 4 {
            index++
        }
        index = (index + 1) % len(s)
    }
    return rotateDirection(s, !reverse, index)
}

func reverse(s string, start, end int) string {
    runes := []rune(s)
    sub := runes[start:end + 1]
    for i, j := 0, len(sub) - 1; i < j; i, j = i+1, j-1 {
        sub[i], sub[j] = sub[j], sub[i]
    }
    return string(append(runes[0:start], append(sub, runes[end + 1:]...)...))
}

func move(s string, from, to int) string {
    runes := []rune(s)
    f := runes[from]
    runes = append(runes[:from], runes[from + 1:]...)
    runes = append(runes[:to], append([]rune{f}, runes[to:]...)...)
    return string(runes)
}
