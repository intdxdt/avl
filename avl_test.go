package avl

import (
	"fmt"
	"testing"
	"github.com/intdxdt/cmp"
	"github.com/franela/goblin"
	"github.com/intdxdt/skiplist"
)

func printVal(n interface{}) string {
	return fmt.Sprintf("%v", n)
}

func TestAVL(t *testing.T) {
	g := goblin.Goblin(t)

	var array = []float64{
		1, 0, 2, 3, 4, 3, 3.3, 9., 29., 3.1, 0.1, 1.1,
		1.81, 0.91, 0.81, 0.71, 0.88, 0.82, 0.81,
	}

	var uniq_array = make([]float64, 0)
	var set = skiplist.NewSkipList(len(array), false, cmp.F64)
	for i := range array {
		set.Insert(array[i])
	}

	//uniq array
	set.Each(func(o interface{}, i int) {
		uniq_array = append(uniq_array, o.(float64))
	})

	g.Describe("AVL - Balanced Binary Search Tree", func() {
		var tree = NewAVL(cmp.F64)
		g.Assert(tree != nil).IsTrue()

		tree.Insert(0.0)
		g.Assert(tree.Height()).Equal(0)
		for i := range array {
			tree.Insert(array[i])
		}

		g.Assert(tree.Search(1.81).Key).Equal(1.81)
		g.Assert(tree.Contains(1.81)).IsTrue()
		g.Assert(tree.Contains(1.817)).IsFalse()
		g.Assert(tree.Contains(0.17)).IsFalse()
		g.Assert(tree.Contains(0.71)).IsTrue()

		treeArray := tree.ToArray()

		g.Assert(len(treeArray)).Eql(len(uniq_array))
		for i := range treeArray {
			g.Assert(treeArray[i].(float64)).Equal(uniq_array[i])
		}

		var trvArray = make([]float64, 0)
		var itemArray = make([]float64, 0)
		tree.Traverse(func(n interface{}) bool {
			trvArray = append(trvArray, n.(float64))
			return true
		})

		tree.EachItem(func(n interface{}) bool {
			itemArray = append(itemArray, n.(float64))
			return true
		})

		g.Assert(uniq_array).Eql(trvArray)
		g.Assert(uniq_array).Eql(itemArray)

		g.Assert(tree.BST.Root.Height).Equal(4)
		g.Assert(tree.BST.Root.Key).Equal(float64(3))

		fmt.Println(tree.Print(printVal))
		fmt.Println("\n\n")

		var rm, parent = tree.Remove(0.81)
		g.Assert(rm.Key).Equal(0.81)
		//parent of the succeeding node
		g.Assert(parent.Key).Equal(0.1)

		rm, parent = tree.Remove(0.0)
		rm, parent = tree.Remove(3.0)
		rm, parent = tree.Remove(3.3)
		rm, parent = tree.Remove(9.0)

		g.Assert(tree.BST.Root.Key).Equal(1.)
		g.Assert(tree.BST.Root.Height).Equal(3)
		g.Assert(tree.BST.Root.Left.Key).Equal(0.71)
		g.Assert(tree.BST.Root.Left.Height).Equal(2)
		g.Assert(tree.BST.Root.Right.Height).Equal(2)

		g.Assert(tree.Size()).Equal(12)

		fmt.Println(tree.Print(printVal))

		fmt.Println("\nWorse Case -- BST as List - AVL as BBST")

		g.Assert(tree.Empty().BST.Root == nil).IsTrue()
		g.Assert(tree.First() == nil).IsTrue()
		g.Assert(tree.Last() == nil).IsTrue()
		g.Assert(tree.NextItem(nil) == nil).IsTrue()
		g.Assert(tree.PrevItem(nil) == nil).IsTrue()

		tree.Insert(1.1)
		tree.Insert(1.3)
		tree.Insert(1.5)
		tree.Insert(1.9)
		tree.Insert(2.1)
		tree.Insert(2.5)
		tree.Insert(2.7)
		var printVal = tree.Print(func(n interface{}) string {
			return fmt.Sprintf("%v", n)
		})
		fmt.Println(printVal)

		g.Assert(tree.BST.Root.Key).Equal(1.9)
		g.Assert(tree.First()).Equal(1.1)
		g.Assert(tree.Last()).Equal(2.7)

		g.Assert(tree.PrevItem(1.9)).Eql(float64(1.5))
		g.Assert(tree.NextItem(1.9)).Eql(float64(2.1))

		g.Assert(tree.PrevItem(1.) == nil).IsTrue()
		g.Assert(tree.NextItem(1.) == nil).IsTrue()

		g.Assert(tree.PrevItem(1.1) == nil).IsTrue()
		g.Assert(tree.NextItem(1.1)).Eql(float64(1.3))

		g.Assert(tree.PrevItem(2.7)).Eql(float64(2.5))
		g.Assert(tree.NextItem(2.7) == nil).IsTrue()

	})
}

