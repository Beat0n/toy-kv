package instance

import "toy-kv/iface"

type KvMap[Key, Value any] struct {
	m map[any]Value
}

func (kvmap *KvMap[Key, Value]) Get(key Key) (Value, bool) {
	value, ok := kvmap.m[key]
	return value, ok
}

func (kvmap *KvMap[Key, Value]) Set(key Key, value Value) {
	kvmap.m[key] = value
}

func (kvmap *KvMap[Key, Value]) Delete(key Key) {
	delete(kvmap.m, key)
}

func (kvmap *KvMap[Key, Value]) Range(f func(key Key, value Value) bool) {
	for k, v := range kvmap.m {
		ok := f(k, v)
		if !ok {
			return
		}
	}
}

func NewKvMap(key, value any) iface.IMap[Key, Value] {

}
