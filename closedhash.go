package gash

// A linear-probed open addressed closed hash.

type ClosedHash struct {
    items []*KvPair
    capacity int
}

func CreateClosedHash(capacity int) ClosedHash {
    table := ClosedHash{}
    table.capacity = capacity
    table.items = make([]*KvPair, capacity)

    return table
}

func (table ClosedHash) Insert(k string, v interface{}) {
    hash := Djb2(k) % table.capacity
    index := hash

    // find first empty space
    for (table.items[index] != nil) {
        index++

        if (index == hash) {
            // we traversed the entire table without finding anything
            return
        }

        if (index >= table.capacity) {
            index = 0;
        }
    }

    item := KvPair{k, v}
    table.items[index] = &item
}

func (table ClosedHash) Find(k string) interface{} {
    hash := Djb2(k) % table.capacity
    index := hash

    for (table.items[index] != nil && table.items[index].Key != k) {
        index++

        if (index == hash) {
            // we traversed the entire table without finding anything
            return nil
        }

        if (index >= table.capacity) {
            index = 0;
        }
    }
    
    result := table.items[index]
    return result.Value
}

func (table ClosedHash) Remove(k string) {
    hash := Djb2(k) % table.capacity
    index := hash

    for (table.items[index] != nil && table.items[index].Key != k) {
        index++

        if (index == hash) {
            // we traversed the entire table without finding anything
            return
        }

        if (index >= table.capacity) {
            index = 0;
        }
    }

    table.items[index] = nil
}
