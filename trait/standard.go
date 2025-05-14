package trait

import (
	"cmp"
	"context"
	"errors"
	"fmt"
	"maps"
	"slices"
	"time"
)

// Slices 新增的切片操作包函数
func Slices() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slices.Sort(s)                  // .SortFunc
	slices.Index(s, 2)              // .IndexFunc
	_ = slices.Clone(s)             // shallow clone：浅克隆
	ns := slices.Clip(s[:len(s)/2]) // 等于 s[:len(s):len(s)]
	fmt.Printf("s -> cap:%d len:%d \nns -> cap %d len:%d\n", cap(s), len(s), cap(ns), len(ns))
	slices.Delete(ns, 2, 3)
	// ... 自行探索其他函数
}

// Maps 新增的map操作包函数
func Maps() {
	m := map[int]int{0: 0, 1: 1, 2: 2}
	nm := make(map[int]int)
	maps.Copy(nm, m)
	for k, v := range nm {
		fmt.Printf("k=%d v=%d\n", k, v)
	}
	if !maps.Equal(nm, m) {
		panic("unexpected not equal")
	}
	maps.DeleteFunc(nm, func(i int, i2 int) bool { return nm[i] == 1 })
	for k, v := range nm {
		fmt.Printf("deleteFunc: k=%d v=%d\n", k, v)
	}
}

// Cmp 提供深度比较和差异报告功能
func Cmp() {
	cmp.Compare(1, 2) // 只能比较 cmp.Ordered 泛型类型的参数
	cmp.Less(1, 2)
	// 更高Go版本支持更多的函数，自己可以研究下
}

// Context 新增允许设置取消原因的上下文
func Context() {
	ctx, causeCancel := context.WithCancelCause(context.Background())
	defer causeCancel(errors.New("process exited"))

	ctx2, cancel := context.WithDeadlineCause(context.Background(), time.Now().Add(100*time.Second), errors.New("deadline"))
	defer cancel()

	_, _ = ctx, ctx2
}

// PanicRecover
// panic: 提升了执行性能，通过复用内存对象和减少锁范围，降低了调用开销。堆栈信息增强，增加显示defer匿名函数。提高并发安全性。
// recover: 通过缓存和快速路径优化，减少条件判断次数。
func PanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic-recover:%v\n", r)
		}
	}()
	panic(nil)
}
