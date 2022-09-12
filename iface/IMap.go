package iface

type IMap[Key, Value any] interface {
	Get(key Key) (value Value, ok bool)
	Set(key Key, value Value)
	Delete(key Key)
	//遍历kv-map，调用f
	Range(f func(key Key, value Value) bool)
}
