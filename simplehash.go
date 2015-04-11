package gash

type SimpleHash struct {
    items []KvPair
    capacity int
}

func Create(capacity int) SimpleHash {
    table := SimpleHash{}
    table.capacity = capacity
    table.items = make([]KvPair, capacity)
    return table
}

func (table SimpleHash) Insert(k string, v interface{}) {
    index := Djb2(k) % table.capacity
    pair := KvPair{k, v}
    table.items[index] = pair
}

func (table SimpleHash) Find(k string) interface{} {
    index := Djb2(k) % table.capacity
    pair := table.items[index]
    return pair.Value
}
