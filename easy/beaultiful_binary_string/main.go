package main

import (
    "fmt"
)

func beautifulBinaryString(b string) int {
    count := 0
    runes := []rune(b)
    for i := 1; i < len(runes)-1; i++ {
        if runes[i] == '1' && runes[i-1] == '0' && runes[i+1] == '0' {
            runes[i+1] = '1'
            count++
        }
    }
    return count
}

func main() {
    tests := []struct {
        input    string
        expected int
    }{
        {"0101010", 2},
        {"01100", 0},
        {"0100101010", 3},
    }

    for _, test := range tests {
        result := beautifulBinaryString(test.input)
        if result != test.expected {
            fmt.Printf("For input '%s', expected %d but got %d\n", test.input, test.expected, result)
        } else {
            fmt.Println("PASS")
        }
    }
}