package factory

import (
	"fmt"
	"testing"
)

// 测试手工创建类时是否正确工作
func TestSampleProductCreate(t *testing.T) {
	product1 := &sampleProduct1{}
	product1.SetName("p1")
	fmt.Println("product1's name: ", product1.GetName())

	product2 := &sampleProduct2{}
	product2.SetName("p2")
	fmt.Println("product2's name: ", product2.GetName())
}

func TestSampleProductFactory(t *testing.T) {
	product1 := NewSampleProductFactory(p1)
	product2 := NewSampleProductFactory(p2)

	product1.SetName("p1")
	product2.SetName("p2")

	fmt.Println("p1's name: ", product1.GetName(), "; p2's name: ", product2.GetName())
}
