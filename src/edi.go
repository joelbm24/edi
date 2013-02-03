package main

import "fmt"
import "./parser"

var input string
var prompt string
var buffer []string

func main() {
    fmt.Print(prompt)
    fmt.Scan(&input)
    expr := parser.Parse(input)
    fmt.Println(expr.Cmd)
}
