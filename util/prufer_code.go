package util

import (
	"slices"

	"github.com/dq1Mango/gograph"
)

func GraphFromPruferCode(prufer ...uint) (gograph.Graph[uint], error) {
	sequence := make([]uint, len(prufer)+2)

	for i := range sequence {
		sequence[i] = uint(i)
	}

	graph := gograph.New[uint]()

	for len(prufer) > 0 {

		var smallest uint
		for _, i := range sequence {

			if !slices.Contains(prufer, i) {
				smallest = i
			}
		}

		graph.AddEdge(gograph.NewVertex(prufer[0]), gograph.NewVertex(smallest))

		sequence = sequence[1:]
		prufer = prufer[1:]
	}

	graph.AddEdge(gograph.NewVertex(sequence[0]), gograph.NewVertex(sequence[1]))

	return graph, nil
}

func PruferCodeFromGraph(graph gograph.Graph[uint]) ([]uint, error) {
	verticies := graph.GetAllVertices()

	// ids := make([]uint, len(verticies))
	//
	// for i, vertex := range verticies {
	// 	ids[i] = vertex.Label()
	// }

	slices.SortFunc(verticies,
		func(v1, v2 *gograph.Vertex[uint]) int {
			if v1.Label() > v2.Label() {
				return 1
			} else if v1.Label() == v2.Label() {
				return 0
			} else {
				return -1
			}
		},
	)

	pruferCode := make([]uint, len(verticies)-2)

	for i := range pruferCode {

		var smallest_vertex int
		var neighbor uint

		for index, vertex := range verticies {
			// we shall see this counts as a 'leaf' for a directed graph
			if vertex.InDegree() == 1 {
				neighbor = vertex.Neighbors()[0].Label()
				smallest_vertex = index
				break
			}
		}

		pruferCode[i] = neighbor

		graph.RemoveVertices(verticies[smallest_vertex])

		verticies = slices.Delete(verticies, smallest_vertex, smallest_vertex+1)
	}

	return pruferCode, nil
}

// func sliceContains(slice []uint, item uint) bool {
//
// 	for _, value := range slice {
// 		if value == item {
// 			return false
// 		}
// 	}
//
// 	return true
// }
//
// func smallest(sequence []uint, set []uint) uint {}
