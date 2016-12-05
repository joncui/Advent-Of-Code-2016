package main

import (
    "bytes"
    "crypto/md5"
    "encoding/hex"
    "fmt"
    "os"
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

    input := "abc"
    if part == "1" {
        fmt.Println(part1(input))
    } else {
        fmt.Println(part2(input))
    }
}

func part1(input string) string {
    index := 0
    var passwordBuffer bytes.Buffer

    for passwordBuffer.Len() < 8 {
        hashString := getHashString(input, index)
        if strings.Index(hashString, "00000") == 0 {
            passwordBuffer.WriteString(string(hashString[5]))
        }
        index++
    }

    return passwordBuffer.String()
}

func part2(input string) string {
    index := 0
    count := 0
    var passwordArr [8]string

    for count < 8 {
        hashString := getHashString(input, index)
        if strings.Index(hashString, "00000") == 0 {
            if position, err := strconv.Atoi(string(hashString[5])); (err == nil) && (position < 8) {
                fmt.Printf("Hash: %s\n", hashString)
                fmt.Printf("Position: %d\n", position)
                if passwordArr[position] == "" {
                    passwordArr[position] = string(hashString[6])
                    count++
                    fmt.Printf("Password at position: %s\n", passwordArr[position])
                    fmt.Println(passwordArr)
                }
            }
        }
        index++
    }

    return strings.Join(passwordArr[0:], "")
}

func getHashString(input string, index int) string {
    hash := md5.Sum([]byte(input + strconv.Itoa(index)))
    hashString := hex.EncodeToString(hash[:])

    return hashString
}
