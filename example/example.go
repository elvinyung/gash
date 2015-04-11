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
    t.Insert("quux", "asdfofsdja")
    fmt.Println(t.Find("foo"))
    fmt.Println(t.Find("quux"))
}
