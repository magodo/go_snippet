package mapkey

import "reflect"

func Mapkey(i interface{}, val interface{}) (key interface{}, ok bool) {
	iv := reflect.ValueOf(i)
	if iv.Kind() != reflect.Map {
		return nil, false
	}
	for _, k := range iv.MapKeys() {
		if iv.MapIndex(k).Interface() == val {
			return k.Interface(), true
		}
	}
	return nil, false
}
