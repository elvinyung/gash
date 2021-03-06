package gash

// A simple chained hash.

type SimpleHash struct {
	items    [][]*KvPair
	capacity int
	fn       HashFn
}

func CreateSimpleHash(capacity int, fn HashFn) SimpleHash {
	table := SimpleHash{}
	table.capacity = capacity
	table.items = make([][]*KvPair, capacity)
	for index := range table.items {
		table.items[index] = []*KvPair{}
	}
	table.fn = fn

	return table
}

func (table SimpleHash) Insert(k string, v interface{}) {
	index := table.fn(k) % table.capacity
	item := KvPair{k, v}

	isSet := false
	for searchIndex, pair := range table.items[index] {
		if pair.Key == k {
			table.items[index][searchIndex].Value = v
			isSet = true
			break
		}
	}
	if !isSet {
		table.items[index] = append(table.items[index], &item)
	}
}

func (table SimpleHash) Find(k string) interface{} {
	index := table.fn(k) % table.capacity
	for _, pair := range table.items[index] {
		if pair.Key == k {
			return pair.Value
		}
	}
	return nil
}

func (table SimpleHash) Remove(k string) {
	index := table.fn(k) % table.capacity
	for searchIndex, pair := range table.items[index] {
		if pair.Key == k {
			table.items[index] = append(table.items[index][:searchIndex],
				table.items[index][searchIndex+1:]...)
			break
		}
	}
}
