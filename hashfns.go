package gash

// a collection of various hash function implementations.

// a simple djb2 implementation
func Djb2(s string) int {
    hash := 5381
    for _, c := range s {
        hash += (hash * 33) + int(c)
    }

    return hash
}

// sdbm
func Sdbm(s string) int {
    hash := 0
    for _, c := range s {
        hash += (hash * 65599) + int(c)
    }

    return hash
}
