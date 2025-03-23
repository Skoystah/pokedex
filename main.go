package main

import "strings"

func cleanInput(text string) []string {
    lower := strings.ToLower(text)
    res := strings.Fields(lower)
    return res
}

func main() {
    print("Hello, World!")
}
