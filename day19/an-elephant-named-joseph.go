package main

import (
    "fmt"
)

const input = 3005290
var elves [input]int

func main() {
    resetElves()
    fmt.Printf("Elf %d has all the presents for Part 1.\n", findElfWithAllPresentsPt1())

    fmt.Printf("Elf %d has all the presents for Part 2.\n", findElfWithAllPresentsPt2())
    // Let x = 3^z that is closest to input without exceeding the input (3^13)
    // If input > x*2, elf = ((input - (x * 2)) * 2) + x
    // else if input < x*2, elf = x - ((x * 2) - input)
    // else elf = x
}

func resetElves() {
    for i := 0; i < input; i++ {
        elves[i] = 1
    }
}

func findElfWithAllPresentsPt1() int {
    for i := 0; ; {
        if i >= input {
            i = 0
        }

        if elves[i] == input {
            return i + 1
        } else if elves[i] == 0 {
            i++
            continue
        }

        next := getNextElfWithPresentsPt1((i + 1) % input)
        elves[i] += elves[next]
        elves[next] = 0
        i = next
    }
}

func findElfWithAllPresentsPt2() (elf int) {
    x := 1
    for i := 0; ; i++ {
        if x*3 > input {
            break
        }
        x *= 3
    }

    doubleX := x * 2
    if input > doubleX {
        elf = ((input - doubleX) * 2) + x
    } else if input < doubleX {
        elf = x - doubleX + input
    } else {
        elf = x
    }

    return
}

func getNextElfWithPresentsPt1(start int) int {
    for {
        if elves[start] > 0 {
            return start
        } else {
            start = (start + 1) % input
        }
    }
}
