package datastruct

import (
	"errors"
	"strconv"
)

// 顺序表

type SqList struct {
	value     []any
	maxLength int
	length    int
}

func NewSqlList(max int) *SqList {
	return &SqList{
		value:     make([]any, max),
		maxLength: max,
		length:    0,
	}
}

func (s SqList) GetLen() int {
	return s.length
}
func (s SqList) GetMaxLen() int {
	return s.maxLength
}

// 在pos位置插入数据
func (s *SqList) Insert(pos int, value any) error {
	if s.length >= s.maxLength {
		return errors.New("顺序表已满")
	}
	if pos < 0 || pos >= s.maxLength || pos > s.length {
		return errors.New("索引越界 pos: " + strconv.Itoa(pos))
	}
	for i := s.length; i > pos; i-- {
		s.value[i] = s.value[i-1]
	}
	s.value[pos] = value
	s.length++

	return nil
}

// 查找值对应的位置, -1 表示不存在
func (s *SqList) IndexOf(v any) int {
	for i := 0; i < s.length; i++ {
		if s.value[i] == v {
			return i
		}
	}
	return -1
}

func (s *SqList) Get(pos int) (any, error) {
	if pos < 0 || pos > s.length {
		return nil, errors.New("未知参数错误")
	}
	return s.value[pos], nil
}

func (s *SqList) Each(f func(i int, v any)) {
	for i := 0; i < s.length; i++ {
		f(i, s.value[i])
	}
}

// 将循序表中的传入 f 函数, 返回由 f 函数返回值组成的 切片
func (s *SqList) Map(f func(i int, v any) any) []any {
	v := make([]any, s.length)
	for i := 0; i < s.length; i++ {
		v[i] = f(i, s.value[i])
	}
	return v
}

// 删除节点
func (s *SqList) Del(pos int) error {
	if pos < 0 || pos >= s.length {
		return errors.New("位置参数错误")
	}

	for i := pos; i < s.length-1; i++ {
		s.value[i] = s.value[i+1]
	}
	s.length--
	return nil
}

// 顺序表扩容, 若新容量max 小于原顺序表的容量,可能会丢失数据
func (s *SqList) ExtendCap(max int) {
	newV := make([]any, max)
	copy(newV, s.value)
	s.value = newV
	s.maxLength = max
}
