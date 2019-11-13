// Алгоритм поиска в ширину

package main

import (
	"container/list"
	"fmt"
)

type graph struct {
	g map[string][]string
}

// Как вариант, инициализация карты через конструктор
func newGraph() *graph {
	return &graph{
		g: map[string][]string{
			"you":    {"alice", "bob", "claire"},
			"bob":    {"you", "john", "peggy"},
			"alice":  {"you", "peggy"},
			"claire": {"you", "tom", "johnny"},
			"john":   {"peggy"},
			"peggy":  {"john"},
			"tom":    {"Andrew"},
			"johnny": {},
			"Andrew": {"Nikolai"},
		},
	}
}

func inSearched(person string, searched []string) bool {
	for _, name := range searched {
		if person == name {
			return true
		}
	}
	return false
}

func pushQueue(q *list.List, name string) {
	for _, n := range newGraph().g[name] {
		q.PushBack(n)
	}
}

// Если длина имени человека = 7, то он продавец
func personIsSeller(p string) bool {
	if len(p) == 7 {
		return true
	}
	return false
}

func search(name string) string {
	var q list.List
	var searched []string
	pushQueue(&q, name)
	for q.Len() != 0 {
		person := fmt.Sprintf("%v", q.Remove(q.Front()))
		if !inSearched(person, searched) {
			if personIsSeller(person) {
				return fmt.Sprintf("%s is seller", person)
			}
			pushQueue(&q, person)
			searched = append(searched, person)
		}
	}
	return "Seller not found"
}

func main() {
	fmt.Println(search("peggy"))
}
