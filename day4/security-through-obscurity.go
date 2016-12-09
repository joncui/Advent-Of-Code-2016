package main

import (
    "bufio"
    "bytes"
    "fmt"
    "os"
    "regexp"
    "sort"
    "strconv"
    "strings"
)

var encryptedNameMap map[string]int

type ByFrequency []string

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, err := os.Open("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day4/data")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    re := regexp.MustCompile(`([a-z\-]+)+\-([0-9]+)\[([a-z]+)\]`)
    sectorIds := 0

    for scanner.Scan() {
        room := scanner.Text()
        roomArr := re.FindAllStringSubmatch(room, -1)[0][1:]
        encryptedNameMap = make(map[string]int)

        encryptedName := roomArr[0]
        sectorId, _ := strconv.Atoi(roomArr[1])
        checksum := roomArr[2]

        // Store frequency in map
        for _,v := range encryptedName {
            if string(v) != "-" {
                encryptedNameMap[string(v)]++
            }
        }

        // Get keys into array
        keys := make([]string, 0, len(encryptedNameMap))
        for key,_ := range encryptedNameMap {
            keys = append(keys, key)
        }

        // Sort the keys by the frequency of the characters
        sort.Sort(ByFrequency(keys))

        // If the checksum is equal to the first 5 elements of the keys
        if strings.Join(keys[0:5], "") == checksum {
            sectorIds += sectorId
            decrypted := rotate(encryptedName, sectorId)
            if (strings.Contains(decrypted, "north")) && (strings.Contains(decrypted, "object")) {
                fmt.Printf("Sector %d is %q\n", sectorId, decrypted)
            }
        }

    }

    fmt.Println("Sector ID Total:", sectorIds)
}

func (s ByFrequency) Less(a, b int) bool {
    if encryptedNameMap[s[a]] == encryptedNameMap[s[b]] {
        return s[a] < s[b]
    }
    return encryptedNameMap[s[a]] > encryptedNameMap[s[b]]
}

func (s ByFrequency) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByFrequency) Len() int {
    return len(s)
}

func rotate(s string, num int) string {
    rotate := num % 26
    var buffer bytes.Buffer

    for _,v := range s {
        if v == 45 {
            buffer.WriteString(" ")
        } else {
            newV := v + rune(rotate)
            if newV > 122 {
                newV -= 26
            }

            buffer.WriteString(string(newV))
        }
    }

    return buffer.String()
}
