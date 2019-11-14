// Версия с представлением узлов и ребер в виде хэш таблиц
package main

import (
	"fmt"
	"math"
)

var inf = math.Inf(1)

type graph struct {
	nodes
	edges
}

type nodes map[string]node

type edges map[string]edge

type node struct {
	name      string
	cost      float64
	status    bool
	neighbour []string
}

func (n node) String() string {
	return fmt.Sprintf("%s(%.2f)", n.name, n.cost)
}

func (g *graph) String() (s string) {
	for i := range g.nodes {
		s = s + fmt.Sprintf("%v ", g.nodes[i])
	}
	return s
}

type edge struct {
	name   string
	weight float64
}

func newGraph() *graph {
	return &graph{nodes: make(map[string]node), edges: make(map[string]edge)}
}

func (g *graph) addNode(n string) node {
	nn := node{name: n, cost: inf, status: true}
	g.nodes[n] = nn
	return nn
}

func (g *graph) addEdge(p1, p2 node, w float64) {
	n := p1.name + "-" + p2.name
	e := edge{name: n, weight: w}
	g.edges[n] = e
	g.addNeighbour(p1, p2.name)
	g.addNeighbour(p2, p1.name)
}

func (g *graph) addNeighbour(n node, nn string) {
	old := g.nodes[n.name]
	g.nodes[n.name] = node{old.name, old.cost, old.status, append(old.neighbour, nn)}
}

func (g *graph) setCost(n node, c float64) {
	old := g.nodes[n.name]
	g.nodes[n.name] = node{old.name, c, old.status, old.neighbour}
}

func (g *graph) offNode(n node) {
	old := g.nodes[n.name]
	g.nodes[n.name] = node{old.name, old.cost, false, old.neighbour}
}

func (g *graph) findLowCosts(n node) {

}

func (g *graph) dijkstra(start node) {
	g.setCost(start, 0)
	g.findLowCosts(start)
}

func main() {
	g := newGraph()
	a := g.addNode("a")
	b := g.addNode("b")
	c := g.addNode("c")
	d := g.addNode("d")
	g.addEdge(a, b, 65)
	g.addEdge(a, c, 74)
	g.addEdge(c, b, 33)
	g.addEdge(c, d, 40)
	g.addEdge(b, d, 50)

	// fmt.Println("Узел в графе по имени:", g.nodes["a"])
	// fmt.Println("Статус узла в графе по имени:", g.nodes["a"].status)
	// fmt.Println("Соседи узла:", g.nodes["a"].neighbour)
	// fmt.Println("***Updates***")

	// g.setCost(b, 123)
	// g.offNode(b)

	// fmt.Println(g.nodes["b"], g.nodes["b"].status)
	// fmt.Println("Весь граф:", g)

	g.dijkstra(a)

}
