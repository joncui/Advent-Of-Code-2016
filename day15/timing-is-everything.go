package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
    "strings"
)

type State struct {
    num, max, pos int
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    part := "1"
    if len(os.Args) > 1 {
        part = os.Args[1]
    }

    states := parseData()
    var time int

    if part == "1" {
        time = find(states)
    } else {
        states = append(states, State{len(states) + 1, 11, 0})
        time = find(states)
    }

    fmt.Printf("Press the button at time %d to get the capsule\n", time)
}

func parseData() (states []State) {
    data, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day15/data")
    check(err)

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    for _, v := range lines {
        line := strings.Split(v, " ")
        disc,_ := strconv.Atoi(strings.Replace(line[1], "#", "", 1))
        max,_ := strconv.Atoi(line[3])
        pos,_ := strconv.Atoi(strings.Replace(line[11], ".", "", 1))

        states = append(states, State{disc, max, pos})
    }

    return
}

func find(states []State) (time int) {
    for {
        var nextPos []int
        for i,disc := range states {
            nextPos = append(nextPos, getNextPosition(disc))

            // Move disc to next start position
            disc.pos++
            states[i] = disc
        }

        if checkAllMatch(nextPos) {
            break
        }
        time++
    }


    return
}

func getNextPosition(discState State) (position int) {
    position = (discState.pos + discState.num) % discState.max
    return
}

func checkAllMatch(arr []int) bool {
    for i := 0; i < len(arr) - 1; i++ {
        if arr[i] != arr[i+1] {
            return false
        }
    }

    return true
}
