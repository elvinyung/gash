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
    for _, pair := range table.items[index] {
        if pair.Key == k {
            return pair
        }
    }
    return nil
}

func (table SimpleHash) Remove(k string) {
    index := Djb2(k) % table.capacity
    for searchIndex, pair := range table.items[index] {
        if pair.Key == k {
            table.items[index] = append(table.items[index][:searchIndex], 
                                        table.items[index][searchIndex+1:]...)
            break
        }
    }
}
