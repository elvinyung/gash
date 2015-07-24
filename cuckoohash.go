package gash

// A cuckoo hash.
// each key is associated with two different indices,
//  and they're switched around when resolving collisions.

type CuckooHash struct {
	items    []*KvPair
	capacity int
	fn1      HashFn
	fn2      HashFn
}

func CreateCuckooHash(capacity int, fn1 HashFn, fn2 HashFn) CuckooHash {
	table := CuckooHash{}
	table.capacity = capacity
	table.items = make([]*KvPair, capacity)
	table.fn1 = fn1
	table.fn2 = fn2
	return table
}

func (table CuckooHash) Insert(k string, v interface{}) {
	// insert original element
	index := table.fn1(k) % table.capacity
	swapItem := table.items[index]
	insertItem := KvPair{k, v}
	table.items[index] = &insertItem
	oldIndex := index

	// cascade displace other elements
	item := swapItem
	for item != nil {
		index = table.fn1(swapItem.Key) % table.capacity
		if index == oldIndex {
			index = table.fn2(swapItem.Key) % table.capacity
		}

		swapItem = table.items[index]
		if swapItem != nil && swapItem.Key == k {
			// somehow, we've reached a loop
			// TODO: implement remove and reinsert everything

			return
		}
		table.items[index] = item
		item = swapItem
	}
}

func (table CuckooHash) Find(k string) interface{} {
	var pair *KvPair

	index1 := table.fn1(k) % table.capacity
	index2 := table.fn2(k) % table.capacity

	pair1 := table.items[index1]
	pair2 := table.items[index2]

	if pair1 != nil && pair1.Key == k {
		pair = pair1
	} else if pair2 != nil && pair2.Key == k {
		pair = pair2
	}

	return pair.Value
}

func (table CuckooHash) Remove(k string) {
	index1 := table.fn1(k) % table.capacity
	index2 := table.fn2(k) % table.capacity

	pair1 := table.items[index1]
	pair2 := table.items[index2]

	if pair1 != nil && pair1.Key == k {
		table.items[index1] = nil
	} else if pair2 != nil && pair2.Key == k {
		table.items[index2] = nil
	}
}
