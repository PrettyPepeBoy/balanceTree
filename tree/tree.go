package tree

import "fmt"

type Tree struct {
	Val     int
	left    *Tree
	right   *Tree
	parent  *Tree
	BFactor int //определяет разницу высот правого и левого поддерева. Если положительно, то левое больше
}

func CreateTree(value int) Tree {
	var t Tree
	t.Val = value
	t.BFactor = 0
	return t
}

// малый поворот вправо
func (t *Tree) rightTurn() *Tree {
	t.BFactor = 0
	if t.left.BFactor > 0 {
		t.left.BFactor--
		t.Val, t.left.Val = t.left.Val, t.Val
		t.right = t.left
		t.left = t.left.left
		t.left.parent = t
		t.right.left = nil
		return t
	} else {
		t.left.BFactor++
		t.Val, t.left.right.Val = t.left.right.Val, t.Val
		t.left.right.parent = t
		t.right = t.left.right
		t.left.right = nil
		return t
	}
}

// малый поворот налево
func (t *Tree) leftTurn() *Tree {
	t.BFactor = 0
	if t.right.BFactor < 0 {
		t.right.BFactor++
		t.Val, t.right.Val = t.right.Val, t.Val
		t.left = t.right
		t.right = t.right.right
		t.right.parent = t
		t.left.right = nil
		return t
	} else {
		t.right.BFactor--
		t.Val, t.right.left.Val = t.right.left.Val, t.Val
		t.right.left.parent = t
		t.left = t.right.left
		t.right.left = nil
		return t
	}
}

// логическая конструкция необходима для того, чтобы значение фактора баланса не переопределялось при удалении сегмента
func (t *Tree) setHeight(put bool) *Tree {
	if t.parent == nil {
		if t.BFactor == 2 {
			t.rightTurn()
			return t
		}
		if t.BFactor == -2 {
			t.leftTurn()
			return t
		}
		return t
	}

	if t.Val > t.parent.Val {
		if put {
			t.parent.BFactor--
		} else {
			t.parent.BFactor++
		}

		if t.parent.BFactor == -2 {
			t.parent.leftTurn()
			return t
		}

		if t.parent.BFactor == 0 {
			return t
		} else {
			return t.parent.setHeight(put)
		}
	} else {
		if put {
			t.parent.BFactor++
		} else {
			t.parent.BFactor--
		}

		if t.parent.BFactor == 2 {
			t.parent.rightTurn()
			return t
		}

		if t.parent.BFactor == 0 {
			return t
		} else {
			return t.parent.setHeight(put)
		}
	}
}

func (t *Tree) Put(value int) bool {
	if t.Val > value {
		if t.left == nil {
			t.left = &Tree{
				Val:     value,
				parent:  t,
				BFactor: 0,
			}
			t.left.setHeight(true)
			return false
		}
		return t.left.Put(value)
	}

	if t.Val < value {
		if t.right == nil {
			t.right = &Tree{
				Val:     value,
				parent:  t,
				BFactor: 0,
			}
			t.right.setHeight(true)
			return false
		}
		return t.right.Put(value)
	}

	return true
}

func (t *Tree) FindMax() *Tree {
	if t.right == nil {
		return t
	}
	return t.right.FindMax()
}

func (t *Tree) FindMin() *Tree {
	if t.left == nil {
		return t
	}
	return t.left.FindMin()
}

func (t *Tree) FindElem(value int) *Tree {

	if t.Val > value {
		if t.left != nil {
			return t.left.FindElem(value)
		} else {
			return nil
		}
	}

	if t.Val < value {
		if t.right != nil {
			return t.right.FindElem(value)
		} else {
			return nil
		}
	}

	return t
}

func (t *Tree) Delete(value int) bool {
	node := t.FindElem(value)
	if node == nil {
		return false
	}

	if node.right == nil {
		if node.parent.Val > node.Val {
			node.parent.left = node.left
			if node.left != nil {
				node.left.parent = node.parent
			}
			node.parent.BFactor--
			node.parent.setHeight(false)
		} else {
			node.parent.right = node.left
			if node.left != nil {
				node.left.parent = node.parent
			}
			node.parent.BFactor++
			node.parent.setHeight(false)
		}
		return true
	}

	minNode := node.right.FindMin()
	if minNode.parent.Val != node.Val {
		node.Val = minNode.Val
		minNode.parent.left = minNode.right
		if minNode.right != nil {
			minNode.right.parent = minNode.right
		}
		return true
	}
	node.Val = minNode.Val
	node.right = minNode.right
	if minNode.right != nil {
		minNode.right.parent = node
	}
	return true
}

func (t *Tree) Print() {
	if t == nil {
		return
	}
	t.left.Print()
	fmt.Println(t)
	t.right.Print()
}
