package main

type Node struct {
	Children map[rune]*Node
	End      bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: &Node{
		Children: map[rune]*Node{},
	}}
}

func (this *Trie) Insert(word string) {
	cur := this.root

	for _, v := range word {
		if _, found := cur.Children[v]; !found {
			cur.Children[v] = &Node{
				Children: map[rune]*Node{},
			}
		}
		cur = cur.Children[v]
	}
	cur.End = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root

	for _, v := range word {
		if _, found := cur.Children[v]; !found {
			return false
		}
		cur = cur.Children[v]
	}
	return cur.End
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root

	for _, v := range prefix {
		if _, found := cur.Children[v]; !found {
			return false
		}
		cur = cur.Children[v]
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
