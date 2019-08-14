package tree

import (
	"errors"
)

type Node struct {
	Value  interface{}
	Parent *Node
	Left   *Node
	Right  *Node
	Comp   CompFunc
}

type Bst struct {
	Root *Node
	comp CompFunc
}

// CompFunc returns:
// 0 : if two value equals
// -1: if a < b
// 1 : if a > b
type CompFunc func(a interface{}, b interface{}) int

/////////////////////// INIT /////////////////////
func (tree *Bst) RegisterCompFunc(f CompFunc) {
	tree.comp = f
}

/////////////////////// INSERT /////////////////////

func (node *Node) Insert(data interface{}) (err error) {
	if node == nil {
		return errors.New("nil node")
	}

	switch node.Comp(data, node.Value) {
	case 0:
		return
	case -1:
		if node.Left == nil {
			node.Left = &Node{
				Parent: node,
				Value:  data,
				Left:   nil,
				Right:  nil,
				Comp:   node.Comp,
			}
			return
		}
		return node.Left.Insert(data)
	case 1:
		if node.Right == nil {
			node.Right = &Node{
				Parent: node,
				Value:  data,
				Left:   nil,
				Right:  nil,
				Comp:   node.Comp,
			}
			return
		}
		return node.Right.Insert(data)
	default:
		return errors.New("malformed comp function")
	}
}

func (tree *Bst) Insert(data interface{}) (err error) {
	if tree.comp == nil {
		return errors.New("comp function not registered")
	}

	if tree.Root == nil {
		tree.Root = &Node{
			Parent: nil,
			Value:  data,
			Left:   nil,
			Right:  nil,
			Comp:   tree.comp,
		}
		return nil
	}
	return tree.Root.Insert(data)
}

/////////////////////// FIND /////////////////////

type dataContext struct {
	val    interface{}
	parent *Node
}

func (node *Node) FindSmallestNode() (*Node, error) {
	if node == nil {
		return nil, errors.New("node is nil")
	}
	if node.Left == nil {
		return node, nil
	}
	n := node.Left
	for ; n.Left != nil; n = n.Left {
	}
	return n, nil
}

func (node *Node) FindBiggestNode() (*Node, error) {
	if node == nil {
		return nil, errors.New("node is nil")
	}
	if node.Right == nil {
		return node, nil
	}
	n := node.Right
	for ; n.Right != nil; n = n.Right {
	}
	return n, nil
}

func (node *Node) FindNearestAscendant() (*Node, error) {
	if node == nil {
		return nil, errors.New("node is nil")
	}
	if node.Right != nil {
		return node.Right.FindSmallestNode()
	}
	for n, p := node, node.Parent; p != nil; n, p = p, p.Parent {
		if n == p.Left {
			return p, nil
		}
	}
	return nil, errors.New("not found")
}

func (node *Node) FindNearestDescendant() (*Node, error) {
	if node == nil {
		return nil, errors.New("node is nil")
	}
	if node.Left != nil {
		return node.Left.FindBiggestNode()
	}
	for n, p := node, node.Parent; p != nil; n, p = p, p.Parent {
		if n == p.Right {
			return p, nil
		}
	}
	return nil, errors.New("not found")
}

func (node *Node) FindNearestNodeAgainstValue(dataCtx dataContext) (*Node, error) {
	if node == nil {
		if dataCtx.parent == nil {
			return nil, errors.New("not found")
		}
		return dataCtx.parent, nil
	}
	dataCtx.parent = node
	switch node.Comp(dataCtx.val, node.Value) {
	case 0:
		return node, nil
	case -1:
		return node.Left.FindNearestNodeAgainstValue(dataCtx)
	case 1:
		return node.Right.FindNearestNodeAgainstValue(dataCtx)
	default:
		return nil, errors.New("malformed comp function")
	}
}

