package main

type OrderedMap struct {
	store map[int]interface{}
	keys  []int
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		store: map[int]interface{}{},
		keys:  []int{},
	}
}

func (o *OrderedMap) Get(key int) (interface{}, bool) {
	val, exists := o.store[key]
	return val, exists
}

func (o *OrderedMap) Set(key int, val interface{}) {
	if _, exists := o.store[key]; !exists {
		o.keys = append(o.keys, key)
	}
	o.store[key] = val
}

func (o *OrderedMap) Delete(key int) {
	delete(o.store, key)

	// Find key in slice
	var idx *int

	for i, val := range o.keys {
		if val == key {
			idx = &[]int{i}[0]
			break
		}
	}
	if idx != nil {
		o.keys = append(o.keys[:*idx], o.keys[*idx+1:]...)
	}
}

func (o *OrderedMap) Iterator() func() (*int, *int, interface{}) {
	var keys = o.keys

	j := 0

	return func() (_ *int, _ *int, _ interface{}) {
		if j > len(keys)-1 {
			return
		}

		row := keys[j]
		j++

		return &[]int{j - 1}[0], &row, o.store[row]
	}
}
