package main

// golang 泛型

// 泛型函数
func GenFunc[T, V any](param T, paramB V) (T, V) {
	return param, paramB
}

// 泛型类型
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

// s := Stack[int64]{}
// s.Push(345)

// 泛型接口
type Container[T any] interface {
	len() int
	Add(T)
	Get() T
}
