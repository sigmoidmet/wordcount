package main

import (
    "fmt"
    "strings"
    "os"
)

func main() {
    fmt.Println(len(strings.Fields(os.Args[1])))
}
