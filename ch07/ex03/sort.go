// Package treesort provides insertion sort using an unbalanced binary tree.
package treesort

import "strconv"

type tree struct {
	value       int
	left, right *tree
}

// Sortはvalues内の値をその中でソートします。
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValuesはtの要素をvaluesの正しい順序に追加し、
// 結果のスライスを返します。
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// return &tree{value: value}と同じ
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

// Stringはツリー内の値の列を表示します。
func (t *tree) String() string {
	return "{" + t.valString() + "}"
}

func (t *tree) valString() string {
	if t == nil {
		return ""
	}

	var ret string

	if t.left != nil {
		ret = t.left.valString() + " "
	}

	ret += strconv.Itoa(t.value)

	if t.right != nil {
		ret += " " + t.right.valString()
	}

	return ret
}
