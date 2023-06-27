package main

import "fmt"

type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, args ...string) (int, error)
}

// 结构体中嵌套了接口, 那么结构体会被认为实现了接口, 但是只能调用结构实现了的方法
// 不能调用在接口中结构体还未实现的方法

type Fake struct {
	Stmt
}

func (f Fake) Exec(stmt string, args ...string) (int, error) {
	return 4, nil
}

// 结构体自定义的方法(非继承接口中的) 在 typ Fake1 Fake 时 Fake1 不会继承该方法
func (f Fake) Private() (int, error) {
	return 4, nil
}

func interfaceTest() {

	f := Fake{}

	// f能成功赋值给 Stmt 型变量
	var stmt Stmt = f
	r, err := stmt.Exec("hello")
	fmt.Println(r, err)

	// 结构体中还未实现NumInput 方法,  以下调用会出错, invalid Memory
	// fmt.Println(f.NumInput())

	// define 类型
	type Fake1 Fake
	type Stmt1 Stmt

	// Fake1 未继承 Fake 中自定义/自己实现的方法, 以下方法会出错
	// Fake1{}.Exec("")
	fmt.Println("\n ------define 类型测试")
	DumpMethodSet(&Fake{})
	DumpMethodSet(&Fake1{})
	DumpMethodSet((*Stmt)(nil))
	DumpMethodSet((*Stmt1)(nil))

	// 类型别名
	fmt.Println("\n -------类型别名测试")
	type FakeAlias = Fake
	type StmtAlias = Stmt

	// 以下语句不会出错
	FakeAlias{}.Exec("")
	DumpMethodSet(&Fake{})
	DumpMethodSet(&FakeAlias{})
	DumpMethodSet((*Stmt)(nil))
	DumpMethodSet((*StmtAlias)(nil))

}
