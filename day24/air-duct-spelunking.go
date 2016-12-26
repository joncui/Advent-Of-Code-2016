package main

import (
    "container/list"
    "fmt"
    "github.com/fighterlyt/permutation"
    "io/ioutil"
    "regexp"
    "sort"
    "strings"
)

type Coord struct {
    x, y, distance int
}

var indices = make([]Coord, 9)
var blueprint [][]string
var locations = []int{0, 1, 2, 3, 4 , 5, 6, 7}
var adjTree = make([][]int, 9)
var width int
var height int

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day24/data")
    check(err)

    parseData(strings.Split(strings.TrimSpace(string(data)), "\n"))
    width = len(blueprint[0])
    height = len(blueprint)

    sort.Ints(locations)
    buildAdjacencyTree()
    steps, path := findSteps(false)
    fmt.Printf("Shortest distance is %d using %v\n", steps, path)

    steps, path = findSteps(true)
    fmt.Printf("Shortest distance is %d using %v for Part 2\n", steps, path)
}

func parseData(lines []string) {
    re := regexp.MustCompile(`(\d)`)

    for y,line := range lines {
        if strings.Contains(line, ".") {
            r := re.FindAllStringIndex(line, -1)
            if len(r) != 0 {
                for _,index := range r {
                    x := index[0]
                    indices[line[x] - '0'] = Coord{x - 1, y - 1, 0}
                }
            }
            blueprint = append(blueprint, strings.Split(line, "")[1:len(line) - 1])
        }
    }
}

func buildAdjacencyTree() {
    for i := 0; i < len(locations); i++ {
        adjTree[i] = make([]int, 9)
    }

    for i := 0; i < len(locations) - 1; i++ {
        for j := i + 1; j < len(locations); j++ {
            startIndex := locations[i]
            endIndex := locations[j]
            start := indices[startIndex]
            end := indices[endIndex]
            dist := search(start, end)
            adjTree[startIndex][endIndex], adjTree[endIndex][startIndex] = dist, dist
        }
    }
}

func search(start, end Coord) int {
    seen := make(map[Coord]bool)
    queue := list.New()

    queue.PushBack(start)

    for q := queue.Front(); q != nil; q = q.Next() {
        val := q.Value.(Coord)
        if seen[val] {
            continue
        } else if (val.x == end.x) && (val.y == end.y) {
            return val.distance
        }
        seen[val] = true
        queue.PushBackList(getAdjacent(val))
    }

    return 0
}

func getAdjacent(current Coord) *list.List {
    l := list.New()
    for x := current.x - 1; x < current.x + 2; x++ {
        if (x == current.x) || (x < 0) || (x >= width) || !isOpen(x, current.y) {
            continue
        }

        l.PushBack(Coord{x, current.y, current.distance + 1})

    }

    for y := current.y - 1; y < current.y + 2; y++ {
        if (y == current.y) || (y < 0) || (y >= height) || !isOpen(current.x, y) {
            continue
        }

        l.PushBack(Coord{current.x, y, current.distance + 1})
    }
    return l
}

func isOpen(x, y int) bool {
    return blueprint[y][x] != "#"
}

func findSteps(part2 bool) (minSteps int, path []int) {
    i := []int{1, 2, 3, 4, 5, 6, 7}
    minSteps = 1000
    p, err := permutation.NewPerm(i, nil)
    check(err)

    for i,e := p.Next(); e == nil; i,e = p.Next() {
        steps := getSteps(i.([]int), part2)
        if steps < minSteps {
            minSteps = steps
            path = i.([]int)
        }
    }

    return
}

func getSteps(perm []int, part2 bool) (steps int) {
    steps = adjTree[0][perm[0]]
    for i := 0; i < len(perm) - 1; i++ {
        steps += adjTree[perm[i]][perm[i + 1]]
    }

    if part2 {
        steps += adjTree[0][perm[len(perm) - 1]]
    }

    return
}
