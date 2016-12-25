package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    "strconv"
    "strings"
)

type df struct {
    size, used, avail int
}

type coord struct {
    x, y int
}

var diskfree map[coord]df
var grid = make([][]string, 29)
var strGrid = make([][]string, 29)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    data, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day22/data")
    check(err)

    arr := strings.Split(strings.TrimSpace(string(data)), "\n")[2:]
    diskfree = make(map[coord]df)

    parseData(arr)
    fmt.Printf("%d viable pairs\n", findPairCount())

    printGrid(grid)
    printGrid(strGrid)

    // Empty node starts at x8-y28
    // moving empty node to x33-y0 takes 67 steps (right next to Goal Data)
    // moving Goal Data left takes 5 steps each. Takes 165 to get Goal Data to x1-y0
    // Plus 1 step to get Goal Data to x0-y0
    fmt.Printf("233 steps to get Goal Data x34-y0 to x0-y0")
}

func parseData(arr []string) {
    re := regexp.MustCompile(`x(\d+)-y(\d+) +(\d+)T +(\d+)T +(\d+)T`)

    for _,a := range arr {
        r := re.FindAllStringSubmatch(a, -1)[0][1:]

        x,_ := strconv.Atoi(r[0])
        y,_ := strconv.Atoi(r[1])
        size,_ := strconv.Atoi(r[2])
        used,_ := strconv.Atoi(r[3])
        avail,_ := strconv.Atoi(r[4])

        c := coord{x, y}
        diskfree[c] = df{size, used, avail}

        val := getVal(x, y, size, used, avail)
        str := getStr(size, used)

        if len(grid[y]) == 0 {
            grid[y] = make([]string, 35)
        }
        grid[y][x] = val

        if len(strGrid[y]) == 0 {
            strGrid[y] = make([]string, 35)
        }
        strGrid[y][x] = str
    }
}

func getVal(x, y, size, used, avail int) (val string) {
    if used == 0 {
        val = "_"
    } else if size < 100 {
        val = "."
    } else {
        val = "#"
    }

    if (x == 0) && (y == 0) {
        val = "(" + val + ")"
    } else if (x == 34) && (y == 0) {
        val = " G "
    } else {
        val = " " + val + " "
    }

    return
}

func getStr(size, used int) (str string) {
    str = strconv.Itoa(used) + "/" + strconv.Itoa(size)
    return
}

func findPairCount() (count int) {
    for k1,v1 := range diskfree {
        if v1.used == 0 {
            continue
        }

        for k2,v2 := range diskfree {
            if k1 == k2 {
                continue
            }

            if v1.used <= v2.avail {
                count++
            }
        }
    }

    return
}

func printGrid(g [][]string) {
    for _,y := range g {
        fmt.Println(y)
    }
}

// func (s ByUsed) Less(a, b int) bool {
//     if encryptedNameMap[s[a]] == encryptedNameMap[s[b]] {
//         return s[a] < s[b]
//     }
//     return encryptedNameMap[s[a]] > encryptedNameMap[s[b]]
// }
//
// func (s ByUsed) Swap(i, j int) {
//     s[i], s[j] = s[j], s[i]
// }
//
// func (s ByUsed) Len() int {
//     return len(s)
// }
//
