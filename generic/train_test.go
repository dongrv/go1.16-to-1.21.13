package generic

import (
	"bytes"
	"github.com/go-playground/assert/v2"
	"strconv"
	"testing"
)

func TestTrain(t *testing.T) {
	t.Log("test train code\n")

	t.Run("LinkedList", func(t *testing.T) {
		list := &LinkedList[int]{}
		list.Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

		t.Run("add", func(t *testing.T) {
			assert.IsEqual(list.Nodes[1].next.value, 3)
		})
		t.Run("find", func(t *testing.T) {
			node := list.Find(func(i int) bool {
				return i == 9
			})
			if node == nil {
				t.Fatalf("node is nil")
			}
			assert.IsEqual(node.value, 9)
		})
		t.Run("remove", func(t *testing.T) {
			list.Remove(list.Nodes[0])
			assert.IsEqual(list.Nodes[3].next, 5)
		})
	})

	t.Run("processor", func(t *testing.T) {
		sp := StringProcessor{}
		stringResult := RunProcess[string, string](sp, []string{"Hello world!"})
		assert.Equal(t, stringResult[0], "Helloworld!")

		ip := IntProcessor{}
		intResult := RunProcess[int, int](ip, []int{10})
		assert.Equal(t, intResult[0], 100)
	})

	t.Run("decorator", func(t *testing.T) {
		logger := &LoggingDecorator[string]{}
		validator := &ValidatorDecorator[string]{
			validate: func(s2 string) bool {
				return s2 == "Alice"
			},
		}

		inputFunc := func(in string) string {
			return "Hello " + in + "!\n"
		}

		outputFunc := ApplyDecorator(inputFunc, logger, validator)
		assert.IsEqual(outputFunc("Alice"), "Hello Alice!")
	})

	t.Run("sort", func(t *testing.T) {
		slice := []string{"1", "3", "2", "8", "6", "9", "10"}
		SortByKey(slice, func(t string) int {
			i, _ := strconv.Atoi(t)
			return i
		}, true)
		t.Logf("slice: %v\n", slice)
	})

	t.Run("database", func(t *testing.T) {
		myString := MyString("Hello world!")
		dbi := &DBInstance[MyString]{store: make(map[string]MyString)}
		err := dbi.Put("Print", myString)
		if err != nil {
			t.Fatal(err)
		}
		data, err := dbi.Get("Print")
		if err != nil {
			t.Fatal(err)
		}
		body, _ := data.Marshal()
		t.Logf("data.marshal: %v\n", body)
		assert.Equal(t, body, []byte("Hello world!"))
	})

	t.Run("pool", func(t *testing.T) {
		p1 := Pool[bytes.Buffer]{}
		p1.Put(*bytes.NewBuffer([]byte("Hello world!")))
		v := p1.Get()
		t.Logf("pool get: %s\n", v.Bytes())

		type Foo struct{}
		p2 := Pool[Foo]{}
		p2.Put(Foo{})
		v2 := p2.Get()
		t.Logf("pool get: %#v\n", v2)
	})

	t.Run("vector", func(t *testing.T) {
		vector := &Vector[float64]{}
		a, b := []float64{1, 2, 3}, []float64{4, 5, 6}
		assert.IsEqual(vector.Add(a, b), []float64{5, 7, 9})
		assert.IsEqual(vector.Dot(a, b), 21)
	})

	t.Run("state-machine", func(t *testing.T) {
		type state uint
		const (
			Freezing state = iota
			Walking
			Running
			Flying
		)

		type Event = string

		moveEvent := "Move"
		fastMoveEvent := "Fast move"
		highSpeedEvent := "High Speed"

		sm := NewStateMachine[state, Event]()
		sm.AddTransition(Freezing, moveEvent, Walking)
		sm.AddTransition(Walking, fastMoveEvent, Running)
		sm.AddTransition(Running, highSpeedEvent, Flying)

		nextState, err := sm.Trigger(moveEvent)
		if err != nil {
			t.Fatal(err)
		}
		nextState, err = sm.Trigger(fastMoveEvent)
		if err != nil {
			t.Fatal(err)
		}
		nextState, err = sm.Trigger(highSpeedEvent)
		if err != nil {
			t.Fatal(err)
		}
		assert.IsEqual(nextState, Flying)
	})

	t.Run("container", func(t *testing.T) {
		loggerContainer := NewContainer[Logger]()

		logger := &ConsoleLogger{}

		loggerContainer.RegisterInstance(`logger`, logger)

		loggerContainer.RegisterFactory("service", NewService)

		service, err := loggerContainer.Get("service")
		if err != nil {
			t.Fatalf("container err: %v\n", err)
		}

		service.Log("Hello world!")

		// 另一种写法，容器分开写，实际意义不大，增加了冗余容器

		serviceContainer := NewContainer[IService]()
		serviceContainer.RegisterFactory("service", func(container *Container[IService]) IService {
			got, err := loggerContainer.Get("logger")
			if err != nil {
				return nil
			}
			return &Service{logger: got}
		})
		service2, err := serviceContainer.Get("service")
		if err != nil {
			t.Fatal(err)
		}
		service2.Do()
	})

}
