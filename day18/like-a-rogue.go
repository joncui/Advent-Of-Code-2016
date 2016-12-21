package main

import (
    "bytes"
    "fmt"
)

var trapLCR = []string{"..^", ".^^", "^..", "^^."}

func main() {
    input := ".^^.^^^..^.^..^.^^.^^^^.^^.^^...^..^...^^^..^^...^..^^^^^^..^.^^^..^.^^^^.^^^.^...^^^.^^.^^^.^.^^.^."
    fmt.Printf("Number of safe tiles with %d rows: %d\n", 40, getSafeTileNum(input, 40))
    fmt.Printf("Number of safe tiles with %d rows: %d\n", 400000, getSafeTileNum(input, 400000))
}

func getSafeTileNum(input string, rows int) (safe int) {
    prev := input
    safe = 43
    for i := 1; i < rows; i++ {
        var buffer bytes.Buffer
        for x := 0; x < len(prev); x++ {
            if isTrap(getLCR(x, prev)) {
                buffer.WriteString("^")
            } else {
                buffer.WriteString(".")
                safe++
            }
        }
        prev = buffer.String()
    }

    return
}

func getLCR(i int, row string) (lcr string) {
    if i == 0 {
        lcr = "." + string(row[i]) + string(row[i + 1])
    } else if i == len(row) - 1 {
        lcr = string(row[i - 1]) + string(row[i]) + "."
    } else {
        lcr = string(row[i - 1]) + string(row[i]) + string(row[i + 1])
    }

    return
}

func isTrap(lcr string) bool {
    for _,v := range trapLCR {
        if lcr == v {
            return true
        }
    }
    return false
}
