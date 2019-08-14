package main

type DefaultDict interface {
	Get(interface{}) interface{}
	Set(interface{}, interface{})
	Delete(interface{})
}

func NewDefaultDict(defaultVal interface{}) DefaultDict {
	return &defaultDict{defaultVal: defaultVal, m: make(map[interface{}]interface{})}
}

type defaultDict struct {
	defaultVal interface{}
	m          map[interface{}]interface{}
}

func (d *defaultDict) Get(key interface{}) interface{} {
	if v, ok := d.m[key]; ok {
		return v
	}
	return d.defaultVal
}

func (d *defaultDict) Set(key interface{}, val interface{}) {
	d.m[key] = val
}

func (d *defaultDict) Delete(key interface{}) {
	delete(d.m, key)
}
