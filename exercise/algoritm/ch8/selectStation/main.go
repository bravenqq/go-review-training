package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

/*求如何选择最少的station能覆盖全美的州*/
func main() {
	var selectedStation []string
	needStates := mapset.NewSetFromSlice([]interface{}{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"})
	stations := make(map[string]mapset.Set)
	stations["kone"] = mapset.NewSetFromSlice([]interface{}{"id", "nv", "ut"})
	stations["ktwo"] = mapset.NewSetFromSlice([]interface{}{"wa", "id", "mt"})
	stations["kthree"] = mapset.NewSetFromSlice([]interface{}{"or", "nv", "ca"})
	stations["kfour"] = mapset.NewSetFromSlice([]interface{}{"nv", "ut"})
	stations["kfive"] = mapset.NewSetFromSlice([]interface{}{"ca", "az"})
	//1.选择一个广播台，覆盖最多未覆盖的州
	//2.重复步骤1直到没有为覆盖的州
	var station string
	var stateCorved int
	var interStates mapset.Set
	for needStates.Cardinality() != 0 {
		station = ""
		stateCorved = 0
		//找出覆盖最多的未覆盖州的广播台
		for k, st := range stations {
			inter := needStates.Intersect(st)
			if inter.Cardinality() > stateCorved {
				station = k
				interStates = inter
			}
		}
		//添加广播台到stasion
		selectedStation = append(selectedStation, station)
		//needStates除去选中广播台所覆盖州与为覆盖州的交集
		needStates = needStates.Difference(interStates)
		//除去已选择的州
		delete(stations, station)
	}
	fmt.Printf("selected stations is:%v\n", selectedStation)
}
