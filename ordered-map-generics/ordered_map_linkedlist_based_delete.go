package main

import "container/list"

type keyValueHolder struct {
	key   int
	value interface{}
}

type LinkedListBasedOrderedMap struct {
	store map[int]*list.Element
	keys  *list.List
}

func NewLinkedListBasedOrderedMap() *LinkedListBasedOrderedMap {
	return &LinkedListBasedOrderedMap{
		store: map[int]*list.Element{},
		keys:  list.New(),
	}
}

func (o *LinkedListBasedOrderedMap) Get(key int) (interface{}, bool) {
	val, exists := o.store[key]
	if !exists {
		return 0, false
	}
	return val.Value.(keyValueHolder).value, true
}

func (o *LinkedListBasedOrderedMap) Set(key int, val interface{}) {
	var e *list.Element
	if _, exists := o.store[key]; !exists {
		e = o.keys.PushBack(keyValueHolder{
			key:   key,
			value: val,
		})
	} else {
		e = o.store[key]
		e.Value = keyValueHolder{
			key:   key,
			value: val,
		}
	}
	o.store[key] = e
}

func (o *LinkedListBasedOrderedMap) Delete(key int) {
	e, exists := o.store[key]
	if !exists {
		return
	}

	o.keys.Remove(e)

	delete(o.store, key)
}

func (o *LinkedListBasedOrderedMap) Iterator() func() (*int, *int, interface{}) {
	e := o.keys.Front()
	j := 0
	return func() (_ *int, _ *int, _ interface{}) {
		if e == nil {
			return
		}

		keyVal := e.Value.(keyValueHolder)
		j++
		e = e.Next()

		return func() *int { v := j - 1; return &v }(), &keyVal.key, keyVal.value
	}
}
