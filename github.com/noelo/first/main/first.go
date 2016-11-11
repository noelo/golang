package main

import "fmt"

// Trie holds the tree
type Trie struct {
	children map[string]Trie
}

func init() {
	fmt.Println("In init ")
}

func trieFactory() *Trie {
	res := new(Trie)
	res.children = make(map[string]Trie)
	return res
}

func dumpTrie(inTree *Trie, level int) {
	currTrie := inTree
	for key, value := range currTrie.children {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		fmt.Println(key)
		dumpTrie(&value, level+1)
	}
}

func addWordToTrie(inTree *Trie, inWord string) {
	curr := inTree
	for _, element := range inWord {
		// fmt.Println("Curr ==>", curr)
		if _, ok := curr.children[string(element)]; ok == false {
			node := trieFactory()
			// fmt.Println("Node ==>", &node)

			(*curr).children[string(element)] = *node
			curr = node
		} else {
			// fmt.Printf("Element %s already exists\n", string(element))
			child := curr.children[string(element)]
			curr = &child
		}
	}
}

func searchForWord(inTree *Trie, inWord string) bool {
	if len(inWord) == 0 {
		return true
	}
	if inTree == nil {
		return false
	}

	next, ok := inTree.children[string(inWord[0])]
	if ok == false {
		return false
	}
	return searchForWord(&next, inWord[1:len(inWord)])
}

func main() {
	root := trieFactory()
	addWordToTrie(root, "test")
	addWordToTrie(root, "ted")
	addWordToTrie(root, "apple")
	addWordToTrie(root, "android")
	addWordToTrie(root, "teddy")
	addWordToTrie(root, "zebra")
	addWordToTrie(root, "zoolander")
	dumpTrie(root, 0)
	fmt.Printf("Found %t \n", searchForWord(root, "test"))
	fmt.Printf("Found %t \n", searchForWord(root, "ted"))
	fmt.Printf("Found %t \n", searchForWord(root, "error"))
	fmt.Printf("Found %t \n", searchForWord(root, "zoolander"))
	fmt.Printf("Found %t \n", searchForWord(root, "flakey"))
	fmt.Printf("Found %t \n", searchForWord(root, ""))
	fmt.Printf("Found %t \n", searchForWord(nil, ""))

}
