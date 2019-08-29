package main

import "fmt"

type Node uint32

type Graph struct {
	CompactAdjacencyList []Node
	Span                 []uint64
}

func (g *Graph) Neighbours(n Node) []Node {
	start := g.Span[n]
	if int(n) == len(g.Span)-1 {
		return g.CompactAdjacencyList[start:]
	}
	return g.CompactAdjacencyList[start:g.Span[n+1]]
}

type Callback func(Node)

func (g *Graph) BFS(f Callback) {
	visited := NewUintSet(len(g.Span) + 1)
	candidateNodes := IntQueue{}
	candidateNodes.Push(0)
	for {
		if candidateNodes.Size() == 0 {
			break
		}
		n := Node(candidateNodes.Pop())
		if !visited.Contains(uint32(n)) {
			f(n)
			visited.Add(uint32(n))
			for _, neibr := range g.Neighbours(n) {
				if !visited.Contains(uint32(neibr)) {
					candidateNodes.Push(int(neibr))
				}
			}
		}
	}
}

func main() {
	g := Graph{
		CompactAdjacencyList: []Node{2, 4, 3, 0, 3, 0, 1, 2, 0},
		Span:                 []uint64{0, 2, 3, 5, 8},
	}
	g.BFS(func(n Node) { fmt.Println(n) })
}
