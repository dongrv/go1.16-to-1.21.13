package generic

import (
	"github.com/go-playground/assert/v2"
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
}
