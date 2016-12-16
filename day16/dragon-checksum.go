package main

import (
    "fmt"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    input := "10010000000110000"
    fillSize := 272

    fillData := getFillData(input, fillSize)
    checkSum := getChecksum(fillData)
    fmt.Printf("The checksum for input %q is %q\n", input, checkSum)

    fillSizePt2 := 35651584
    fillDataPt2 := getFillData(input, fillSizePt2)
    checkSumPt2 := getChecksum(fillDataPt2)
    fmt.Printf("The checksum for input %q is %q\n", input, checkSumPt2)
}

func getFillData(currentData string, desiredLength int) string {
    if len(currentData) >= desiredLength {
        return currentData[0:desiredLength]
    }
    a := currentData
    b := swapString(reverseString(a), '0', '1')

    return getFillData(a + "0" + b, desiredLength)
}

func reverseString(str string) string {
    runes := []rune(str)
    for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

func swapString(str string, swapA, swapB rune) string {
    runes := []rune(str)
    for i := 0; i < len(runes); i++ {
        if runes[i] == swapA {
            runes[i] = swapB
        } else if runes[i] == swapB {
            runes[i] = swapA
        }
    }

    return string(runes)
}

func getChecksum(currentCheckSum string) string {
    if len(currentCheckSum) % 2 == 1 {
        return currentCheckSum
    }

    var newCheckSum []rune
    for i := 0; i < len(currentCheckSum) - 1; i += 2 {
        if currentCheckSum[i] == currentCheckSum[i+1] {
            newCheckSum = append(newCheckSum, '1')
        } else {
            newCheckSum = append(newCheckSum, '0')
        }
    }

    return getChecksum(string(newCheckSum))
}
