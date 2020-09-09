package _138_copy_list_with_random_pointer

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var nodeMap = make(map[*Node]*Node, 1)
	var node *Node
	var nodeCopy *Node

	node = head
	for node != nil {
		nodeCopy = new(Node)
		nodeCopy.Val = node.Val

		nodeMap[node] = nodeCopy
		node = node.Next
	}

	node = head
	for node != nil {
		nodeMap[node].Next = nodeMap[node.Next]
		nodeMap[node].Random = nodeMap[node.Random]
		node = node.Next
	}

	return nodeMap[head]

}