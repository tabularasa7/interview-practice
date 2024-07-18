package breadthfirst

import "fmt"

type vertex struct {
	data     int
	visited  bool
	distance int
}

type vertexQueue []vertex

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

func BFSearch(adjacencyList map[int][]int, start, val int) (int, error) {
	var queue vertexQueue
	qVertex := vertex{data: start, visited: false, distance: 0}
	queue.Push(qVertex)

	for !queue.isEmpty() && qVertex.data != val {
		qVertex = queue.Pop()
		qVertex.visited = true

		edges, ok := adjacencyList[qVertex.data]

		if !ok {
			return 0, fmt.Errorf("vertex does not exist in graph")
		}

		for _, vertexVal := range edges {
			queue.Push(vertex{data: vertexVal, distance: qVertex.distance + 1})
		}
	}

	return qVertex.distance, nil

}

func (q *vertexQueue) Push(val vertex) {
	*q = append(*q, val)
}

func (q *vertexQueue) Pop() vertex {
	if q.isEmpty() {
		return vertex{}
	}

	val := (*q)[0]
	*q = (*q)[1:]
	return val

}

func (q *vertexQueue) isEmpty() bool {
	return len(*q) == 0
}
