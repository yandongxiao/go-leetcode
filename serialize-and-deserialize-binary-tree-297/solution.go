package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	queue := []*TreeNode{root}
	var children []*TreeNode
	var result strings.Builder
	result.WriteString("[")
	for len(queue) > 0 {
		for _, node := range queue {
			if node == nil {
				result.WriteString("null,")
			} else {
				result.WriteString(strconv.Itoa(node.Val) + ",")
				children = append(children, node.Left)
				children = append(children, node.Right)
			}
		}
		queue = children
		children = nil
	}
	result.WriteString("]")
	return result.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1]
	strs := strings.Split(data, ",")
	root := &TreeNode{
		Val: toInt(strs[0]),
	}
	strs = strs[1:]
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if len(strs) >= 2 {
			if strs[0] != "null" {
				node.Left = &TreeNode{
					Val: toInt(strs[0]),
				}
				queue = append(queue, node.Left)
			}
			if strs[1] != "null" {
				node.Right = &TreeNode{
					Val: toInt(strs[1]),
				}
				queue = append(queue, node.Right)
			}
			strs = strs[2:]
		} else if len(strs) == 1 {
			if strs[0] != "null" {
				node.Left = &TreeNode{
					Val: toInt(strs[0]),
				}
				queue = append(queue, node.Left)
			}
			strs = strs[1:]
		}
	}
	return root
}

func toInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
