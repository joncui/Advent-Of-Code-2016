package main
import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strconv"
    "strings"
)

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

    file, err := os.Open("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day3/data")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    count := 0
    if part == "1" {
        count = countHorizontalTriangles(scanner)
    } else {
        count = countVerticalTriangles(scanner)
    }

    fmt.Println(count)
}

func countHorizontalTriangles(scanner *bufio.Scanner) (int) {
    count := 0
    for scanner.Scan() {
        sides, err := sliceAtoi(regexp.MustCompile(" +").Split(strings.Trim(scanner.Text(), " "), 3))
        check(err)

        if isTriangle(sides) {
            count++
        }
    }

    return count
}

func countVerticalTriangles(scanner *bufio.Scanner) (int) {
    count := 0
    for scanner.Scan() {
        row1, err := sliceAtoi(regexp.MustCompile(" +").Split(strings.Trim(scanner.Text(), " "), 3))
        check(err)

        var row2 []int
        if scanner.Scan() {
            row2, err = sliceAtoi(regexp.MustCompile(" +").Split(strings.Trim(scanner.Text(), " "), 3))
            check(err)
        } else {
            return count
        }

        var row3 []int
        if scanner.Scan() {
            row3, err = sliceAtoi(regexp.MustCompile(" +").Split(strings.Trim(scanner.Text(), " "), 3))
            check(err)
        } else {
            return count
        }

        for i := 0; i < 3; i++ {
            sides := []int{row1[i], row2[i], row3[i]}
            if isTriangle(sides) {
                count++
            }
        }
    }

    return count
}

func isTriangle(sides []int) (bool) {
    sort.Ints(sides)
    return sides[0] + sides[1] > sides[2]
}

func sliceAtoi(sa []string) ([]int, error) {
    si := make([]int, 0, len(sa))
    for _,a := range sa {
        i, err := strconv.Atoi(a)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }

    return si, nil
}
