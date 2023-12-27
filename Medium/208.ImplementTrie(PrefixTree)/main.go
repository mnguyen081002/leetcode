package main

type Node struct {
	Children [26]*Node
	End      bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{}}
}

func (this *Trie) Insert(word string) {
	current := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if current.Children[idx] == nil {
			current.Children[idx] = &Node{}
		}
		current = current.Children[idx]
	}
	current.End = true
}

func (this *Trie) Search(word string) bool {
	current := this.root
	for _, ch := range word {
		idx := ch - 'a'
		if current.Children[idx] == nil {
			return false
		}
		current = current.Children[idx]
	}
	return current.End
}

func (this *Trie) StartsWith(prefix string) bool {
	current := this.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if current.Children[idx] == nil {
			return false
		}
		current = current.Children[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
