// Версия с представлением узлов и ребер в виде хэш таблиц в 17 раз быстрее версии с обходами узлов и ребер в циклах
// и это при небольшом графе. Вообще не требует аллокаций памяти.
// goos: windows
// goarch: amd64
// pkg: github.com/NuclearLouse/grokking_algorithms_Aditya_Bhargava/chapter_07
// BenchmarkDijckstra-2     5024041               222 ns/op               0 B/op          0 allocs/op
// PASS
// ok      github.com/NuclearLouse/grokking_algorithms_Aditya_Bhargava/chapter_07  1.496s
package main

import (
	"container/list"
	"fmt"
	"math"
)

type graph struct {
	nodes
	edges
	sequence *list.List
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
	nn := node{name: n, cost: math.Inf(1), status: true}
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

func (g *graph) setCost(n node, c float64) node {
	old := g.nodes[n.name]
	g.nodes[n.name] = node{old.name, c, old.status, old.neighbour}
	return g.nodes[n.name]
}

func (g *graph) offNode(n node) {
	old := g.nodes[n.name]
	g.nodes[n.name] = node{old.name, old.cost, false, old.neighbour}
}

func (g *graph) checkEdge(p1, p2 string) float64 {
	if e, ok := g.edges[p1+"-"+p2]; ok {
		return e.weight
	}
	return g.edges[p2+"-"+p1].weight
}

func minCost(n []node) node {
	if len(n) == 2 {
		if n[0].cost < n[1].cost {
			return n[0]
		}
		return n[1]
	}
	min := minCost(n[1:]) //рекурсия
	if n[0].cost < min.cost {
		return n[0]
	}
	return min
}

func (g *graph) findLowCosts(n node) {
	var listNeigbours []node                // объявляю будущий список соседей для поиска мин.
	neighbours := g.nodes[n.name].neighbour // получаю список соседей начальной ноды
	for i := range neighbours {             // начинаю обход списка
		neighbour := g.nodes[neighbours[i]] // выбираю i-го соседа
		if !neighbour.status {              // проверяю, чтоб он был активен
			continue // если отключен иду к другому
		}
		cw := g.nodes[n.name].cost + g.checkEdge(n.name, neighbours[i])
		if neighbour.cost > cw { // сравниваю стоимость соседа со стоимостью база+ребро
			neighbour = g.setCost(neighbour, cw) // если сосед дороже, то делаю его дешевле
		}
		listNeigbours = append(listNeigbours, neighbour) // добавляю соседа в список соседей для поиска мин.
	} // повторяю со всеми соседями
	g.offNode(n) // отключаю базу

	if len(listNeigbours) <= 1 { // прекращаю если в списке 1 или меньше узлов
		return
	}
	lowNode := minCost(listNeigbours) // нахожу соседа с минимальной стоимостью
	g.findLowCosts(lowNode)           // рекурсивно повторяю поиск соседей для найденного минимального узла
}

func (g *graph) createRoute(n node) {
	neighbours := g.nodes[n.name].neighbour
	for i := range neighbours {
		if (n.cost - g.checkEdge(n.name, neighbours[i])) == 0 {
			// if g.nodes[neighbours[i]].cost == 0 {
			g.sequence.PushFront(g.nodes[neighbours[i]])
			return
		}
		if g.nodes[neighbours[i]].cost == (n.cost - g.checkEdge(n.name, neighbours[i])) {
			g.sequence.PushFront(g.nodes[neighbours[i]])
			g.createRoute(g.nodes[neighbours[i]])
		}
	}
}

func (g *graph) dijkstra(start node, end ...node) {
	n := g.setCost(start, 0)
	g.findLowCosts(n)
	if len(end) != 0 {
		g.sequence = list.New()
		endNode := g.nodes[end[0].name]
		g.sequence.PushBack(endNode)
		g.createRoute(endNode)
		f := g.sequence.Front()
		s := fmt.Sprintf("%v", f.Value)
		for i := 1; i < g.sequence.Len(); i++ {
			if f.Value == nil {
				break
			}
			s = s + fmt.Sprintf("-->%v", f.Next().Value)
			f = f.Next()
		}
		fmt.Println("Кратчайший путь: ", s)
	}
}

func main() {
	g := newGraph()
	a1 := g.addNode("1")
	b2 := g.addNode("2")
	c3 := g.addNode("3")
	d4 := g.addNode("4")
	e5 := g.addNode("5")
	f6 := g.addNode("6")

	g.addEdge(a1, b2, 7)
	g.addEdge(a1, c3, 9)
	g.addEdge(a1, f6, 14)
	g.addEdge(b2, c3, 10)
	g.addEdge(b2, d4, 15)
	g.addEdge(c3, d4, 11)
	g.addEdge(f6, c3, 2)
	g.addEdge(e5, f6, 9)
	g.addEdge(d4, e5, 6)

	g.dijkstra(e5, a1) //при поиске e5-a1 ошибки
	fmt.Println("Весь граф:", g)

}
