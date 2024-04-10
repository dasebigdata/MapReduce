package main

import "mr/mr"

var Reduce mr.Reduce = func(key any, values []any) *mr.Pair {
	count := 0
	for _, value := range values {
		count += value.(int)
	}
	return &mr.Pair{Key: key, Value: count}
}
