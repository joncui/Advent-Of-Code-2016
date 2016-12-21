package main

import (
    "container/list"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "sort"
)

type Coord struct {
    x, y int
    path string
}

var directions = []string{"U", "D", "L", "R"}

func main() {
    input := "dmypynyp"
    paths := search(input)
    sort.Sort(sort.Reverse(sort.IntSlice(paths)))
    fmt.Printf("Longest Path is of length %d\n", paths[0])
}

func search(input string) (pathSizes []int) {
    queue := list.New()
    queue.PushBack(Coord{0, 0, ""})

    for q := queue.Front(); q != nil; q = q.Next() {
        val := q.Value.(Coord)
        if (val.x == 3) && (val.y == 3) {
            if len(pathSizes) == 0 {
                fmt.Printf("Shortest Path is %v\n", val.path)
            }
            pathSizes = append(pathSizes, len(val.path))
            continue
        }
        queue.PushBackList(getAdjacent(input, val))
    }

    return
}

func getAdjacent(input string, coord Coord) *list.List {
    adjacent := list.New()
    hash := getHashString(input + coord.path)

    for i,c := range hash {
        if (c >= 98) && (c <= 102) {
            newCoord, valid := getNewCoord(i, coord)
            if valid {
                adjacent.PushBack(newCoord)
            }
        }
    }

    return adjacent
}

func getHashString(str string) string {
    hash := md5.Sum([]byte(str))
    hashString := hex.EncodeToString(hash[:])

    return hashString[0:4]
}

func getNewCoord(index int, currentCoord Coord) (Coord, bool) {
    var c Coord
    path := currentCoord.path + directions[index]

    x := currentCoord.x
    y := currentCoord.y
    switch directions[index] {
        case "U":
            y--
        case "D":
            y++
        case "L":
            x--
        case "R":
            x++
    }

    if (x < 0) || (x > 3) || (y < 0) || (y > 3) {
        return c, false
    }

    c = Coord{x, y, path}

    return c, true
}
