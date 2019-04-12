package iota

import (
	"fmt"
    "os"
    "time"
    "math/rand"
)

func Unique(a []int) {
    var value int

    for i:= range a {
        value = value ^ a[i]

    }
    fmt.Println(value)

}

func Random() int {
    rand.Seed(time.Now().UnixNano())
    min := 1
    max := 73
    return (rand.Intn(max - min) + min)
}
   

func TestMe(a string, b string, c string) {
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]
    fmt.Println(argsWithoutProg) // [a1 b1 c1 d1]
    fmt.Println(arg) // c1

    fmt.Println(333, a, b, c) // c1
}