func (tree *Bst) FindNearestDescendantNodeAgainstValue(data interface{}) (*Node, error) {
	n, err := tree.FindNearestNodeAgainstValue(data)
	if err != nil {
		return nil, err
	}
	if tree.comp(n.Value, data) < 0 {
		return n, nil
	}
	return n.FindNearestDescendant()
}

func (tree *Bst) FindNearestAscendantNodeAgainstValue(data interface{}) (*Node, error) {
	n, err := tree.FindNearestNodeAgainstValue(data)
	if err != nil {
		return nil, err
	}
	if tree.comp(n.Value, data) > 0 {
		return n, nil
	}
	return n.FindNearestAscendant()
}

func (tree *Bst) FindNearestNodeAgainstValue(data interface{}) (*Node, error) {
	if tree.comp == nil {
		return nil, errors.New("comp function not registered")
	}
	if tree.Root == nil {
		return nil, errors.New("empty tree")
	}
	return tree.Root.FindNearestNodeAgainstValue(dataContext{data, nil})
}

func (node *Node) Find(data interface{}) (*Node, error) {
	if node == nil {
		return nil, errors.New("not found")
	}
	switch node.Comp(data, node.Value) {
	case 0:
		return node, nil
	case -1:
		return node.Left.Find(data)
	case 1:
		return node.Right.Find(data)
	default:
		return nil, errors.New("malformed comp function")
	}
}

func (tree *Bst) Find(data interface{}) (*Node, error) {
	if tree.comp == nil {
		return nil, errors.New("comp function not registered")
	}
	if tree.Root == nil {
		return nil, errors.New("empty tree")
	}
	return tree.Root.Find(data)
}

/////////////////////// DELETE /////////////////////

func (node *Node) findMax() *Node {
	if node == nil {
		return nil
	}

	if node.Right == nil {
		return node
	}

	return node.Right.findMax()
}

func (node *Node) replaceMeInParentWith(newnode *Node) error {
	if node.Parent == nil {
		return errors.New("no parent node")
	}

	if newnode != nil {
		newnode.Parent = node.Parent
	}

	switch node {
	case node.Parent.Left:
		node.Parent.Left = newnode
	case node.Parent.Right:
		node.Parent.Right = newnode
	}
	return nil
}

func (node *Node) Delete(data interface{}) error {
	if node == nil {
		return errors.New("nil node")
	}

	switch node.Comp(data, node.Value) {
	case -1:
		return node.Left.Delete(data)
	case 1:
		return node.Right.Delete(data)
	}

	// delete this node
	switch {
	case node.Right != nil && node.Left != nil:
		lmaxNode := node.Left.findMax()
		node.Value = lmaxNode.Value
		return lmaxNode.replaceMeInParentWith(lmaxNode.Left)
	case node.Right != nil:
		return node.replaceMeInParentWith(node.Right)
	case node.Left != nil:
		return node.replaceMeInParentWith(node.Left)
	default:
		return node.replaceMeInParentWith(nil)
	}
}

func (tree *Bst) Delete(data interface{}) error {
	if tree.comp == nil {
		return errors.New("comp function not registered")
	}
	if tree.Root == nil {
		return errors.New("failed to delete from empty tree")
	}

	rootParent := &Node{
		Parent: nil,
		Value:  nil,
		Left:   nil,
		Right:  tree.Root,
		Comp:   tree.comp,
	}

	tree.Root.Parent = rootParent
	defer func() {
		tree.Root = rootParent.Right
		tree.Root.Parent = nil
	}()

	return tree.Root.Delete(data)
}

/////////////////////// Traverse /////////////////////
type TraverseFunc func(*Node)

func (node *Node) Traverse(f TraverseFunc) {
	if node == nil {
		return
	}

	if node.Left != nil {
		node.Left.Traverse(f)
	}

	f(node)

	if node.Right != nil {
		node.Right.Traverse(f)
	}
}

func (tree *Bst) Traverse(f TraverseFunc) {
	tree.Root.Traverse(f)
}
