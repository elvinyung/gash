package gash

// a collection of various hash function implementations.

// a simple djb2 implementation
func Djb2(s string) int {
    hash := 5381
    for c := range s {
        hash += (hash * 33) + c
    }

    return hash
}
