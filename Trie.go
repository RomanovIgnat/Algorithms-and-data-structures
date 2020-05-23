type Node struct {
    Val rune
    Next []*Node
    End bool
}


type Trie struct {
    Root *Node
}


func Constructor() Trie {
    trie := new(Trie)
    node := new(Node)
    trie.Root = node
    return *trie
}


func (this *Trie) Insert(word string)  {
    cur_node := this.Root
    
    LOOP: for _, sym := range word {
        for idx := 0; idx < len(cur_node.Next); idx++ {
            if (cur_node.Next[idx]).Val == sym {
                cur_node = cur_node.Next[idx]
                continue LOOP
            }
        }
        new_node := new(Node)
        new_node.Val = sym
        cur_node.Next = append(cur_node.Next, new_node)
        cur_node = new_node
    }
    
    cur_node.End = true
}


func (this *Trie) Search(word string) bool {
    cur_node := this.Root
    
    LOOP: for _, sym := range word {
        for idx := 0; idx < len(cur_node.Next); idx++ {
            if (cur_node.Next[idx]).Val == sym {
                cur_node = cur_node.Next[idx]
                continue LOOP
            }
        }
        return false
    }
    
    if cur_node.End == true {
        return true
    }
    return false
}


func (this *Trie) StartsWith(prefix string) bool {
    cur_node := this.Root
    
    LOOP: for _, sym := range prefix {
        for idx := 0; idx < len(cur_node.Next); idx++ {
            if (cur_node.Next[idx]).Val == sym {
                cur_node = cur_node.Next[idx]
                continue LOOP
            }
        }
        return false
    }
    
    return true
}
