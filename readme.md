# Gash
### by [Elvin Yung](https://github.com/elvinyung)

Gash is a simple Go hashmap library, containing a few implementations of various types of hashmaps. 

### Quickstart
```go
package main

import (
    "fmt"
    "github.com/elvinyung/gash"
)

func main() {
    t := gash.CreateSimpleHash(1000, gash.Djb2)
    t.Insert("foo", 4)
    t.Insert("bar", "asdfofsdja")
    t.Insert("baz", []int{1,2,3})
    fmt.Println(t.Find("foo"))
    fmt.Println(t.Find("bar"))
    fmt.Println(t.Find("baz"))
}
```

### Other stuff

Checklist:
* [x] chain (simple) hash
* [x] closed hash
* [x] cuckoo hash
* [ ] double hash
* [ ] coalesced hash
* [ ] some sort of hash set
* [ ] some sort of bloom filter
