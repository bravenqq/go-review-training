/*
带权的求起点到终点路径
使用狄克斯特拉算法
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	//graph表示从某一点到所有邻居节点的距离
	graph := make(map[string]map[string]int)
	graph["start"] = make(map[string]int)
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2
	graph["a"] = make(map[string]int)
	graph["a"]["fin"] = 1
	graph["b"] = make(map[string]int)
	graph["b"]["a"] = 3
	graph["b"]["fin"] = 5
	graph["fin"] = make(map[string]int)

	//costs 表示从起点到某点的开销
	costs := make(map[string]int)
	costs["a"] = 6
	costs["b"] = 2
	costs["fin"] = math.MaxInt64

	//parents 记录对应的父节点
	parents := make(map[string]string)
	parents["a"] = "start"
	parents["b"] = "start"
	processed := make(map[string]bool)
	var findLowestCostNode func() string
	findLowestCostNode = func() string {
		lowestCost := math.MaxInt64
		var n string
		for node, cost := range costs {
			if cost < lowestCost && processed[node] == false {
				lowestCost = cost
				n = node
			}
		}
		return n
	}
	node := findLowestCostNode()
	for len(node) != 0 {
		cost := costs[node]
		neighbors := graph[node]
		for n := range neighbors {
			newCosts := cost + neighbors[n]
			if costs[n] > newCosts {
				costs[n] = newCosts
				parents[n] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode()
	}

	fmt.Println("从start到fin最短路径长度：", costs["fin"])
	for k, v := range parents {
		fmt.Println(k, "--->", v)
	}
	path := "fin"
	p := path
	for parents[p] != "start" {
		path = path + "<---" + parents[p]
		p = parents[p]
	}
	path = path + "<---" + "start"
	fmt.Println("最短路径：", path)
}
