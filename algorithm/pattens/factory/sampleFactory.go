package factory

// 简单工厂模式, 又名静态工厂模式(static Factory Method)
// 简单工厂模式通过不同的参数返回不同的实例
// 专门定义一个类负责创建其他类的实例

type sampleProduct interface {
	SetName(name string)
	GetName() string
}

// --------------实现产品1 --------------------
type sampleProduct1 struct {
	name string
}

func (p1 *sampleProduct1) SetName(name string) {
	p1.name = name
}
func (p1 *sampleProduct1) GetName() string {
	return p1.name
}

// ------------实现产品2-------------------------
type sampleProduct2 struct {
	name string
}

func (p2 *sampleProduct2) SetName(name string) {
	p2.name = name
}
func (p2 *sampleProduct2) GetName() string {
	return p2.name
}

// ------------------简单工厂类------------------
type sampleProductType int

const (
	p1 = iota
	p2
)

func NewSampleProductFactory(productType sampleProductType) sampleProduct {
	// 根据传入的不同参数返回不同的实例
	switch productType {
	case p1:
		{
			return &sampleProduct1{}
		}
	case p2:
		{
			return &sampleProduct2{}
		}
	default:
		return nil
	}
}

// 简单工厂模式的优缺点
// 优点: 创建时只需要知道工厂类和需要传入的参数,不需要关心具体的类是怎么创建的, 实现了解耦
// 缺点: 违背了开闭原则
// 适合: 创建对象比较少
