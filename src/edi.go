package main

import "fmt"
import "./parser"

var input string
var prompt string


func main() {
    fmt.Print(prompt)
    fmt.Scan(&input)
    fmt.Println(parser.Parse(input))
}
