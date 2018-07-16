package avl

import "github.com/intdxdt/bst"

/*
compute tree height
 *  hack for avl -1 if nil else node.Height
 param node
 returns {Number}
 */
func (avl *AVL) height(node *bst.Node) int {
	if bst.IsNil(node) {
		return -1
	}
	return node.Height
}

// update height
func (avl *AVL) updateHeight(node *bst.Node) {
	var l = avl.height(node.Left)
	var r = avl.height(node.Right)
	if l > r {
		node.Height = l + 1
	} else {
		node.Height = r + 1
	}
}

//Rebalance BST Tree
func (avl *AVL) rebalance(node *bst.Node) {
	left := bst.NewBranch().AsLeft()
	right := bst.NewBranch().AsRight()

	for node != nil {
		avl.updateHeight(node)
		var balancefactor = avl.height(node.Right) - avl.height(node.Left)
		//left heavy  --> rotate right
		if -(balancefactor) > 1 {
			avl.rotateHeavybranch(node, right)
		}
		//right heavy --> rotate left
		if balancefactor > 1 {
			avl.rotateHeavybranch(node, left)
		}
		node = node.Parent
	}
}

/*
 rotate heavy branch to new branch
 */
func (avl *AVL) rotateHeavybranch(node *bst.Node, tobranch *bst.Branch) {
	//get heavy branch
	var branch = tobranch.ConjBranch()
	var n = node.GetNode(branch)
	nTobr, nBr := n.GetNode(tobranch), n.GetNode(branch)
	if avl.height(nTobr) > avl.height(nBr) {
		avl.rotate(n, branch)
	}
	avl.rotate(node, tobranch)
}

//rotate node to a branch
func (avl *AVL) rotate(node *bst.Node, tobranch *bst.Branch) {
	//opposite tobranch to rotation => opposite heavy branch
	var conjbranch = tobranch.ConjBranch()

	//heavychild in opposite tobranch of rotate
	var heavychild = node.GetNode(conjbranch)

	//node.Parent[node_branch] child <--> heavychild
	bst.Ptr(node.Parent, heavychild, node.Branch())

	//node[conjbranch] child  <--> heavychild[tobranch]
	bst.Ptr(node, heavychild.GetNode(tobranch), conjbranch)

	//heavychild[tobranch] <--> node[tobranch]
	bst.Ptr(heavychild, node, tobranch)

	//if heavy child parent is nil :  update root
	if bst.IsNil(heavychild.Parent) {
		avl.BST.Root = heavychild
	}

	//update heights
	avl.updateHeight(node)
	avl.updateHeight(heavychild)
}
