package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, err := os.Open("/Users/jonathan.cui/go/src/github.com/joncui/Advent-Of-Code-2016/day7/data")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    re := regexp.MustCompile(`([a-z]+)`)
    var tlsCount int
    var sslCount int

    for scanner.Scan() {
        ip7 := scanner.Text()
        sections := re.FindAllStringSubmatch(ip7, -1)
        tlsCount += getTLSCount(sections)
        sslCount += getSSLCount(sections)
    }

    fmt.Printf("IPs that support TLS: %d\n", tlsCount)
    fmt.Printf("IPs that support SSL: %d\n", sslCount)
}

func getTLSCount(sections [][]string) int {
    bracketSection, noBracketSection := getSeparateSections(sections)

    hasBracketPalindrome, _ := hasNPalindromes(bracketSection, 4)
    hasNoBracketPalindrome, _ := hasNPalindromes(noBracketSection, 4)

    if hasNoBracketPalindrome && !hasBracketPalindrome {
        return 1
    }

    return 0
}

func getSSLCount(sections [][]string) int {
    bracketSection, noBracketSection := getSeparateSections(sections)

    _, bracketPalindrome := hasNPalindromes(bracketSection, 3)
    hasNoBracketPalindrome, noBracketPalindrome := hasNPalindromes(noBracketSection, 3)

    if hasNoBracketPalindrome {
        for _,v1 := range bracketPalindrome {
            for _,v2 := range noBracketPalindrome {
                if isReverse(v1, v2) {
                    return 1
                }
            }
        }
    }

    return 0
}

func getSeparateSections(sections [][]string) (string, string) {
    var bracketSection []string
    var noBracketSection []string

    for i,v := range sections {
        if i % 2 == 0 {
            noBracketSection = append(noBracketSection, v[0])
        } else {
            bracketSection = append(bracketSection, v[0])
        }
    }

    return strings.Join(bracketSection, " "), strings.Join(noBracketSection, " ")
}

func hasNPalindromes(s string, n int) (bool, []string) {
    hasPalindrome := false
    var palindromes []string
    for i := 0; i < len(s) - n + 1; i++ {
        if abba(s[i:i+n]) {
            hasPalindrome = true
            palindromes = append(palindromes, s[i:i+n])
        }
    }

    return hasPalindrome, palindromes
}

func abba(s string) bool {
    if len(s) == 4 {
        return (s[0] != s[1]) && (s[0] == s[3]) && (s[1] == s[2])
    } else if len(s) == 3 {
        return (s[0] != s[1]) && (s[0] == s[2])
    }
    return false
}

func isReverse (s1, s2 string) bool {
    if (len(s1) == 3) && (len(s2) == 3) {
        return (s1[0] == s2[1]) && (s2[0] == s1[1])
    }

    return false
}
