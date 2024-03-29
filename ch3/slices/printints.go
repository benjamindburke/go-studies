package main

import (
    "bytes"
    "fmt"
    "strconv"
)

func intsToString(values []int) string {
    var buf bytes.Buffer
    buf.WriteByte('[')
    for i, v := range values {
        if i > 0 {
            buf.WriteString(", ")
        }
        fmt.Fprintf(&buf, "%d", v)
    }
    buf.WriteByte(']')
    return buf.String()
}

func main() {
    fmt.Println(intsToString([]int{1, 2, 3}))
    a := 123
    b := fmt.Sprintf("%d", a)
    fmt.Println(b, strconv.Itoa(a)) // "123 123"
    fmt.Println(strconv.FormatInt(int64(a), 2)) // "1111011"
    s := fmt.Sprintf("a=%b", a) // "a=1111011"
    fmt.Println(s)
    x, _ := strconv.Atoi("123") // x is an int
    fmt.Println(x)
    y, _ := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
    fmt.Println(y)
}

