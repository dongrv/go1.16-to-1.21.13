package generic

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"strings"
	"sync"
	"time"
)

// 双向链表

type LinkedList[T comparable] struct {
	Nodes []*Node[T]
}

type Node[T comparable] struct {
	prev  *Node[T]
	value T
	next  *Node[T]
}

func (list *LinkedList[T]) Add(ts ...T) {
	if len(list.Nodes) == 0 {
		list.Nodes = make([]*Node[T], 0, 100)
	}
	for _, t := range ts {
		node := &Node[T]{value: t}
		if list.len() > 0 {
			last := list.Nodes[len(list.Nodes)-1]
			last.next, node.prev = node, last // 更新连表指针
		}
		list.Nodes = append(list.Nodes, node)
	}
}

func (list *LinkedList[T]) Remove(node *Node[T]) {
	if len(list.Nodes) == 0 || node == nil {
		return
	}
	switch len(list.Nodes) {
	case 0:
		return
	case 1:
		list.Nodes[0] = nil
		list.Nodes = nil
		return
	default:
		for i, n := range list.Nodes {
			if n.value == node.value && n.prev == node.prev && n.next == node.next {
				if n.prev != nil && n.next != nil {
					n.next.prev, n.prev.next = n.prev, n.next
				}
				if n.prev == nil {
					n.next.prev = nil
				}
				if n.next == nil {
					n.prev.next = nil
				}
				list.Nodes[i] = nil
				list.Nodes = append(list.Nodes[0:i], list.Nodes[i+1:]...)
			}
		}
	}
}

func (list *LinkedList[T]) len() int {
	return len(list.Nodes)
}

func (list *LinkedList[T]) Find(f func(T) bool) *Node[T] {
	if list.len() == 0 {
		return nil
	}
	for node := list.Nodes[0]; node.next != nil; node = node.next {
		if f(node.value) {
			return node
		}
	}
	return nil
}

type Processor[T, R any] interface {
	Process(T) R
}

type StringProcessor struct{}

func (s StringProcessor) Process(in string) string {
	return strings.ReplaceAll(in, " ", "")
}

type IntProcessor struct{}

func (i IntProcessor) Process(in int) int {
	return in * in
}

func RunProcess[T, R any](p Processor[T, R], ts []T) []R {
	outs := make([]R, 0, len(ts))
	for _, v := range ts {
		outs = append(outs, p.Process(v))
	}
	return outs
}

// 装饰器模式泛型实现

type Decorator[T any] interface {
	Decorate(func(T) T) func(T) T
}

type LoggingDecorator[T any] struct{}

func (l *LoggingDecorator[T]) Decorate(f func(T) T) func(T) T {
	return func(t T) T {
		v := f(t)
		fmt.Printf("logging: %v", v)
		return v
	}
}

type ValidatorDecorator[T any] struct {
	validate func(T) bool
}

func (v *ValidatorDecorator[T]) Decorate(f func(T) T) func(T) T {
	return func(t T) T {
		if !v.validate(t) {
			var zero T
			fmt.Println("validate fail")
			return zero
		}
		return f(t)
	}
}

func ApplyDecorator[T any](f func(T) T, decorators ...Decorator[T]) func(T) T {
	for _, decorator := range decorators {
		f = decorator.Decorate(f)
	}
	return f
}

func SortByKey[T any, K cmp.Ordered](ts []T, keyFunc func(T) K, ascending bool) {
	slices.SortFunc(ts, func(a, b T) int {
		keyA, keyB := keyFunc(a), keyFunc(b)
		if keyA < keyB {
			if ascending {
				return -1
			}
			return 1
		}
		if keyA > keyB {
			if ascending {
				return 1
			}
			return -1
		}
		return 0
	})
}

// 泛型数据库接口

type Marshaller interface {
	~struct{} | ~string
	Marshal() ([]byte, error)
}

type DB[T Marshaller] interface {
	Get(id string) (T, error)
	Put(id string, data T) error
}

type MyString string

func (m MyString) Marshal() ([]byte, error) {
	return []byte(m), nil
}

type DBInstance[T Marshaller] struct {
	store map[string]T
}

