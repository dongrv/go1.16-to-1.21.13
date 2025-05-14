package generic

// 编译器提示报错：

// type NewType[T int | string] T // Cannot use a type parameter as RHS in type declaratio. 不允许新定义的泛型类型只有类型形参，无意义。
// type NewType1 [T * int][]T // 编译器会认为 T *int是一个表达式，在只有一个类型约束时，需要改成 [T *int,] 才能识别
// type NewType2 [T *int|*float32,] // 但是约束存在多个指针类型时，约束尾部增加逗号也是不合法的，无法识别为有效的泛型约束

// NewTypeOK 解决方法： 将类型约束包在interface{}，避免编译器误解为表达式
type NewTypeOK[T interface{ *int | *float32 }] []T

type DependOnUnderlyingType[T int | string] int
