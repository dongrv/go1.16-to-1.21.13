package generic

import (
	"fmt"
	"strconv"
)

// Min 泛型类型函数
func Min[T int | int32 | int64 | float32 | float64](ts ...T) T {
	var a T
	for _, t := range ts {
		if t > a {
			a = t
		}
	}
	return a
}

// Struct 泛型类型的结构体
type Struct[K int | string, V any] struct {
	ID    K
	Data  V
	Other float64
}

var s Struct[string, int] = Struct[string, int]{ID: "Alice", Data: 100, Other: 0.} // 泛型类型结构体实例化

// GroupStruct 泛型类型的复合类型结构体
type GroupStruct[T int | float32 | string, V []T] struct {
	One  T
	List V
}

var gs GroupStruct[int, []int] = GroupStruct[int, []int]{ // 类型实参T类型一旦确定为int，[]T也随之确定为[]int，存在强依赖关系
	One:  1,
	List: []int{1, 2, 3},
}

// Map 泛型类型的 map
type Map[K int, V string | float64] map[K]V

var m Map[int, string] = Map[int, string]{100: "A+"} // 泛型类型map实例化

// Chan 泛型类型 channel
type Chan[T int | float64] chan T

var c Chan[float64] = make(Chan[float64], 100) // 泛型类型channel实例化

// IExample 泛型类型接口定义
type IExample[T any] interface {
	Print(T)
}

type Example[T int] struct{}

func (e Example[T]) Print(t T) {} // 实现泛型类型接口

// 空接口：包含所有类型的类型集
// type any = interface{} // builtin alias[别名]

type NewSlice[T any] []T // 支持所有类型元素的泛型类型切片

// 从Go1.18开始接口分为两种类型：
// 	- 基本接口类型（Basic interface），向前兼容，识别特征：接口类型只有方法签名或空接口。
// 	- 一般接口类型（General interface），新特性，识别特征：接口类型同时包含方法签名和类型，或只有类型集合。

// 基本接口类型（Basic interface）
// 基本接口类型可以放在类型约束中

type IPrint interface {
	Print(string)
}

type Foo struct{}

func (f Foo) Print(s string) { println(s) }

type IOPrint[T IPrint] []T // 用基本接口类型 IPrint 作为约束条件定义一个泛型类型 IOPrint

var iop = IOPrint[Foo]{Foo{}, Foo{}, Foo{}}

func LoopBasicInterface() {
	for i, foo := range iop {
		foo.Print("BasicInterface: " + strconv.Itoa(i))
	}
}

// 一般接口类型
// 一般接口类型不能用于定义变量，只能用于泛型的类型约束中。
// 分析以下代码示例：

// 基本接口类型，只有方法签名，没有类型

type DataProcessor[T any] interface {
	Process(T) T
	Save(T) error
}

// CSVProcessor 实现了DataProcessor[string]类型接口
type CSVProcessor struct{}

func (csv CSVProcessor) Process(data string) string { return data }
func (csv CSVProcessor) Save(data string) error     { return nil }

// 一般接口类型，存在类型并集

type M int

type DataProcessor2[T any] interface {
	~int32 | ~int

	Process(T) T
	Save(T) error
}

func (m *M) Process(i int) int {
	return i
}
func (m *M) Save(i int) error {
	return nil
}

// 这里 String 也实现了 DataProcessor2[string] 接口
// Go 语言中，类型隐式实现接口的条件是：类型的方法集合完全包含接口定义的所有方法，且方法的签名（参数类型、返回类型）完全一致。

type String string

func (s String) Process(i string) string {
	return ""
}

func (s String) Save(i string) error {
	return nil
}

type NumberI64OrI interface {
	int64 | ~int
}

func Bar[T NumberI64OrI](t T) T {
	return t
}

var result = Bar(M(1)) // Bar接收底层类型为int的类型，M底层类型为int，所以符合条件

// IntNumber 实现了基于底层类型为int、包含 Process(T) T 、Save(T) error 方法的接口 DataProcessor2[T any]
type IntNumber int

func (i IntNumber) Process(number int) int { return number }
func (i IntNumber) Save(number int) error  { return nil }

var p DataProcessor[string]

// var p DataProcessor2[int] // 一般接口类型不能定义变量类型：Interface includes constraint elements '~int', '~struct{ Data interface{} }', can only be used in type parameters

// 其他定义接口类型的限制：

// 1. 使用 | 连接并集时，底层类型不能重叠，但是相交的类型是接口的话，允许连接

type INumber interface {
	// ~int | int // overlapping terms int and ~int
}

type INumber2 interface {
	~int | interface{ int }
}

// 2. 类型的并集不能有类型形参

type Type[T ~int | ~string] interface {
	// int | T // 不能嵌入一个类型形参：Cannot embed a type parameter
}

// 3. 接口不能直接或间接的并入自己

// 直接 Invalid recursive interface 'Bad'

//type Bad interface {
//	Bad
//}

// 间接 Invalid recursive interface 'Bad1' Bad1 → Bad2 → Bad

//type Bad1 interface {
//	Bad2
//}
//
//type Bad2 interface {
//	Bad1
//}

// 4. 接口的并集成员个数大于1的时候，不能直接或间接的并入comparable接口

type OK interface {
	comparable
}

type OK2 interface {
	comparable
	~int // 这样是可以的，OK2是 comparable 和 int 的交集，实际OK2就是 ~int 的一般泛型类型，这种交集没有实际意义
}

//type Bad interface {
//	~int | comparable // Cannot use comparable in union
//}

// 5. 带方法的接口，无论是基本接口类型还是一般接口类型，都不能写入接口的并集中

// 允许的情况：

type Allow interface {
	int | ~float32
}

type Allow2 interface {
	~int | Allow
}

// 不允许的情况：

type IAllow interface {
	Allow()
}

//type INotAllow interface {
//	~int | IAllow // Cannot use interfaces with methods in union
//}

// 类型约束为：结构体、底层为结构体
// 类型约束为结构体 struct{} ，则实现的具体结构体必须保持和约束的结构体类型完全一致（字段名字、数量、类型、顺序、方法），才能视为满足类型约束条件
// 类型约束为底层结构体 ~struct{} ，所有类型的底层类型实现了约束规则（同struct{}）视为满足了类型约束条件

type IDefault interface {
	~struct {
		ID   int
		Name string
	}
	Hi() string
}

type Person struct {
	ID   int
	Name string
}

func (p Person) Hi() string {
	return fmt.Sprintf("Hi, my name is %s and ID is %d", p.Name, p.ID)
}

type American Person

func (a American) Hi() string {
	return fmt.Sprintf("Hi, my name is %s and ID is %d. I'm an American.", a.Name, a.ID)
}

func Hi[T IDefault](t T) {
	println(t.Hi())
}

// 综上：太多了容易混淆，只需要按照面向编译器编程，编译器没报错就说明泛型代码没问题。
