package gash

type SimpleHash struct {
    items [][]KvPair
    capacity int
}

func Create(capacity int) SimpleHash {
    table := SimpleHash{}
    table.capacity = capacity
    table.items = make([][]KvPair, capacity)
    for index := range table.items {
        table.items[index] = []KvPair{}
    }

    return table
}

func (table SimpleHash) Insert(k string, v interface{}) {
    index := Djb2(k) % table.capacity
    item := KvPair{k, v}

    isSet := false
    for searchIndex, pair := range table.items[index] {
        if pair.Key == k {
            table.items[index][searchIndex].Value = v
            isSet = true
            break
        }
    }
    if (!isSet) {
        table.items[index] = append(table.items[index], item)
    }
}

func (table SimpleHash) Find(k string) interface{} {
    index := Djb2(k) % table.capacity
    var result KvPair
    for _, pair := range table.items[index] {
        if pair.Key == k {
            result = pair
            break
        }
    }
    return result.Value
}
