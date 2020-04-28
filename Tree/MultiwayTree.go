package tree

//Tree data structure is encapsulated within this struct
type Tree struct {
	rootNode *treeNode
	size     int
}

type treeNode struct {
	Data     int
	children []*treeNode
	parent   *treeNode
}

func (node treeNode) getChildren() []*treeNode {
	return node.children
}

//finds all treeNodes having their data field same as the argument data using breadth-first algorithm
func getNodes(tree Tree, data int) []*treeNode {
	var matchingNodes []*treeNode
	var queue = []*treeNode{tree.rootNode}

	for _, node := range queue {
		if node.Data == data {
			matchingNodes = append(matchingNodes, node)
			queue = append(queue, node.children...)
		} else {
			queue = append(queue, node.children...)
		}
	}
	return matchingNodes
}

func (node treeNode) getNumberOfChildren() int {
	return len(node.children)
}

func (node treeNode) getParent() *treeNode {
	if node.parent != nil {
		return node.parent
	}
	return &node
}

//returns the sibling which contains the data passed, returns a slice of sibling if more than one sibling have data field = data
func (node treeNode) getSibling(data int) []*treeNode {
	var siblingsWithSameData []*treeNode
	var allSiblings []*treeNode
	allSiblings = append(allSiblings, node.getSiblings()...)
	for _, node := range allSiblings {
		if node.Data == data {
			siblingsWithSameData = append(siblingsWithSameData, node)
		}
	}
	return siblingsWithSameData
}

//returns a slice of sibling of node excluding node itself
func (node treeNode) getSiblings() []*treeNode {
	var siblings []*treeNode
	siblings = append(siblings, node.parent.getChildren()...)
	return siblings
}

//finds out if any descendant of node has their data field equal to data, returns true on first match
//implemented using the depth-first-search
func (node *treeNode) isAncestorOf(data int) bool {
	var queue = []*treeNode{node}
	for _, eachNode := range queue {
		if eachNode.Data == data {
			return true
		}
		queue = append(queue, eachNode.children...)
	}
	return false
}

//finds out if otherNode is a descendant of node
//implemented using the depth-first-search
func (node *treeNode) isAncestorOfNode(possibleChildNode *treeNode) bool {
	var queue = []*treeNode{node}
	for _, eachNode := range queue {
		if eachNode == possibleChildNode {
			return true
		}
		queue = append(queue, eachNode.children...)
	}
	return false
}

func (node treeNode) isChildrenContains(data int) bool {
	for _, value := range node.children {
		if data == value.Data {
			return true
		}
	}
	return false
}

func (node *treeNode) isDescendantOfNode(ancestorNode *treeNode) bool {
	return ancestorNode.isAncestorOfNode(node)
}

func (node treeNode) IsEmpty() bool {
	return len(node.children) == 0
}

func (node treeNode) isParentOf(data int) bool {
	return node.isChildrenContains(data)
}

func (node treeNode) isSiblingOf(data int) bool {
	return node.parent.isChildrenContains(data)
}

//**************************Tree methods************************************

//IsEmpty checks if the receiver tree is empty
func (tree Tree) IsEmpty() bool {
	return tree.size == 0 || tree.rootNode.Data == 0
}

func (tree Tree) IsContains(node *treeNode) bool {

}

//GetPath returns the path traversing from an ancestor node to a descendant node
//Returns an error if one of the nodes are not found on the tree
func (tree Tree) GetPath(fromThisNode treeNode, toThisNode treeNode) (*Tree, error) {

}

//Add a new treeNode to this tree
func (tree *Tree) Add(newNode *treeNode) {

}

//Remove an existing node from this tree, returns error if the node is not found on the tree
func (tree *Tree) Remove(existingNode *treeNode) error {

}
