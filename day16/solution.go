package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strings"
)

type State struct {
	cost, x, y int
	facing     string
	path       [][2]int
}

type PriorityQueue []State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(State))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func findUniqueNodes(parents map[[3]int]map[[3]int]struct{}, endState [3]int, uniqueNodes map[[2]int]struct{}) {
	for node := range parents[endState] {
		findUniqueNodes(parents, node, uniqueNodes)
	}
	uniqueNodes[[2]int{endState[0], endState[1]}] = struct{}{}
}

func dijkstraPathfinding(grid [][]rune, start, end [2]int) (int, map[[2]int]struct{}) {
	directions := []struct {
		dx, dy int
		name   string
	}{
		{0, 1, "right"},
		{0, -1, "left"},
		{1, 0, "down"},
		{-1, 0, "up"},
	}
	directionMap := map[string]int{"right": 0, "up": 1, "left": 2, "down": 3}

	Sx, Sy := start[0], start[1]
	Ex, Ey := end[0], end[1]
	n, m := len(grid), len(grid[0])
	uniqueNodes := make(map[[2]int]struct{}) //this is the set of unique nodes that we have visited on the way to the goal

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{cost: 0, x: Sx, y: Sy, facing: "right", path: [][2]int{{Sx, Sy}}})
	visited := make(map[[3]int]int) // (x, y, facing_index)->cost
	costToGoal := math.MaxInt
	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)
		// Goal check
		if current.x == Ex && current.y == Ey {
			//check current direction against others, in case more than one neighbor of the goal leads into it.
			canBreak := false
			for i := range directions {
				costAtKey, ok := visited[[3]int{current.x, current.y, i}]
				if ok && costAtKey < current.cost {
					canBreak = true
					break
				}
			}
			if canBreak {
				continue
			}
			for _, node := range current.path {
				uniqueNodes[node] = struct{}{}
			}
			costToGoal = current.cost
		}
		// check if we have already visited this node going this direction with a lower cost
		key := [3]int{current.x, current.y, directionMap[current.facing]}
		bestCostAtKey, ok := visited[key]
		if (ok && bestCostAtKey < current.cost) || costToGoal < current.cost {
			continue
		}
		visited[key] = current.cost

		// Explore neighbors
		for _, dir := range directions {
			nx, ny := current.x+dir.dx, current.y+dir.dy
			if nx >= 0 && ny >= 0 && nx < n && ny < m && grid[nx][ny] != '#' {
				turnCost := 0
				if current.facing != dir.name {
					diff := int(math.Abs(float64((directionMap[current.facing] - directionMap[dir.name]))))
					if diff == 3 {
						diff = 1
					}
					turnCost += 1000 * diff
				}
				newCost := current.cost + 1 + int(turnCost)
				copyPath := make([][2]int, len(current.path))
				copy(copyPath, current.path)
				copyPath = append(copyPath, [2]int{nx, ny})
				newState := State{cost: newCost, x: nx, y: ny, facing: dir.name, path: copyPath}
				newKey := [3]int{nx, ny, directionMap[dir.name]}
				if _, ok := visited[newKey]; !ok {
					heap.Push(pq, newState)
				}
			}
		}
	}
	return costToGoal, uniqueNodes
}

func printMap(uniqueNodes map[[2]int]struct{}, grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			found := false
			if _, ok := uniqueNodes[[2]int{i, j}]; ok {
				fmt.Print("O")
				found = true
			}

			if !found {
				fmt.Print(string(grid[i][j]))
			}
		}
		fmt.Println()
	}
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := os.ReadFile("../inputs/day16.txt")
	check(err)
	gridLines := strings.Split(string(input), "\n")
	grid := make([][]rune, len(gridLines))
	var start [2]int
	var end [2]int
	for i, line := range gridLines {
		grid[i] = []rune(line)
		for j, char := range line {
			if char == 'S' {
				start = [2]int{i, j}
			} else if char == 'E' {
				end = [2]int{i, j}
			}
		}
	}
	cost, uniqueNodes := dijkstraPathfinding(grid, start, end)
	if cost != -1 {
		fmt.Printf("Minimum cost to reach the goal: %d\n", cost)
		fmt.Println("Path length:", len(uniqueNodes))
	} else {
		fmt.Println("No path found!")
	}
}
