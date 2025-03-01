package http

import (
	"testing"
)

func TestRadixTreeNode_IsPrefix(t *testing.T) {
	radix := &RadixTreeNode{
		word:      "foo",
		nextNodes: []*RadixTreeNode{},
	}

	if result := radix.IsPrefix("foobar"); result != true {
		t.Fatalf("expected true but received %+v", result)
	}
}

func TestRadixTreeNode_Search(t *testing.T) {
	radix := &RadixTreeNode{
		word: "",
		nextNodes: []*RadixTreeNode{
			{
				word: "foo",
				nextNodes: []*RadixTreeNode{
					{word: "bar", nextNodes: []*RadixTreeNode{}},
					{word: "baz", nextNodes: []*RadixTreeNode{
						{word: "hello", nextNodes: []*RadixTreeNode{}},
					}},
				},
			},
		},
	}

	if result := radix.Search("foobar"); result == nil || result.word != "bar" {
		t.Fatalf("expected to receive the foobar node %+v", result)
	}
}

func TestRadixTreeNode_SearchLeaf(t *testing.T) {
	radix := &RadixTreeNode{
		word: "",
		nextNodes: []*RadixTreeNode{
			{
				word: "foo",
				nextNodes: []*RadixTreeNode{
					{word: "bar", nextNodes: []*RadixTreeNode{}},
					{word: "baz", nextNodes: []*RadixTreeNode{
						{word: "hello", nextNodes: []*RadixTreeNode{}},
					}},
				},
			},
		},
	}

	if result := radix.Search("foobazhello"); result == nil || result.word != "hello" {
		t.Fatalf("expected to receive the hello node %+v", result)
	}
}

func TestRadixTreeNode_SearchInserted(t *testing.T) {
	radix := &RadixTreeNode{
		word: "",
		nextNodes: []*RadixTreeNode{
			{
				word: "he",
				nextNodes: []*RadixTreeNode{
					{word: "llo", nextNodes: []*RadixTreeNode{
						{word: "_world", nextNodes: []*RadixTreeNode{}},
					}},
				},
			},
		},
	}

	if result := radix.Search("hello_world"); result == nil || result.word != "_world" {
		t.Fatalf("expected to receive the _world node %+v", result)
	}
}

func TestRadixTreeNode_Insert(t *testing.T) {
	radix := NewRadixTree()

	radix.Insert("he")
	radix.Insert("hello")
	radix.Insert("hello_world")
	if result := radix.Search("hello_world"); result == nil || result.word != "_world" {
		t.Fatalf("expected to receive the _world node %+v", result)
	}
}
