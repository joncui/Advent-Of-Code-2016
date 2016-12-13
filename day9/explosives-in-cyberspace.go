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
    data, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day9/test")
    check(err)

    dataString := strings.TrimSpace(string(data))

    // count1 := partOne(dataString)
    count2 := partTwo(dataString)

    // fmt.Printf("Decompressed length pt 1: %d\n", count1)
    fmt.Printf("Decompressed length pt 2: %d\n", count2)
}

func partOne(data string) (count int) {
    markerStart := strings.Index(data, "(")
    markerEnd := strings.Index(data, ")")
    if (markerStart == -1) || (markerEnd == -1) {
        count += len(data)
    } else {
        if markerStart != 0 {
            count += markerStart
        }

        x, y := parseMarker(data[markerStart + 1:markerEnd])
        count += (x * y) + partOne(data[markerEnd + x + 1:])
    }

    return
}

func partTwo(data string) (count int) {
    markerStart := strings.Index(data, "(")
    markerEnd := strings.Index(data, ")")
    if (markerStart == -1) || (markerEnd == -1) {
        count += len(data)
    } else {
        if markerStart != 0 {
            count += markerStart
        }

        x, y := parseMarker(data[markerStart + 1:markerEnd])
        fmt.Printf("%s <> %d => %d, %d\n", data, markerEnd, x, y)
        count += (partTwo(data[markerEnd+1:markerEnd+x+1]) * y) + (partTwo(data[markerEnd + x + 1:]))
    }

    return
}

func parseMarker(marker string) (x, y int) {
    splitMarker := strings.Split(marker, "x")
    x,_ = strconv.Atoi(splitMarker[0])
    y,_ = strconv.Atoi(splitMarker[1])

    return
}
