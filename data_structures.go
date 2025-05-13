// Exemplo de uso de diferentes estruturas de dados em Go

package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

// 1. Array (tamanho fixo)
func arrayExample() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	fmt.Println("Array:", arr)
}

// 2. Slice (tamanho dinâmico)
func sliceExample() {
	s := []int{1, 2, 3}
	s = append(s, 4)
	fmt.Println("Slice:", s)
}

// 3. Lista Encadeada (usando container/list)
func listExample() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
}

// 4. Pilha (Stack) - usando slice
type Stack []int

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Pop() int {
	l := len(*s)
	v := (*s)[l-1]
	*s = (*s)[:l-1]
	return v
}

func stackExample() {
	var s Stack
	s.Push(1)
	s.Push(2)
	fmt.Println("Pop:", s.Pop())
}

// 5. Fila (Queue) - usando slice
type Queue []int

func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() int {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func queueExample() {
	var q Queue
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println("Dequeue:", q.Dequeue())
}

// 6. Mapa (HashMap)
func mapExample() {
	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Println("Map:", m)
}

// 7. Árvore Binária e BST
// Nó da árvore
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// Inserção em BST
func bstInsert(root *TreeNode, v int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if v < root.Val {
		root.Left = bstInsert(root.Left, v)
	} else {
		root.Right = bstInsert(root.Right, v)
	}
	return root
}

func bstExample() {
	var root *TreeNode
	vals := []int{5, 3, 7, 2, 4}
	for _, v := range vals {
		root = bstInsert(root, v)
	}
	fmt.Println("BST inserida com valores", vals)
}

// 8. Heap (Min-Heap)
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	h := *h
	v := h[len(h)-1]
	*h = h[:len(h)-1]
	return v
}

func heapExample() {
	h := &IntHeap{5, 2, 8}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Println("Min-Heap:", heap.Pop(h))
}

// 9. Grafo (lista de adjacência)
type Graph struct {
	Adj map[int][]int
}

func graphExample() {
	g := Graph{Adj: make(map[int][]int)}
	g.Adj[1] = []int{2, 3}
	g.Adj[2] = []int{1, 4}
	fmt.Println("Graph adjacency list:", g.Adj)
}

// 10. Trie (Prefix Tree)
type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

func trieExample() {
	root := &TrieNode{Children: make(map[rune]*TrieNode)}
	// Inserir
	word := "go"
	cur := root
	for _, ch := range word {
		if cur.Children[ch] == nil {
			cur.Children[ch] = &TrieNode{Children: make(map[rune]*TrieNode)}
		}
		cur = cur.Children[ch]
	}
	cur.IsWord = true
	fmt.Println("Trie com palavra 'go' inserida")
}

// 11. Union-Find (Disjoint Set)
type UnionFind struct {
	Parent, Rank []int
}

func newUF(n int) *UnionFind {
	uf := &UnionFind{Parent: make([]int, n), Rank: make([]int, n)}
	for i := 0; i < n; i++ {
		uf.Parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.Parent[x] != x {
		uf.Parent[x] = uf.Find(uf.Parent[x])
	}
	return uf.Parent[x]
}

func (uf *UnionFind) Union(a, b int) {
	aRoot := uf.Find(a)
	bRoot := uf.Find(b)
	if aRoot == bRoot {
		return
	}
	if uf.Rank[aRoot] < uf.Rank[bRoot] {
		uf.Parent[aRoot] = bRoot
	} else if uf.Rank[aRoot] > uf.Rank[bRoot] {
		uf.Parent[bRoot] = aRoot
	} else {
		uf.Parent[bRoot] = aRoot
		uf.Rank[aRoot]++
	}
}

func ufExample() {
	uf := newUF(5)
	uf.Union(0, 1)
	fmt.Println("Find(1):", uf.Find(1))
}

func main() {
	arrayExample()
	sliceExample()
	listExample()
	stackExample()
	queueExample()
	mapExample()
	bstExample()
	heapExample()
	graphExample()
	trieExample()
	ufExample()
}
