package main

import (
    "fmt"
    "regexp"
    "log"
    "strings"
)

func main() {
    expr := "2333,33m5"

    cmd, err := regexp.Compile("[a-zA-Z]")
    number, err := regexp.Compile("[1-9]*")
    num_range, err := regexp.Compile("[1-9]+,[1-9]+")
    re_expr, err := regexp.Compile("[a-z]/[a-zA-Z]+/")

    var parsed_expr []string

    if err != nil {
        log.Fatal(err)
    }

    for expr != "" {
        switch {
            case re_expr.MatchString(expr):
                tmp := strings.SplitAfter(expr, re_expr.FindString(expr))[0]
                parsed_expr = append(parsed_expr, tmp)
                if expr != "" { expr = strings.SplitAfter(expr, re_expr.FindString(expr))[1] }

            case num_range.MatchString(expr):
                tmp := strings.SplitAfter(expr, num_range.FindString(expr))[0]
                parsed_expr = append(parsed_expr, tmp)
                if expr != "" { expr = strings.SplitAfter(expr, num_range.FindString(expr))[1] }

            case number.MatchString(expr):
                tmp := strings.SplitAfter(expr, number.FindString(expr))[0]
                parsed_expr = append(parsed_expr, tmp)
                if expr != "" { expr = strings.SplitAfter(expr, number.FindString(expr))[1] }

            case cmd.MatchString(expr):
                tmp := strings.SplitAfter(expr, cmd.FindString(expr))[0]
                parsed_expr = append(parsed_expr, tmp)
                if expr != "" { expr = strings.SplitAfter(expr, cmd.FindString(expr))[1] }

            default: log.Fatal("Invalid") // TODO make it exit better
        }
    }
    fmt.Println(parsed_expr)
}
