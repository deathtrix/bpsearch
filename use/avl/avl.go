package avl

// Node - node type
type Node struct {
	Key            string
	Value          interface{}
	Height         int
	Lchild, Rchild *Node
}

func leftRotate(root *Node) *Node {
	node := root.Rchild
	// fmt.Println(node.Key)
	root.Rchild = node.Lchild
	node.Lchild = root

	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	node.Height = max(height(node.Rchild), height(node.Lchild)) + 1
	return node
}

func leftRigthRotate(root *Node) *Node {
	root.Lchild = leftRotate(root.Lchild)
	root = rightRotate(root)
	return root
}

func rightRotate(root *Node) *Node {
	node := root.Lchild
	root.Lchild = node.Rchild
	node.Rchild = root
	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	node.Height = max(height(node.Lchild), height(node.Rchild)) + 1
	return node
}

func rightLeftRotate(root *Node) *Node {
	root.Rchild = rightRotate(root.Rchild)
	root = leftRotate(root)
	return root
}

func height(root *Node) int {
	if root != nil {
		return root.Height
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Insert - insert a new node
func Insert(root *Node, key string) *Node {
	if root == nil {
		root = &Node{key, nil, 0, nil, nil}
		root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
		return root
	}

	if key < root.Key {
		root.Lchild = Insert(root.Lchild, key)
		if height(root.Lchild)-height(root.Rchild) == 2 {
			if key < root.Lchild.Key {
				root = rightRotate(root)
			} else {
				root = leftRigthRotate(root)
			}
		}
	}

	if key > root.Key {
		root.Rchild = Insert(root.Rchild, key)
		if height(root.Rchild)-height(root.Lchild) == 2 {
			if key > root.Rchild.Key {
				root = leftRotate(root)
			} else {
				root = rightLeftRotate(root)
			}
		}
	}

	root.Height = max(height(root.Lchild), height(root.Rchild)) + 1
	return root
}

type action func(node *Node)

// InOrder - parse tree in order
func InOrder(root *Node, action action) {
	if root == nil {
		return
	}

	InOrder(root.Lchild, action)
	action(root)
	InOrder(root.Rchild, action)
}

// Find - find a key in tree
func Find(root *Node, key string) *Node {
	if root == nil {
		return nil
	}

	if root.Key == key {
		return root
	}
	if key < root.Key {
		return Find(root.Lchild, key)
	}
	return Find(root.Rchild, key)
}

/*
Usage:

func main() {
	var root *Node
	keys := []int{2, 6, 1, 3, 5, 7, 16, 15, 14, 13, 12, 11, 8, 9, 10}
	for _, key := range keys {
		root = Insert(root, key)
	}

	InOrder(root, func(node *Node) {
		fmt.Println(node.Key, node.Height)
	})
}
*/
