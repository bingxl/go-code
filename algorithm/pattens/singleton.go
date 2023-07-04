package pattens

import "sync"

var instanceSingleton []int
var insSync sync.Once

func New() []int {
	// sync.Once 指挥执行一次, 从而得到单列模式
	insSync.Do(func() {
		instanceSingleton = []int{1, 2, 3}
	})

	return instanceSingleton
}
