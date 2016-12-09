package main
import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

var keypad1 = [3][3]string{[3]string{"1", "2", "3"}, [3]string{"4", "5", "6"}, [3]string{"7", "8", "9"}}
var keypad2 = [5][5]string{[5]string{"x", "x", "1", "x", "x"}, [5]string{"x", "2", "3", "4", "x"}, [5]string{"5", "6", "7", "8", "9"}, [5]string{"x", "A", "B", "C", "x"}, [5]string{"x", "x", "D", "x", "x"}}
var part string
var max int

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    if part = "1"; len(os.Args) > 1 {
        part = os.Args[1]
    }

    instructions, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day2/data")
    check(err)
    arr := strings.Split(string(instructions), "\n")
    code := ""

    var x, y int
    if part == "1" {
        max = 2
        x, y = 1, 1
    } else {
        max = 4
        x, y = 0, 2
    }

    fmt.Printf("%d, %d\n", x, y)

    for _,v := range arr {
        if len(v) > 0 {
            x, y = getNextButton(v, x, y)
            fmt.Printf("%d, %d\n", x, y)
            if part == "1" {
                code += keypad1[y][x]
            } else {
                code += keypad2[y][x]
            }
        }
    }

    fmt.Println(code)
}

func getNextButton(instruction string, x, y int) (int, int) {
    for _,direction := range instruction {
        x, y = moveButton(string(direction), x, y)
    }

    return x, y
}

func moveButton(direction string, x, y int) (int, int) {
    multiplier := getMultiplier(direction)
    axis := getAxis(direction, &x, &y)
    *axis += multiplier

    // Fix if out of bounds
    if *axis < 0 {
        *axis = 0
    } else if *axis > max {
        *axis = max
    }

    // Fix for pt2 if coords do not point to a key
    if (part == "2") && (keypad2[y][x] == "x") {
        *axis -= multiplier
    }

    return x, y
}

func getMultiplier(direction string) (int) {
    if (direction == "U") || (direction == "L") {
        return -1
    } else {
        return 1
    }
}

func getAxis(direction string, x, y *int) (*int) {
    if (direction == "U") || (direction == "D") {
        return y
    } else {
        return x
    }
}
