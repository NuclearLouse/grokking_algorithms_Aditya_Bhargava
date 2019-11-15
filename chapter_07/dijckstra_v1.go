// Версия с обходами узлов и ребер  в циклах даже при небольшом графе в 17 раз медленнее версии
// с представлением узлов и ребер в виде хэш таблиц. И требует операций с памятью.
// goos: windows
// goarch: amd64
// pkg: github.com/NuclearLouse/grokking_algorithms_Aditya_Bhargava/chapter_07
// BenchmarkDijckstra-2      285903              3600 ns/op             160 B/op          4 allocs/op
// PASS
// ok      github.com/NuclearLouse/grokking_algorithms_Aditya_Bhargava/chapter_07  1.206s
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

// Ребро графа состоящее из двух узлов-точек и имеющее вес
type edge struct {
	point1 node
	point2 node
	weight float64
}

// Узел графа имеющий базовое состояние имя:стоимость и статус - активен узел или нет
type node struct {
	state
	status bool
}

// Состояние узла является картой в которой ключ имя узла, а значение его стоимость
type state map[string]float64

func newGraph() *graph {
	return &graph{}
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

// Добавляет новое ребро в граф
func (g *graph) addEdge(p1, p2 node, w float64) {
	// проверка на существование родительского и дочернего узла
	if len(p1.state) == 0 || len(p2.state) == 0 {
		// ?может быть нужно сообщение что один из узлов ребра не существует
		return
	}
	// проверка на существование такого же ребра в графе
	for i := range g.edges {
		if g.edges[i].point1.name() == p1.name() && g.edges[i].point2.name() == p2.name() {
			// ?может быть нужно сообщение о существовании такого ребра, если найдется
			return
		}
	}
	e := edge{p1, p2, w}
	g.edges = append(g.edges, e)
}

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

// Строковое представление результатов графа
func (g *graph) String() (s string) {
	for i := range g.nodes {
		s = s + fmt.Sprintf("%v ", g.nodes[i])
	}
	return s
}

// Исключает узел из дальнейших поисков выключая его статус
func (g *graph) offNode(n node) {
	for i := range g.nodes {
		if g.nodes[i].name() == n.name() {
			g.nodes[i].status = false
			break
		}
	}
	for i := range g.edges {
		switch n.name() {
		case g.edges[i].point1.name():
			g.edges[i].point1.status = false
			break
		case g.edges[i].point2.name():
			g.edges[i].point2.status = false
			break
		}
	}
	return
}

// Установка стоимости узла
func (g *graph) setCostNode(n node, c float64) {
	name := n.name()
	for i := range g.nodes {
		if g.nodes[i].name() == name {
			g.nodes[i].state[name] = c
			return
		}
	}
}

//Рекурсивный поиск минимального элемента в списке
func minCost(ns []node) node {
	if len(ns) == 2 {
		if ns[0].cost() < ns[1].cost() {
			return ns[0]
		}
		return ns[1]
	}
	min := minCost(ns[1:]) //рекурсия
	if ns[0].cost() < min.cost() {
		return ns[0]
	}
	return min
}

// Обновление стоимости узла
func (g *graph) updateCost(ns []node, n node, cw float64) []node {
	if !n.status {
		return ns
	}
	c := n.cost()
	if c > cw {
		c = cw
	}
	g.setCostNode(n, c)
	return append(ns, n)
}

// Поиск узла с минимальной стоимостью относительно заданного узла
func (g *graph) lowCostNode(n node) {
	// Обхожу граф по ребрам и нахожу ребра у которых одна из точек заданный узел и собираю в список противоположный узел.
	// В момент добавления узла в список, определяется и выставляется его стоимость равная текущей стоимости + вес ребра
	// Обхожу список этих узлов в поиске узла с наименьшей стоимостью и при условии, что узел активен.
	var nodes []node
	for i := range g.edges {
		switch n.name() {
		case g.edges[i].point1.name():
			nodes = g.updateCost(nodes, g.edges[i].point2, g.edges[i].point1.cost()+g.edges[i].weight)

		case g.edges[i].point2.name():
			nodes = g.updateCost(nodes, g.edges[i].point1, g.edges[i].point2.cost()+g.edges[i].weight)
		}
	}
	// fmt.Println(nodes)
	g.offNode(n)
	// fmt.Println("off:", n)
	if len(nodes) <= 1 {
		return
	}
	min := minCost(nodes) //рекурсия
	// fmt.Println("min:", min)
	g.lowCostNode(min) //рекурсия
	return
}

// Запуск основного алгоритма поиска по заданным начальному и конечному узлу
// в качестве возвращаемого значения можно использовать очередь или двусвязный список
// второй аргумент функции(end) является не обязательным. Если его нет, то функция просто подсчитает стоимость
// всех узлов графа относительно стартового узла. Если он есть, то функция выведет кратчайший маршрут от начала до конца.
// Еще не дописано...
func (g *graph) dijkstra(start node) {
	g.setCostNode(start, 0)
	g.lowCostNode(start)
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

	g.dijkstra(a1)
	fmt.Println(g)
}
