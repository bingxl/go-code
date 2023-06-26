# golang 与 OOP

## golang 中的对象

golang 使用结构体模拟类与对象

```go
type Bike struct {
    color string // 首字母小写表示私有属性
    Name string // 首字母大写表示公有属性
}

// 注册结构体的方法, 首字母大写表示该方法对外公开
func (b *Bike) Move() string {
    return b.color + "自行车"
}
```

- 封装: 首字母大小写表示方法/属性的公有私有权限
- 继承: 使用内嵌的方式对结构体进行组合
- 多态: 使用 interface 实现

```go

// ---------------封装------------
type Person struct {
    name string
}
func (p *Person) Walk() {
    fmt.Println(p.name + " is working ")
}

// ---------------继承---------------
type Chinese struct {
    Person // 继承
    skin string
}
func (c *Chinese) GetName() string {
    // 已经可以直接访问内嵌结构体的成员, 也可以通过结构体名访问 比如 c.Person.name
    return c.name
}

// ---------------多态---------------
type Human interface {
    Speak()
}

// 实现接口, 没有显示说明实现Human接口, 只要接口中的方法结构体都实现了则视为实现了接口
type Chinese struct {
    name string
    Language string
}

func (c *Chinese) Speak() {
    fmt.Println("chinese " + c.name + " speak " + c.Language)
}

// 使用接口类型
var h Human = new(Chinese)
h.Speak()

```

## 五大基本原则
- 单一功能原则
- 开闭原则: 对扩展是打开的,对修改时关闭的
- 里氏替换原则: 子类可以替换父类,而逻辑不变, golang中是面向接口编程
- 接口隔离原则: 接口最小化
- 依赖反转原则: 不依赖具体实现,依赖接口开发

## 设计模式分类
- 创建者型: 单例模式/简单工厂模式/工厂方法模式/抽象工厂模式/建造者模式/原型模式
- 结构型: 代理模式/适配器模式/装饰模式/桥接模式/组合模式/亨元模式/外观模式
- 行为型: 观察者模式/模板方法模式/命令模式/状态模式/职责链模式/解释器模式/中介者模式/访问者模式/策略模式/备忘录模式/迭代器模式