func TestAVL_Set(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("AVL - Set Opt", func() {
		var tree1 = NewAVL(cmp.Int)
		var tree2 = NewAVL(cmp.Int)
		var tree3 = NewAVL(cmp.Int)

		tree1.Insert(1)
		tree1.Insert(0)
		tree1.Insert(5)
		tree1.Insert(4)
		tree1.Insert(7)
		tree1.Insert(10)
		tree1.Insert(13)

		//===
		tree2.Insert(4)
		tree2.Insert(7)
		tree2.Insert(9)
		tree2.Insert(10)
		tree2.Insert(20)
		tree2.Insert(17)
		tree2.Insert(91)
		//===
		tree3.Insert(21)
		tree3.Insert(11)
		tree3.Insert(12)
		tree3.Insert(41)
		tree3.Insert(92)

		utree := tree1.Union(tree2)
		intertree := tree1.Intersection(tree2)
		d1tree := tree1.Difference(tree2)
		d2tree := tree2.Difference(tree1)
		symtree := tree2.SymDifference(tree1)
		uset := make([]interface{}, 0)
		for _, v := range []int{0, 1, 4, 5, 7, 9, 10, 13, 17, 20, 91} {
			uset = append(uset, v)
		}
		g.Assert(utree.ToArray()).Eql(uset)
		g.Assert(utree.Clone().ToArray()).Eql(uset)
		g.Assert(intertree.ToArray()).Eql([]interface{}{4, 7, 10})

		g.Assert(d1tree.ToArray()).Eql([]interface{}{0, 1, 5, 13})
		g.Assert(d2tree.ToArray()).Eql([]interface{}{9, 17, 20, 91})
		g.Assert(symtree.ToArray()).Eql([]interface{}{0, 1, 5, 9, 13, 17, 20, 91})
		g.Assert(len(tree1.String()) > 0).IsTrue()
		//fmt.Println(tree1.String(), "\n")
		//fmt.Println(tree2.String(), "\n")
	})
}

func TestSetAVL(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST - Binary Search Tree - Intersection", func() {
		var tree1 = NewAVL(cmp.Int)
		var tree2 = NewAVL(cmp.Int)
		var tree3 = NewAVL(cmp.Int)
		var tree4 = NewAVL(cmp.Int)
		var tree5 = NewAVL(cmp.Int)

		tree1.Insert(5)
		tree1.Insert(1)
		tree1.Insert(10)
		tree1.Insert(0)
		tree1.Insert(4)
		tree1.Insert(7)
		tree1.Insert(9)

		//===
		tree2.Insert(10)
		tree2.Insert(7)
		tree2.Insert(20)
		tree2.Insert(4)
		tree2.Insert(91)
		//===
		tree3.Insert(11)
		tree3.Insert(12)
		tree3.Insert(21)
		tree3.Insert(41)
		tree3.Insert(92)

		var empty = make([]interface{}, 0)

		g.Assert(tree1.Intersection(tree2).ToArray()).Eql(
			[]interface{}{4, 7, 10},
		)
		g.Assert(tree2.Intersection(tree1).ToArray()).Eql(
			[]interface{}{4, 7, 10},
		)
		g.Assert(tree3.Intersection(tree1).ToArray()).Eql(empty)
		g.Assert(tree2.Intersection(tree3).ToArray()).Eql(empty)

		diff1 := []interface{}{0, 1, 5, 9}
		diff2 := []interface{}{20, 91}
		sdiff := []interface{}{0, 1, 5, 9, 20, 91}
		g.Assert(tree1.Difference(tree2).ToArray()).Eql(diff1)
		g.Assert(tree2.Difference(tree1).ToArray()).Eql(diff2)
		g.Assert(tree1.SymDifference(tree2).ToArray()).Eql(sdiff)

		g.Assert(tree1.Difference(tree1).ToArray()).Eql(empty)
		g.Assert(tree2.Difference(tree2).ToArray()).Eql(empty)

		//merge two trees
		mged := []interface{}{
			0, 1, 4, 5, 7,
			9, 10, 20, 91}

		mgd := tree1.Union(tree2)
		fmt.Println("merged->\n", mgd)
		g.Assert(mgd.ToArray()).Eql(mged)

		mgd2 := tree1.Union(tree2)
		g.Assert(mgd.ToArray()).Eql(mgd2.ToArray())

		//union with an empty tree
		mgd4 := tree4.Union(tree1)

		//clone an empty tree
		mgd5 := tree5.Clone()
		tree1.EachItem(func(o interface{}) bool {
			mgd5.Insert(o)
			return true
		})
		g.Assert(mgd5.ToArray()).Eql(mgd4.ToArray())

		fmt.Println("\ntree-----1:\n", tree1)
		fmt.Println("\ntree-----2:\n", tree2)
		fmt.Println("\ntree union:\n", mgd2)
	})
}

func TestSetAVL2(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("BST - Binary Search Tree 2 - difference ", func() {
		var tree1 = NewAVL(cmp.Int)
		var tree2 = NewAVL(cmp.Int)

		tree1.Insert(53)
		tree1.Insert(33)
		tree1.Insert(63)
		tree1.Insert(03)
		tree1.Insert(43)
		tree1.Insert(73)
		tree1.Insert(93)
		//===
		tree2.Insert(13)
		tree2.Insert(71)
		tree2.Insert(20)
		tree2.Insert(42)
		tree2.Insert(91)

		var empty = make([]interface{}, 0)

		g.Assert(tree1.Intersection(tree2).ToArray()).Eql(empty)
		g.Assert(tree2.Intersection(tree1).ToArray()).Eql(empty)

		g.Assert(tree1.Difference(tree2).ToArray()).Eql(tree1.ToArray())
		g.Assert(tree2.Difference(tree1).ToArray()).Eql(tree2.ToArray())

		xor12 := []interface{}{
			03, 13, 20, 33, 42, 43,
			53, 63, 71, 73, 91, 93}
		g.Assert(tree1.SymDifference(tree2).ToArray()).Eql(xor12)
	})
}
