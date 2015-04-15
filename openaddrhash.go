package gash
/*

// A linear-probed open addressed closed hash.

type OpenAddrHash struct {
    items []KvPair
    capacity int
}

func CreateSimpleHash(capacity int) OpenAddrHash {
    table := OpenAddrHash{}
    table.capacity = capacity
    table.items = make([]KvPair, capacity)

    return table
}

func (table OpenAddrHash) Insert(k string, v interface{}) {
    hash := Djb2(k) % table.capacity
    item := KvPair{k, v}

    
    hash++
}

func (table OpenAddrHash) Find(k string) interface{} {
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

func (table OpenAddrHash) Remove(k string) {
    index := Djb2(k) % table.capacity
    for searchIndex, pair := range table.items[index] {
        if pair.Key == k {
            table.items[index] = append(table.items[index][:searchIndex], 
                                        table.items[index][searchIndex+1:]...)
            break
        }
    }
}
*/