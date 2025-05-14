package generic

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
