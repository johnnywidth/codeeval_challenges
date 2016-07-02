package main

import (
    "fmt"
    "math/big"
    "strconv"
    "unicode/utf8"
)

var N = 1000

func main() {
    var max int64
    var left string
    var right string

    for i := 0; i <= N; i++ {
        a := strconv.Itoa(i)
        l := len(a)

        switch {
        case l == 1:
            continue
        case l%2 == 0:
            left = a[:(l / 2)]
            right = a[(l / 2):]
        case l%2 == 1:
            left = a[:(l / 2)]
            right = a[(l/2 + 1):]
        }

        reverse := func(s string) string {
            size := len(s)
            buf := make([]byte, size)
            for start := 0; start < size; {
                r, n := utf8.DecodeRuneInString(s[start:])
                start += n
                utf8.EncodeRune(buf[size-start:], r)
            }
            return string(buf)
        }

        if left == reverse(right) {
            ni := big.NewInt(int64(i))
            if ni.ProbablyPrime(1) {
                m := ni.Int64()
                if max < m {
                    max = m
                }
            }
        }
    }

    if max > 0 {
        fmt.Print(max)
    }
}
