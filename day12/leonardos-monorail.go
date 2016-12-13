package main

import (
    "fmt"
    "io/ioutil"
    "strconv"
    "strings"
)

var registers map[string]int

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day12/data")
    check(err)

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    var instructions [][]string

    for _, line := range lines {
        instructions = append(instructions, strings.Split(line, " "))
    }

    fmt.Printf("Register A is %d for Part 1\n", partOne(instructions))
    fmt.Printf("Register A is %d for Part 2\n", partTwo(instructions))
}

func partOne(instructions [][]string) int {
    registers = make(map[string]int)
    performInstructions(instructions)
    return registers["a"]
}

func partTwo(instructions [][]string) int {
    registers = map[string]int{"c": 1}
    performInstructions(instructions)
    return registers["a"]
}

func performInstructions(instructions [][]string) {
    for i := 0; i < len(instructions); i++ {
        switch instructions[i][0] {
            case "cpy":
                performCopy(instructions[i][1], instructions[i][2])
            case "inc":
                performInc(instructions[i][1])
            case "dec":
                performDec(instructions[i][1])
            case "jnz":
                i += performJnz(instructions[i][1], instructions[i][2])
            default:
                panic("Unrecognized instruction!")
        }
    }
}

func performCopy(x, y string) {
    val := strToInt(x)
    registers[y] = val
}

func performInc(x string) {
    registers[x]++
}

func performDec(x string) {
    registers[x]--
}

func performJnz(x, y string) int {
    val := strToInt(x)
    if val != 0 {
        return strToInt(y) - 1
    }

    return 0
}

func strToInt(s string) int {
    num, err := strconv.Atoi(s)
    if err != nil {
        return registers[s]
    }

    return num
}
