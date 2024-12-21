package main

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}
type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}
func (this *WordDictionary) AddWord(word string) {
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
func (this *WordDictionary) Search(word string) bool {
	return this.searchRecursive(word, 0, this.root)
}

func (this *WordDictionary) searchRecursive(word string, index int, node *TrieNode) bool {
	if index == len(word) {
		return node.isEnd
	}
	char := word[index]
	if char == '.' {
		for _, child := range node.children {
			if child != nil && this.searchRecursive(word, index+1, child) {
				return true
			}
		}
		return false
	}
	ind := char - 'a'
	if node.children[ind] == nil {
		return false
	}
	return this.searchRecursive(word, index+1, node.children[ind])

}

func main() {
	// Example Usage
	wordDict := Constructor()
	wordDict.AddWord("bad")
	wordDict.AddWord("dad")
	wordDict.AddWord("mad")

	// Search for exact matches and with wildcards
	fmt.Println(wordDict.Search("pad")) // false
	fmt.Println(wordDict.Search("bad")) // true
	fmt.Println(wordDict.Search(".ad")) // true
	fmt.Println(wordDict.Search("b..")) // true
}
