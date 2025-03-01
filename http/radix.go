package http

import (
	"fmt"
	"strings"
)

type RadixTreeNode struct {
	nextNodes []*RadixTreeNode
	word      string
}

func (node *RadixTreeNode) IsLeaf() bool {
	return len(node.nextNodes) == 0
}

func (node *RadixTreeNode) IsPrefix(word string) bool {
	//fmt.Printf("%+v, %+v, %+v \n", node.word, word, strings.HasPrefix(word, node.word))
	return strings.HasPrefix(word, node.word)
}

func (node *RadixTreeNode) Search(word string) *RadixTreeNode {
	fmt.Printf("searching word %s\n", word)

	if word == "" && node.IsLeaf() {
		return node
	}

	if node.IsLeaf() || word == "" {
		return nil
	}

	for _, nextNode := range node.nextNodes {
		if nextNode.IsPrefix(word) {
			if result := nextNode.Search(strings.Replace(word, nextNode.word, "", 1)); result != nil {
				return result
			}
		}
	}

	return nil
}

func (node *RadixTreeNode) Insert(word string) {

	for _, nextNode := range node.nextNodes {
		if nextNode.IsPrefix(word) {
			//return nextNode.Search(strings.Replace(word, node.word, "", 1))
			wordWithoutPrefix := strings.Replace(word, nextNode.word, "", 1)
			nextNode.Insert(wordWithoutPrefix)
			return
		}
	}

	node.nextNodes = append(node.nextNodes, &RadixTreeNode{word: word, nextNodes: []*RadixTreeNode{}})
}

func NewRadixTree() *RadixTreeNode {
	return &RadixTreeNode{
		word:      "",
		nextNodes: []*RadixTreeNode{},
	}
}
