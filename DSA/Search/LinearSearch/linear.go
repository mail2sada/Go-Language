package linearsearch

import "reflect"

func LinearSearch[T any](dataList []T, key interface{}) (pos int, found bool) {
	pos = -1
	found = false
	for idx, val := range dataList {
		if reflect.DeepEqual(val, key) {
			pos = idx
			found = true
			return
		}
	}
	return
}
