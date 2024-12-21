package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{},
	}
}
func (this *Trie) Insert(word string) {
	node := this.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}
func (this *Trie) Search(word string) bool {
	node := this.root
	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return node.isEnd
}
func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
	for _, ch := range prefix {
		index := ch - 'a'
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // true
	fmt.Println(trie.Search("app"))     // false
	fmt.Println(trie.StartsWith("app")) // true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // true
}
