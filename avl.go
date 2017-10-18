//Package avl is a balanced binary search tree.
package avl

import (
	"github.com/intdxdt/bst"
	"github.com/intdxdt/cmp"
)

//AVL Type
type AVL struct {
	BST *bst.BST
}

//NewAVL- Creates a New AVL Tree
func NewAVL(comparator cmp.Compare) *AVL {
	return &AVL{BST: bst.NewBST(comparator)}
}

//Clone AVL
func (avl *AVL) Clone() *AVL {
	return &AVL{BST: avl.BST.Clone()}
}

//Search item
func (avl *AVL) Search(val interface{}) *bst.Node {
	return bst.SearchItem(avl.BST.Root, val)
}

//Insert into AVL
func (avl *AVL) Insert(item interface{}) *bst.Node {
	var node = avl.BST.Insert(item)
	avl.rebalance(node)
	return node
}

//Remove node with value at given key
func (avl *AVL) Remove(val interface{}) (*bst.Node, *bst.Node) {
	node, successorParent := avl.BST.Remove(val)
	avl.rebalance(successorParent) //rebalance successor_parent
	return node, successorParent   //return deleted node, succ_parent
}

//Empty AVL Tree
func (avl *AVL) Empty() *AVL {
	avl.BST.Empty()
	return avl
}

//Traverse each node in AVL
func (avl *AVL) Traverse(fn func(interface{}) bool) {
	avl.BST.Traverse(func(n *bst.Node) bool {
		fn(n.Key)
		return true
	})
}

//EachItem - Iterates over each item in the AVL
func (avl *AVL) EachItem(fn func(interface{}) bool) {
	avl.BST.EachItem(fn)
}

//Size of bst node
func (avl *AVL) Size() int {
	return avl.BST.Size()
}

//Height- computes height of bst tree
func (avl *AVL) Height() int {
	return avl.BST.Height()
}

//ToArray - tree as array
func (avl *AVL) ToArray() []interface{} {
	return avl.BST.ToArray()
}

//First item in the Tree
func (avl *AVL) First() interface{} {
	return avl.BST.First()
}

//Last item in the Tree
func (avl *AVL) Last() interface{} {
	return avl.BST.Last()
}

//NextItem - next item given item in the Tree
func (avl *AVL) NextItem(v interface{}) interface{} {
	return avl.BST.NextItem(v)
}

//PrevItem - previous given item in the Tree
func (avl *AVL) PrevItem(v interface{}) interface{} {
	return avl.BST.PrevItem(v)
}

//Print tree structure as string
func (avl *AVL) Print(keyfn func(interface{}) string) string {
	return avl.BST.Print(keyfn)
}

//String - avl as string
func (avl *AVL) String() string {
	return avl.BST.String()
}

//Union of two trees
func (avl *AVL) Union(other *AVL) *AVL {
	var tree = NewAVL(avl.BST.Cmp)
	for _, o := range avl.BST.Union(other.BST) {
		tree.Insert(o)
	}
	return tree
}

//Intersection of two trees
func (avl *AVL) Intersection(other *AVL) *AVL {
	tree := NewAVL(avl.BST.Cmp)
	vals := avl.BST.Intersection(other.BST)
	for _, v := range vals {
		tree.Insert(v)
	}
	return tree
}

//Difference - intersection of two trees
func (avl *AVL) Difference(other *AVL) *AVL {
	tree := NewAVL(avl.BST.Cmp)
	diff := avl.BST.Difference(other.BST)
	for _, v := range diff {
		tree.Insert(v)
	}
	return tree
}

//SymDifference - intersection of two trees
func (avl *AVL) SymDifference(other *AVL) *AVL {
	tree := NewAVL(avl.BST.Cmp)
	diff := avl.BST.XOR(other.BST)
	for _, v := range diff {
		tree.Insert(v)
	}
	return tree
}
