package main

type WordDictionary struct {
	root *Node
}

type Node struct {
	children [26]*Node
	end      bool
}

func Constructor() WordDictionary {
	return WordDictionary{root: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this.root

	for _, w := range word {
		ix := w - 'a'
		if curr.children[ix] == nil {
			curr.children[ix] = &Node{}
		}
		curr = curr.children[ix]
	}
	curr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	curr := this.root
	return curr.Search(word)
}

func (this *Node) Search(word string) bool {
	curr := this
	for i, w := range word {
		if w == '.' {
			for _, v := range curr.children {
				if v != nil && v.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else if curr.children[w-'a'] == nil {
			return false
		}
		curr = curr.children[w-'a']
	}
	return curr.end
}
