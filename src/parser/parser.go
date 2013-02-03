package parser

import (
    "regexp"
    "log"
    "strings"
    "strconv"
)

type ParsedExpr struct {
    Cmd string
    Number []int64 //want to change this to int
    Num_range string
    Re_expr string
}

func Parse(expr string) ParsedExpr {

    pe := ParsedExpr{Cmd: ""} // this way we can test for a len of zero

    cmd, err := regexp.Compile("^[a-zA-Z]")
    number, err := regexp.Compile("^([1-9]+)")
    num_range, err := regexp.Compile("[1-9]+,[1-9]+")
    re_expr, err := regexp.Compile("[a-z]/[a-zA-Z]+/")


    if err != nil {
        log.Fatal(err)
    }

    for expr != "" {
        switch {
            case re_expr.MatchString(expr):
                tmp := strings.SplitAfter(expr, re_expr.FindString(expr))[0]
                pe.Re_expr = tmp
                if expr != "" { expr = strings.SplitAfter(expr, re_expr.FindString(expr))[1] }

            case num_range.MatchString(expr):
                tmp := strings.SplitAfter(expr, num_range.FindString(expr))[0]
                pe.Num_range = tmp
                if expr != "" { expr = strings.SplitAfter(expr, num_range.FindString(expr))[1] }

            case number.MatchString(expr):
                //TODO convert numbers to ints when they are split
                tmp, _ := strconv.ParseInt(strings.SplitAfter(expr, number.FindString(expr))[0], 10, 0)
                pe.Number = append(pe.Number, tmp)
                if expr != "" { expr = strings.SplitAfter(expr, number.FindString(expr))[1] }

            case cmd.MatchString(expr):
                tmp := strings.SplitAfter(expr, cmd.FindString(expr))[0]
                pe.Cmd = tmp
                if expr != "" { expr = strings.SplitAfter(expr, cmd.FindString(expr))[1] }

            default: log.Fatal("Invalid") // TODO make it exit better

        }
    }
    return pe
}
