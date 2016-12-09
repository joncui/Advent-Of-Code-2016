package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const width = 50
const height = 6
var screen [height][width]bool

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, err := os.Open("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day8/data")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lit := 0

    for scanner.Scan() {
        instructions := strings.Split(scanner.Text(), " ")

        if len(instructions) == 2 {
            dimensions := strings.Split(instructions[1], "x")
            lit += turnOnRect(dimensions[0], dimensions[1])
        } else {
            index,_ := strconv.Atoi(strings.Split(instructions[2], "=")[1])
            shift,_ := strconv.Atoi(instructions[4])
            if instructions[1] == "row" {
                for i := 0; i < shift; i++ {
                    copy(screen[index][:], append(screen[index][width - 1:], screen[index][0:width - 1]...))
                }
            } else {
                column := getColumn(index)
                for i := 0; i < shift; i++ {
                    column = append(column[height - 1:], column[0:height - 1]...)
                }
                setColumn(index, column)
            }
        }
    }

    fmt.Printf("Number of pixels lit: %d\n", lit)
    printGrid()
}

func turnOnRect(dimX, dimY string) (count int) {
    x,_ := strconv.Atoi(dimX)
    y,_ := strconv.Atoi(dimY)

    for i := 0; i < y; i++ {
        for j := 0; j < x; j++ {
            count += turnOn(j, i)
        }
    }

    return
}

func turnOn(x, y int) (count int) {
    if !screen[y][x] {
        count = 1
    }
    screen[y][x] = true

    return
}

func getColumn(index int) (column []bool) {
    for i := 0; i < 6; i++ {
        column = append(column, screen[i][index])
    }

    return column
}

func setColumn(index int, newColumn []bool) {
    for i := 0; i < 6; i++ {
        screen[i][index] = newColumn[i]
    }
}

func printGrid() {
    for i := 0; i < 6; i++ {
        for j := 0; j < 50; j++ {
            if screen[i][j] {
                fmt.Print("+")
            } else {
                fmt.Print("-")
            }
        }
        fmt.Print("\n")
    }
}
