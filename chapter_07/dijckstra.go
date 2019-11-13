package main

import (
	"container/list"
	"fmt"
	"math"
)

var inf = math.Inf(1)

type graph struct {
	edges    []edge
	nodes    []node
	sequence *list.List
}

type edge struct {
	parent node
	child  node
	weight float64
}

type node struct {
	state
	status bool
}

// Состояние узла является картой в которой ключ имя узла, а значение его стоимость
type state map[string]float64

// Возвращает имя узла
func (n node) name() (key string) {
	for key = range n.state {
		return key
	}
	return
}

// Возвращает стоимость узла
func (n node) cost() (value float64) {
	for _, value := range n.state {
		return value
	}
	return
}

// Строковое представление отдельного узла
func (n node) String() string {
	return fmt.Sprintf("%s(%.2f)", n.name(), n.cost())
}

// Строковое представление искомого списка графа
func (g *graph) String() string {
	f := g.sequence.Front()
	s := fmt.Sprintf("%v", f.Value)
	for i := 1; i < g.sequence.Len(); i++ {
		if f.Value == nil {
			break
		}
		s = s + fmt.Sprintf("-->%v", f.Next().Value)
		f = f.Next()
	}
	return s
}

// Добавляет новое ребро в граф
func (g *graph) addEdge(p, c node, w float64) {
	// проверка на существование родительского и дочернего узла
	if len(p.state) == 0 || len(c.state) == 0 {
		// ?может быть нужно сообщение что один из узлов ребра не существует
		return
	}
	// проверка на существование такого же ребра в графе
	for i := range g.edges {
		if g.edges[i].parent.name() == p.name() && g.edges[i].child.name() == c.name() {
			// ?может быть нужно сообщение о существовании такого ребра, если найдется
			return
		}
	}
	e := edge{p, c, w}
	g.edges = append(g.edges, e)
}

func (g *graph) addNode(n string) node {
	// проверка на существование такого же узла в графе
	for i := range g.nodes {
		if g.nodes[i].name() == n {
			// ?может быть нужно сообщение о существовании такого узла, если найдется
			return node{}
		}
	}
	// по умолчанию узел добавляется со стоимостью равной бесконечености и включенным статусом
	nn := node{state: map[string]float64{n: inf}, status: true}
	g.nodes = append(g.nodes, nn)
	return nn
}

// Исключает узел из дальнейших поисков выключая его статус
func (g *graph) offNode(n node) {
	for i := range g.nodes {
		if g.nodes[i].name() == n.name() {
			g.nodes[i].status = false
			return
		}
	}
}

// Поиск узла с минимальной стоимостью относительно заданного узла
func (g *graph) lowCostNode(n node) string {
	// обхожу граф по ребрам и нахожу ребра у которых родитель=заданный узел и собираю в список дочерние узлы
	// в момент сбора дочерних узлов их стоимость надо установить равной текущая стоимость+вес ребра
	// обхожу этот список в поиске узла с минимальной стоимостью (и статусом=труу?, чтоб не проверять выключенные) и возвращаю его
	// если две вершины имеют одинаковую стоимость, то будет выбрана одна из них, а вторая выберется на втором шаге
	var nodes []node
	for i := range g.edges {
		if g.edges[i].parent.name() == n.name() {
			if !g.edges[i].child.status {
				continue
			}
			w := g.edges[i].weight
			c := g.edges[i].child.cost()
			if c == inf {
				c = w
			}
			c = c + w
			g.setCostNode(g.edges[i].child.name(), c)
			nodes = append(nodes, g.edges[i].child)
		}
	}
	min := minCost(nodes)
	return min.name()
}

func minCost(ns []node) node {
	if len(ns) == 2 {
		if ns[0].cost() < ns[1].cost() {
			return ns[0]
		}
		return ns[1]
	}
	min := minCost(ns[1:])
	if ns[0].cost() < min.cost() {
		return ns[0]
	}
	return min
}

// Установка стоимости узла
func (g *graph) setCostNode(n string, c float64) {
	for i := range g.nodes {
		if g.nodes[i].name() == n {
			g.nodes[i].state[n] = c
			return
		}
	}
}

// Запуск основного алгоритма поиска по заданным начальному и конечному узлу
// в качестве возвращаемого значения можно очередь или двусвязный список
func (g *graph) Dijkstra(start, end node) *list.List {
	// у стартого узла надо установить стоимость = 0 и вставить его в начало списка
	g.setCostNode(start.name(), 0)
	g.sequence = list.New()
	g.sequence.PushFront(start)
	// g.sequence.InsertAfter(end, e)
	g.sequence.PushBack(end)
	fmt.Println(g.lowCostNode(start))
	return g.sequence
}

func main() {
	g := &graph{}
	s := g.addNode("start")
	a := g.addNode("a")
	b := g.addNode("b")
	c := g.addNode("с")

	g.addEdge(s, a, 7)
	g.addEdge(s, b, 7)
	g.addEdge(b, c, 14)
	g.addEdge(a, c, 25)

	// g.offNode(a)
	// g.setCostNode("b", 777)

	// fmt.Println(b)

	g.Dijkstra(s, c)

	fmt.Println(g)
}
