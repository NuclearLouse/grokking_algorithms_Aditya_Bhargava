package main

import (
	"fmt"

	"github.com/zoumo/goset"
)

func main() {
	statesNeeded := goset.NewSet("mt", "wa", "or", "id", "nv", "ut", "ca", "az")
	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	var selectedStations []string
	for statesNeeded.Len() > 0 {
		statesCovered := goset.NewSet()
		var bestStation string
		for st := range stations { //обхожу карту радиостанций
			states := goset.NewSetFromStrings(stations[st]) //
			covered := statesNeeded.Intersect(states)
			if covered.Len() > statesCovered.Len() {
				bestStation = st
				statesCovered = covered
			}
		}
		selectedStations = append(selectedStations, bestStation)
		delete(stations, bestStation)
		statesNeeded.Remove(statesCovered.Elements()...)
	}

	fmt.Println(selectedStations) // ответ должен быть 1,2,3,5
}
