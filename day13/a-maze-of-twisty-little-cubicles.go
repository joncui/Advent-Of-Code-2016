package main

import (
    "container/list"
    "fmt"
    "strconv"
    "strings"
)

type Coords struct {
    x, y int
}

var seen map[Coords]int

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    steps := search(Coords{1, 1}, 1364, 31, 39)
    fmt.Printf("It takes %d steps to reach 31, 39\n", steps) // Part 1
    fmt.Printf("There are %d coordinates that are at most 50 steps from 1, 1\n", getSeenWithNSteps(51)) // Part 2

    // Test
    steps = search(Coords{1, 1}, 10, 7, 4)
    fmt.Printf("It takes %d steps to reach %d, %d\n", steps, 7, 4)
}

func search(start Coords, fav, x, y int) int {
    seen = make(map[Coords]int)
    queue := list.New()

    queue.PushBack(start)
    seen[start] = 1

    for q := queue.Front(); q != nil; q = q.Next() {
        val := q.Value.(Coords)
        adjacent := getAdjacent(val)
        for e := adjacent.Front(); e != nil; e = e.Next() {
            space := e.Value.(Coords)
            if space.x == x && space.y == y {
                return seen[val]
            }

            if (!isOpenSpace(space, fav)) || (seen[space] > 0) {
                continue
            } else {
                seen[space] = seen[val] + 1
                queue.PushBack(space)
            }
        }

    }

    return 1
}

func getAdjacent(current Coords) *list.List {
    l := list.New()
    for x := current.x - 1; x < current.x + 2; x++ {
        if x < 0 {
            continue
        }

        l.PushBack(Coords{x, current.y})

    }

    for y := current.y - 1; y < current.y + 2; y++ {
        if y < 0 {
            continue
        }

        l.PushBack(Coords{current.x, y})
    }
    return l
}

func getSeenWithNSteps(n int) (count int) {
    for _, v := range seen {
        if v <= n {
            count++
        }
    }

    return
}

func isOpenSpace(coord Coords, fav int) bool {
    x, y := coord.x, coord.y
    formula := int64((x * x) + (3 * x) + (2 * x * y) + y + (y * y) + fav)
    numberOfOnes := len(strings.Replace(strconv.FormatInt(formula, 2), "0", "", -1))

    return (numberOfOnes % 2) == 0
}
