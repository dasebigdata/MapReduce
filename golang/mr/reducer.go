package mr

type Reduce func(key any, values []any) *Pair
