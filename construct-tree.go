package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Integer struct {
	Val *int64
}

type QNode struct {
	Level int
	Data  *TreeNode
	Next  *QNode
}

func string2Slice(str string) []*int {
	d := strings.TrimRight(strings.TrimLeft(str, "["), "]")
	var intList []*int
	for _, v := range strings.Split(d, ",") {
		a := strings.TrimSpace(v)
		var x *int
		if a == "null" {
			x = nil
		} else {
			tp, _ := strconv.ParseInt(a, 10, 64)
			tp1 := int(tp)
			x = &tp1

		}
		intList = append(intList, x)
	}
	return intList
}

func main() {
	b := "[1, null, 2, 2, 32, 31, 3, 23, 1, 23, 123, 12, 3, 12, 31, 23, 2]"

	intList := string2Slice(b)
	fmt.Println("Hello, playground", intList)
	for _, v := range intList {
		if v != nil {
			fmt.Println("Hello, playground", *v)
		} else {
			fmt.Println("Hello, playground", v)
		}

	}
	root := constructTree(intList)
	fmt.Println("Hello, playground", inorderTraversal(root))
	fmt.Println("Hello, playground", levelOrder(root))

	printTreeHorizontal(root)
	fmt.Println()
	fmt.Println("Hello, playground", "------------")
	printTree(root)
}

func constructTree(array []*int) *TreeNode {
	if array == nil || len(array) == 0 || array[0] == nil {
		return nil
	}

	index := 0
	length := len(array)

	root := &TreeNode{
		Val: *array[index],
	}
	queue := list.New()
	queue.PushBack(root)

	var currNode *TreeNode
	for index < length {
		index++
		if index >= length {
			return root
		}
		if queue.Len() >= 1 {
			//var currNode *TreeNode

			xxx := queue.Remove(queue.Front())
			currNode = xxx.(*TreeNode)
		}

		leftChild := array[index]
		if leftChild != nil {
			currNode.Left = &TreeNode{
				Val: *leftChild,
			}
			queue.PushBack(currNode.Left)
		}
		index++
		if index >= length {
			return root
		}
		rightChild := array[index]
		if rightChild != nil {

			currNode.Right = &TreeNode{

				Val: *rightChild}
			queue.PushBack(currNode.Right)

		}
	}

	return root
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	st := make([]*TreeNode, 1531)
	top := -1
	p := root
	for {
		for p != nil {
			top++
			st[top] = p
			p = p.Left
		}
		if top == -1 {
			break
		}
		for top > -1 {
			q := st[top]
			top--
			//fmt.Println(q.Val)
			res = append(res, q.Val)
			if q.Right != nil {
				p = q.Right
				break
			}

		}

	}
	return res
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	var front, rear *QNode
	front = &QNode{
		Level: 0,
		Data:  root,
	}
	rear = front
	flg := -1

	for front != nil {
		if front.Level != flg {
			t := make([]int, 0)
			res = append(res, t)

			flg = front.Level
		}
		res[flg] = append(res[flg], front.Data.Val)
		if front.Data.Left != nil {
			qnode := &QNode{
				Level: front.Level + 1,
				Data:  front.Data.Left,
			}
			rear.Next = qnode
			rear = qnode
		}
		if front.Data.Right != nil {
			qnode := &QNode{
				Level: front.Level + 1,
				Data:  front.Data.Right,
			}
			rear.Next = qnode
			rear = qnode
		}
		front = front.Next
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(getTreeDepth(root.Left), getTreeDepth(root.Right))
}

func traversePreOrder(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var sb bytes.Buffer
	//StringBuilder sb = new StringBuilder();
	sb.WriteString(fmt.Sprint(root.Val))

	pointerRight := "└──"
	var pointerLeft string
	if root.Right != nil {
		pointerLeft = "├──"
	} else {
		pointerLeft = "└──"
	}

	traverseNodes(&sb, "", pointerLeft, root.Left, root.Right != nil)
	traverseNodes(&sb, "", pointerRight, root.Right, false)

	return sb.String()
}

func traverseNodes(sb *bytes.Buffer, padding string, pointer string, node *TreeNode,
	hasRightSibling bool) {
	if node == nil {
		return
	}

	sb.WriteString("\n")
	sb.WriteString(padding)
	sb.WriteString(pointer)
	sb.WriteString(fmt.Sprint(node.Val))

	paddingBuilder := bytes.NewBufferString(padding)
	if hasRightSibling {
		paddingBuilder.WriteString("│  ")
	} else {
		paddingBuilder.WriteString("   ")
	}

	paddingForBoth := paddingBuilder.String()
	pointerRight := "└──"
	// String pointerLeft = (node.right != null) ? "├──" : "└──";
	var pointerLeft string
	if node.Right != nil {
		pointerLeft = "├──"
	} else {
		pointerLeft = "└──"
	}

	traverseNodes(sb, paddingForBoth, pointerLeft, node.Left, node.Right != nil)
	traverseNodes(sb, paddingForBoth, pointerRight, node.Right, false)
}

func printTreeHorizontal(root *TreeNode) {
	print(traversePreOrder(root))
}


func printTree(root *TreeNode) {
	maxLevel := getTreeDepth(root)
	l := list.New()
	l.PushBack(root)
	printNodeInternal(l, 1, maxLevel)
}

func printNodeInternal(nodes *list.List, level int, maxLevel int) {
	if nodes == nil || nodes.Len() == 0 || isAllElementsNull(nodes) {
		return
	}

	floor := maxLevel - level
	endgeLines := int(math.Pow(2, float64((max(floor-1, 0)))))
	firstSpaces := int(math.Pow(2, float64(floor)) - 1)
	betweenSpaces := int(math.Pow(2, float64((floor+1))) - 1)

	printWhitespaces(firstSpaces)

	newNodes := list.New()
	for e := nodes.Front(); e != nil; e = e.Next() {
		node, ok := e.Value.(*TreeNode)
		if ok && node != nil {
			print(node.Val)
			newNodes.PushBack(node.Left)
			newNodes.PushBack(node.Right)
		} else {
			newNodes.PushBack(nil)
			newNodes.PushBack(nil)
			print(" ")
		}
		printWhitespaces(betweenSpaces)
	}

	print("\n")

	for i := 1; i <= endgeLines; i++ {
		for e := nodes.Front(); e != nil; e = e.Next() {
			//for j := 0; j < nodes.Len(); j++ {
			node, ok := e.Value.(*TreeNode)
			printWhitespaces(firstSpaces - i)
			if !ok || node == nil{
				printWhitespaces(endgeLines + endgeLines + i + 1)
				continue
			}
			//printWhitespaces(firstSpaces - i)
			//if node == nil {
			//	printWhitespaces(endgeLines + endgeLines + i + 1)
			//	continue
			//}

			if node.Left != nil {
				print("/")
			} else {
				printWhitespaces(1)
			}

			printWhitespaces(i + i - 1)
			if node.Right != nil {
				print("\\")
			} else {
				printWhitespaces(1)
			}
			printWhitespaces(endgeLines + endgeLines - i)
		}

		print("\n")
	}

	printNodeInternal(newNodes, level+1, maxLevel)
}

func printWhitespaces(count int) {
	for i := 0; i < count; i++ {
		print(" ")
	}
}

func isAllElementsNull(l *list.List) bool {
	for e := l.Front(); e != nil; e = e.Next() {
		node, ok := e.Value.(*TreeNode)
		if ok && node != nil {
			return false
		}
	}
	return true

}

