package main
import (
    "fmt"
    "io/ioutil"
    "math"
    "os"
    "strings"
    "strconv"
)

var nsew = [4]string{"N", "E", "S", "W"}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// x, y: the coordinates
// distance: The given number of blocks to go one direction
// direaction: The given direction (Left or RIght)
// orientation: North/East/South/West
func main() {
    var part string
    if part = "1"; len(os.Args) > 1 {
        part = os.Args[1]
    }

    blockMap := make(map[string]bool)

    sequences, err := ioutil.ReadFile("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day1/data")
    check(err)
    var arr []string = strings.Split(string(sequences), ", ")
    x, y, orientation := 0, 0, 0
    done := false

    for _,v := range arr {
        value := strings.TrimSpace(v)
        direction := value[0:1]
        distance, err := strconv.Atoi(value[1:len(value)])
        check(err)
        if part == "1" {
            x, y, orientation = nextLocation(x, y, orientation, int(distance), direction)
        } else {
            x, y, orientation, done = nextLocationPt2(x, y, orientation, int(distance), direction, blockMap)
            if done {
                break
            }
        }
        fmt.Printf("%q -> %d, %d\n", nsew[orientation], x, y)
    }

    fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func nextLocation(x, y, orientation, distance int, direction string) (int, int, int) {
    orientation = getNextOrientation(orientation, direction)
    axis := getAxisToUpdate(&x, &y, orientation)
    multiplier := getMultiplier(orientation)

    *axis += (distance * multiplier)

    return x, y, orientation
}

func nextLocationPt2(x, y, orientation, distance int, direction string, blockMap map[string]bool) (int, int, int, bool) {
    orientation = getNextOrientation(orientation, direction)
    axis := getAxisToUpdate(&x, &y, orientation)
    multiplier := getMultiplier(orientation)

    for i := 0; i < distance; i++ {
        key := strconv.Itoa(x) + "," + strconv.Itoa(y)
        if blockMap[key] {
            return x, y, orientation, true
        } else {
            blockMap[key] = true
        }
        *axis += multiplier
    }
    return x, y, orientation, false
}

// Which direction will we be headed next?
func getNextOrientation(orientation int, direction string) (int) {
    if direction == "R" {
        orientation++
    } else {
        orientation--
    }

    // Add 4 to orientation so the results dont ever come back as negative
    // Even if orientation is initially negative
    orientation = (orientation + 4) % 4
    return orientation
}

func getAxisToUpdate(x *int, y *int, orientation int) (*int) {
    // Update North/South (y-axis) if orientation is 0 or 2
    // Update East/West (x-axis) if orientation is 1 or 3
    if orientation % 2 == 0 {
        return y
    } else {
        return x
    }
}

// Multiplier is used to determine if the distance is added or subtracted to the axis
func getMultiplier(orientation int) (int) {
    if orientation <= 1 {
        return 1
    } else {
        return -1
    }
}
