package generic

// 编译器提示报错：

// type NewType[T int | string] T // Cannot use a type parameter as RHS in type declaration. 不允许新定义的泛型类型只有类型形参，无意义。
// type NewType1 [T * int][]T // 编译器会认为 T *int是一个表达式，在只有一个类型约束时，需要改成 [T *int,] 才能识别
// type NewType2 [T *int|*float32,] // 但是约束存在多个指针类型时，约束尾部增加逗号也是不合法的，无法识别为有效的泛型约束

// NewTypeOK 解决方法： 将类型约束包在interface{}，避免编译器误解为表达式
type NewTypeOK[T interface{ *int | *float32 }] []T

// 指定底层类型
// 具有相同底层类型的值能够相互赋值
// 使用 ~ 时有一定的限制：
// 	~ 后面的类型不能为接口
//	~ 后面的类型必须为基本类型

type Float interface {
	~float32 | ~float64 // 并集
}

type Int interface {
	~int8 | ~uint8 | ~int | ~uint // .. cmp.Ordered
}

type UnderlyingType[T Float | Int] []T // 这和[T ~float32 | ~float64 | ~int8 | ~uint8 | ~int | ~uint ]在功能上并没有什么不同。

type IByte interface { // 交集，IByte代表string和[]byte的交集 ~[]byte
	string
	~[]byte
}

type EmptySet interface { // 约束类型为空，没有相交的类型，可以编译通过但没有实际意义
	~float64
	int
}

type INumber3 interface {
	EmptySet
}

func Hello[T INumber3](t T) {
	println(t)
}
