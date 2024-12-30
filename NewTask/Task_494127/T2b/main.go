package main  
import "fmt"

// TrieNode represents a node in the trie data structure
type TrieNode struct {
    children map[string]*TrieNode
    isValue  bool
    value    int
}

func NewTrieNode() *TrieNode {
    return &TrieNode{children: make(map[string]*TrieNode)}
}

// Trie represents the trie data structure
type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: NewTrieNode()}
}

// Put inserts a key-value pair into the trie
func (t *Trie) Put(keys []string, value int) {
    current := t.root
    for _, key := range keys {
        if current.children[key] == nil {
            current.children[key] = NewTrieNode()
        }
        current = current.children[key]
    }
    current.isValue = true
    current.value = value
}

// Get retrieves the value for a given key from the trie
func (t *Trie) Get(keys []string) (int, bool) {
    current := t.root
    for _, key := range keys {
        next, ok := current.children[key]
        if !ok {
            return 0, false
        }
        current = next
    }
    if current.isValue {
        return current.value, true
    }
    return 0, false
}

func main() {
    trie := NewTrie()
  
    // Populate the trie with nested keys (similar to nested map initialization)
    trie.Put([]string{"Alice", "Math"}, 95)
    trie.Put([]string{"Alice", "Science"}, 88)
    trie.Put([]string{"Bob", "Math"}, 85)
    trie.Put([]string{"Bob", "Science"}, 92)
    trie.Put([]string{"Charlie", "Math"}, 78)
  
    // Get values from the trie using keys
    mathScore, exists := trie.Get([]string{"Alice", "Math"})
    if exists {
        fmt.Println("Alice's Math score:", mathScore)
    } else {
        fmt.Println("Alice's Math score not found.")
    }

    _, exists = trie.Get([]string{"Charlie", "Science"})
    if !exists {
        fmt.Println("Charlie's Science score not found (which is correct as it's sparse).")
    }

    // You can further optimize the trie for memory use by implementing garbage collection for unused subtrees.
} 