package generic

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
