package Trie

// 前缀树 是一种树形结构
// 前缀树的定义
/**
1.根节点不包含字符，除根节点外每一个节点都只包含一个字符。
2.从根节点到某一节点，路径上经过的字符连接起来，为该节点对应的字符串。
3.每个节点的所有子节点包含的字符都不相同。
*/

type Trie struct {
	// 表示以当前字符结尾的字符串的数量
	Count int
	// 表示以当前路径组合而成的字符串为前缀的字符串的数量
	Prefix int
	Nexts  []*Trie
}

func Constructor() Trie {
	// 因为总共有26个字母 之所以节点中不包含值的原因是可以用过Next就可以得到节点的值
	return Trie{Nexts: make([]*Trie, 26)}
}

func (this *Trie) Insert(word string) {
	p := this
	for i := 0; i < len(word); i++ {
		nextIndex := int(word[i] - 'a')
		if p.Nexts[nextIndex] == nil {
			p.Nexts[nextIndex] = &Trie{Nexts: make([]*Trie, 26)}
		}
		p = p.Nexts[nextIndex]
		// 经过该字符就需要加1 因为这个是前缀
		p.Prefix++
		if i == len(word)-1 {
			// 以该单词结尾
			p.Count++
		}
	}
}

func (this *Trie) Search(word string) bool {
	p := this
	for i := 0; i < len(word); i++ {
		nextIndex := int(word[i] - 'a')
		if p.Nexts[nextIndex] == nil {
			return false
		}
		p = p.Nexts[nextIndex]
	}
	return p.Count > 0
}

func (this *Trie) StartsWith(prefix string) bool {
	p := this
	for i := 0; i < len(prefix); i++ {
		nextIndex := int(prefix[i] - 'a')
		if p.Nexts[nextIndex] == nil {
			return false
		}
		p = p.Nexts[nextIndex]
	}
	return p.Prefix > 0
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
