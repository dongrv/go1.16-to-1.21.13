package generic

type Slice[T int | float32 | string] []T

/*

· 类型组成说明：
 	T -> 类型形参(Type parameter)
	int | float32 | string -> 类型约束(Type constraint)
 	[T int | float32 | string] -> 类型形参列表(Type parameter list)
 	Slice[T int | float32 | string] -> 泛型类型(Generic type)
· 泛型类型实例化(Instantiations)：
	ints := Slice[int]{1, 2, 3} // 此处的int为类型实参(Type argument)

*/

/*

· 等同于：
	type Slice_int 		[]int
	type Slice_float32 	[]float32
	type Slice_string 	[]string

· 上述的这个过程叫单态化（Monomorphization），单态化是指编译器为每个具体类型实例化一份泛型代码的过程。

*/

func Search[T int | float32 | string](slice Slice[T], find T) (ok bool) {
	if len(slice) == 0 {
		return
	}
	for _, v := range slice {
		if v == find {
			ok = true
			break
		}
	}
	return
}

// 编译器会在编译时针对每个使用该泛型函数的具体类型（如int、float64、string等）生成以下独立的代码：

func SearchInt(slice []int, find int) (ok bool) {
	if len(slice) == 0 {
		return
	}
	for _, v := range slice {
		if v == find {
			ok = true
			break
		}
	}
	return
}

func SearchFloat32(slice []float32, find float32) (ok bool) {
	if len(slice) == 0 {
		return
	}
	for _, v := range slice {
		if v == find {
			ok = true
			break
		}
	}
	return
}

func SearchString(slice []string, find string) (ok bool) {
	if len(slice) == 0 {
		return
	}
	for _, v := range slice {
		if v == find {
			ok = true
			break
		}
	}
	return
}

/*

单态化
· 优点：
	性能接近原生代码：无需运行时类型检查或动态调度，直接使用具体类型的方法和操作符。
	类型安全：所有类型检查都在编译时完成，避免运行时错误。
· 缺点：
	二进制文件膨胀：每个类型实例都会生成重复代码，但 Go 通过内联和优化技术减少了这种影响。

*/
