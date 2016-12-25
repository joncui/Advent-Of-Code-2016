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
    data, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day23/data")
    check(err)

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    var instructions [][]string

    for _, line := range lines {
        instructions = append(instructions, strings.Split(line, " "))
    }

    registers = make(map[string]int)
    registers["a"] = 12
    fmt.Printf("Register A is %d\n", performInstructions(instructions))
}

func performInstructions(instructions [][]string) int {
    for i := 0; i < len(instructions); i++ {
        switch instructions[i][0] {
            case "cpy":
                performCopy(instructions[i][1], instructions[i][2])
            case "inc":
                if !performPeephole(instructions, i) {
                    performInc(instructions[i][1])
                }
            case "dec":
                performDec(instructions[i][1])
            case "jnz":
                i += performJnz(instructions[i][1], instructions[i][2])
            case "tgl":
                val := i + strToInt(instructions[i][1])
                if val < len(instructions) {
                    newI := performTgl(instructions[val])
                    instructions[val][0] = newI
                }
            default:
                panic("Unrecognized instruction!")
        }
    }

    return registers["a"]
}

func performPeephole(instructions [][]string, i int) bool {
    // TODO
    return false
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

func performTgl(toggledInstruction []string) (newI string) {
    if len(toggledInstruction) == 2 {
        if toggledInstruction[0] == "inc" {
            newI = "dec"
        } else {
            newI = "inc"
        }
    } else if len(toggledInstruction) == 3 {
        if toggledInstruction[0] == "jnz" {
            if _, err := strconv.Atoi(toggledInstruction[2]); err != nil {
                newI = "cpy"
            }
        } else {
            newI = "jnz"
        }
    }

    return
}

func strToInt(s string) int {
    num, err := strconv.Atoi(s)
    if err != nil {
        return registers[s]
    }

    return num
}
