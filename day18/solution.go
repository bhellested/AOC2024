package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	cost, x, y int
	path [][2]int
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

	Sx, Sy := start[0], start[1]
	Ex, Ey := end[0], end[1]
	n, m := len(grid), len(grid[0])
	uniqueNodes := make(map[[2]int]struct{}) //this is the set of unique nodes that we have visited on the way to the goal

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, State{cost: 0, x: Sx, y: Sy, path: [][2]int{{Sx, Sy}}})
	visited := make(map[[2]int]int) // (x, y)->cost
	costToGoal := -1
	for pq.Len() > 0 {
		current := heap.Pop(pq).(State)
		// Goal check
		if current.x == Ex && current.y == Ey {
			costToGoal = current.cost
			break
		}
		// check if we have already visited this node going this direction with a lower cost
		key := [2]int{current.x, current.y}
		_, ok := visited[key]
		if ok  {
			continue
		}
		visited[key] = current.cost

		// Explore neighbors
		for _, dir := range directions {
			nx, ny := current.x+dir.dx, current.y+dir.dy
			if nx >= 0 && ny >= 0 && nx < n && ny < m && grid[nx][ny] != '#' {
				newCost := current.cost + 1
				copyPath := make([][2]int, len(current.path))
				copy(copyPath, current.path)
				copyPath = append(copyPath, [2]int{nx, ny})
				newState := State{cost: newCost, x: nx, y: ny, path: copyPath}
				newKey := [2]int{nx, ny}
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
	input, err := os.ReadFile("../inputs/day18.txt")
	check(err)
	gridLines := strings.Split(string(input), "\n")
	grid := make([][]rune, 71)
	for i := range grid {
		grid[i] = make([]rune, 71)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	//this is a little bit slow for part2. I could do it a little faster by either using A* for quicker pathfinding, or binarysearching the grids but it only takes a couple seconds so It's sufficient
	for i, line := range gridLines {
		x, err := strconv.Atoi(strings.Split(line, ",")[0])
		check(err)
		y, err := strconv.Atoi(strings.Split(line, ",")[1])
		check(err)
		grid[x][y] = '#'
		cost, _ := dijkstraPathfinding(grid, [2]int{0, 0}, [2]int{70, 70})
		if i==1024{
			fmt.Println("cost at 1024 (part1)",cost)
		}
		if cost ==-1 {
			//printMap(make(map[[2]int]struct{}), grid)
			fmt.Println("No path found at i:",i)
			fmt.Println("x,y (part2):", x,",", y)
			break;
		}
	}
}
