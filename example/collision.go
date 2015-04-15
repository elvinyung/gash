package main

import (
    "fmt"
    "github.com/elvinyung/gash"
)

func main() {
    p := gash.KvPair{"samplekey", "samplevalue"}
    fmt.Println(p.Key)

    t := gash.CreateSimpleHash(1000, gash.Djb2)
    t.Insert("foo", 4)
    fmt.Println(t.Find("foo"))
    tg.Insert("foo", "asdfofsdja")
    fmt.Println(t.Find("foo"))
}
