package depthfirst

import "fmt"

func CreateGraph() map[int][]int {

	adjacencyList := map[int][]int{
		1: {3, 4},
		2: {5},
		3: {1, 5},
		4: {1, 5},
		5: {2, 3, 4},
	}

	return adjacencyList
}

func DFSearch(adjacencyList map[int][]int, startingVal int) int {
	visited := make(map[int]bool)
	depth := 0

	return search(adjacencyList, visited, startingVal, depth)

}

func search(adjacencyList map[int][]int, visited map[int]bool, val, depth int) int {
	visited[val] = true
	fmt.Printf("Depth first traversal node: %d\n", val)

	if adjV, ok := adjacencyList[val]; ok {
		for _, vertex := range adjV {
			if _, ok := visited[vertex]; !ok {
				depth = search(adjacencyList, visited, vertex, depth)
			}
		}
	}

	return depth
}