func (db *DBInstance[T]) Get(id string) (T, error) {
	if v, ok := db.store[id]; ok {
		return v, nil
	}
	var zero T
	return zero, errors.New("no data")
}

func (db *DBInstance[T]) Put(id string, data T) error {
	db.store[id] = data
	return nil
}

// Pool

type Pool[T ~struct{} | bytes.Buffer] struct {
	pool sync.Pool
}

func (p *Pool[T]) Get() T {
	v, ok := p.pool.Get().(T) // 需要显示的断言
	if ok {
		return v
	}
	var zero T
	return zero
}

func (p *Pool[T]) Put(t T) {
	p.pool.Put(t)
}

// 向量运算

type IVectorNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | uintptr |
		~float32 | ~float64
}

type Vector[T IVectorNumber] struct{}

// Add 向量加法
func (v *Vector[T]) Add(a, b []T) []T {
	if len(a) != len(b) {
		return nil
	}
	result := make([]T, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}
	return result
}

// Dot 向量点积
func (v *Vector[T]) Dot(a, b []T) T {
	var zero T
	if len(a) != len(b) {
		return zero
	}
	var sum T
	for i := range a {
		sum += a[i] + b[i]
	}
	return sum
}

// Normalize 向量归一化
func (v *Vector[T]) Normalize(vec []T) []T {
	magnitude := v.Magnitude(vec)
	result := make([]T, len(vec))
	for i := range vec {
		result[i] = vec[i] / magnitude
	}
	return result
}

// Magnitude 向量模长
func (v *Vector[T]) Magnitude(vec []T) T {
	var sum T
	for _, v := range vec {
		sum += v * v
	}
	return T(math.Sqrt(float64(sum)))
}

// 泛型状态机

type State interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type StateMachine[S State, E comparable] struct {
	mu     sync.RWMutex
	state  S             // 当前状态
	bundle map[S]map[E]S // 起始状态->触发事件->结束状态
}

func NewStateMachine[S State, E comparable]() *StateMachine[S, E] {
	return &StateMachine[S, E]{bundle: make(map[S]map[E]S)}
}

func (s *StateMachine[S, E]) AddTransition(from S, event E, to S) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.bundle[from]; !ok {
		s.bundle[from] = make(map[E]S)
	}
	if _, ok := s.bundle[from][event]; !ok {
		s.bundle[from][event] = to
	}
}

func (s *StateMachine[S, E]) Trigger(event E) (S, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var zero S
	if _, ok := s.bundle[s.state]; !ok {
		return zero, errors.New("state error")
	}
	if _, ok := s.bundle[s.state][event]; !ok {
		return zero, errors.New("not found event")
	}
	s.state = s.bundle[s.state][event]
	return s.state, nil
}

// 泛型依赖容器

type Container[T any] struct {
	mu       sync.RWMutex
	registry map[string]interface{}
}

func NewContainer[T any]() *Container[T] {
	return &Container[T]{registry: map[string]interface{}{}}
}

func (c *Container[T]) RegisterInstance(name string, instance T) {
	c.mu.Lock()
	c.registry[name] = instance
	c.mu.Unlock()
}

func (c *Container[T]) RegisterFactory(name string, fn func(container *Container[T]) T) {
	c.mu.Lock()
	c.registry[name] = fn
	c.mu.Unlock()
}

func (c *Container[T]) Get(name string) (T, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if val, ok := c.registry[name]; ok {
		switch v := val.(type) {
		case T:
			return v, nil
		case func(container *Container[T]) T:
			return v(c), nil
		default:
			var zero T
			return zero, fmt.Errorf("type mismatch for %s: expected: %T or func(*Container) %T", name, zero, zero)
		}
	}
	var zero T
	return zero, fmt.Errorf("no registration found %s", name)
}

type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(msg string) {
	fmt.Printf("%s %s\n", time.Now().Format(time.DateTime), msg)
}

type IService interface {
	Do()
}

type Service struct {
	logger Logger
}

func (s *Service) Do() {
	s.logger.Log("Server is running")
}

func (s *Service) Log(msg string) {
	s.logger.Log("[Service]" + msg)
}

func NewService(c *Container[Logger]) Logger {
	logger, err := c.Get("logger")
	if err != nil {
		return nil
	}

	return &Service{logger: logger}
}
