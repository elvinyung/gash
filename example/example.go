package main

import (
    "fmt"
    "github.com/elvinyung/gash"
)

func main() {
    p := gash.KvPair{"samplekey", "samplevalue"}
    fmt.Println(p.Key)

    t := gash.CreateSimpleHash(1000)
    t.Insert("foo", 4)
    t.Insert("bar", "asdfofsdja")
    t.Insert("baz", []int{1,2,3})
    fmt.Println(t.Find("foo"))
    fmt.Println(t.Find("bar"))
    fmt.Println(t.Find("baz"))
}
