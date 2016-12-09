package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var maps [8]map[string]int

    file, err := os.Open("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day6/data")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        for i,v := range line {
            if maps[i] == nil {
                maps[i] = make(map[string]int)
            }
            maps[i][string(v)]++
        }
    }

    var mostFrequent [8]string
    for i,v := range maps {
        prev := -1
        for k,value := range v {
            if prev < value {
                prev = value
                mostFrequent[i] = k
            }
        }
    }
    fmt.Printf("Most Frequent: %s\n", strings.Join(mostFrequent[0:], ""))

    var leastFrequent [8]string
    for i,v := range maps {
        prev := 999
        for k,value := range v {
            if prev > value {
                prev = value
                leastFrequent[i] = k
            }
        }
    }

    fmt.Printf("Least Frequent: %s\n", strings.Join(leastFrequent[0:], ""))
}
