package main

import (
	"fmt"
)

type Node struct {
	value int
}

var g ItemGraph

type ItemGraph struct {
	node  []*Node
	edges map[Node][]*Node
}

func (node *Node) String() string {
	return fmt.Sprintf("%v", node.value)
}

//AddNode adds a new node to the graph
func (graph *ItemGraph) AddNode(node *Node) {
	graph.node = append(graph.node, node)
}

//AddEdge adds an edge to link n1 to n2
func (graph *ItemGraph) AddEdge(n1, n2 *Node) {
	if graph.edges == nil {
		graph.edges = make(map[Node][]*Node)
	}
	//Adjecency list using maps and slices
	graph.edges[*n1] = append(graph.edges[*n1], n2)
	graph.edges[*n2] = append(graph.edges[*n2], n1)
}

//A-B-C
//C-S
//

func (graph *ItemGraph) String() {
	s := ""
	for i := 0; i < len(graph.node); i++ {
		s += graph.node[i].String() + " -> "
		near := graph.edges[*graph.node[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func FillGraph() {
	nA := Node{1}
	nB := Node{2}
	nC := Node{3}
	nD := Node{4}

	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nC, &nD)
}

//BFS Search
func (graph *ItemGraph) Bfs(f func(*Node)) {
	queue := make([]*Node, 0)
	visited := make(map[*Node]bool)
	n := graph.node[0]
	//Enqueue operation
	queue = append(queue, n)
	for {
		if len(queue) == 0 {
			break
		}
		//Dequeue oparation
		node := queue[0]
		queue = append(queue[:0], queue[1:]...)

		visited[node] = true
		near := graph.edges[*node]
		fmt.Println("Near edges", near)
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				queue = append(queue, j)
				fmt.Println("Queue state:", queue)
				visited[j] = true
			}
		}
		if f != nil {
			f(node)
		}
	}
}

func main() {
	//Undirected graph A->B then B->A
	FillGraph()
	g.String()

	g.Bfs(func(n *Node) {
		fmt.Println(n.String())
	})
}
