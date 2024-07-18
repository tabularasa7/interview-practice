package depthfirst

import "fmt"

// type vertex struct {
// 	data     int
// 	visited  bool
// 	distance int
// }

// type vertexQueue []vertex

func CreateGraph() map[int][]int {
	// vertex1 := vertex{data: 1}
	// vertex2 := vertex{data: 2}
	// vertex3 := vertex{data: 3}
	// vertex4 := vertex{data: 4}
	// vertex5 := vertex{data: 5}

	adjacencyList := map[int][]int{
		1: {3, 4},
		2: {5},
		3: {1, 5},
		4: {1, 5},
		5: {2, 3, 4},
	}

	return adjacencyList
}

func DFSearch(adjacencyList map[int][]int, val int) int {
	visited := make(map[int]bool)
	depth := 0

	return search(adjacencyList, visited, val, depth)

}

func search(adjacencyList map[int][]int, visited map[int]bool, val, depth int) int {
	visited[val] = true
	fmt.Printf("Depth first traversal node: %d\n", val)

	if adjV, ok := adjacencyList[val]; ok {
		for _, vertex := range adjV {
			if !visited[vertex] {
				depth = search(adjacencyList, visited, vertex, depth)
			}
		}
	}

	return depth
}
