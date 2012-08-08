package parser

import (
    "regexp"
    "log"
    "strings"
    "strconv"
)

type ParsedExpr struct {
    cmd string
    number []int64 //want to change this to int
    num_range string
    re_expr string
}

func Parse(expr string) ParsedExpr {

    pe := ParsedExpr{cmd: ""} // this way we can test for a len of zero

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
                pe.re_expr = tmp
                if expr != "" { expr = strings.SplitAfter(expr, re_expr.FindString(expr))[1] }

            case num_range.MatchString(expr):
                tmp := strings.SplitAfter(expr, num_range.FindString(expr))[0]
                pe.num_range = tmp
                if expr != "" { expr = strings.SplitAfter(expr, num_range.FindString(expr))[1] }

            case number.MatchString(expr):
                //TODO convert numbers to ints when they are split
                tmp, _ := strconv.ParseInt(strings.SplitAfter(expr, number.FindString(expr))[0], 10, 0)
                pe.number = append(pe.number, tmp)
                if expr != "" { expr = strings.SplitAfter(expr, number.FindString(expr))[1] }

            case cmd.MatchString(expr):
                tmp := strings.SplitAfter(expr, cmd.FindString(expr))[0]
                pe.cmd = tmp
                if expr != "" { expr = strings.SplitAfter(expr, cmd.FindString(expr))[1] }

            default: log.Fatal("Invalid") // TODO make it exit better

        }
    }
    return pe
}
