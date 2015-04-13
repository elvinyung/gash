package main

import (
    "fmt"
    "github.com/elvinyung/gash"
)

func main() {
    p := gash.KvPair{"samplekey", "samplevalue"}
    fmt.Println(p.Key)

    t := gash.Create(1000)
    t.Insert("foo", 4)
    fmt.Println(t.Find("foo"))
    t.Remove("foo")
    fmt.Println(t.Find("foo"))
}
