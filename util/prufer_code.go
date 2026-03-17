package util

import (
	"slices"

	"github.com/hmdsefi/gograph"
)

type InvalidPruferCodeError struct{}

func (e *InvalidPruferCodeError) Error() string {
	return "Prufer Code does not construct valid tree"
}

type NonTreeError struct{}

func (e *NonTreeError) Error() string {
	return "Graph is not a tree"
}

func GraphFromPruferCode(prufer ...uint) (gograph.Graph[uint], error) {
	N := uint(len(prufer) + 2)

	for _, p := range prufer {
		if p >= N {
			return nil, &InvalidPruferCodeError{}
		}
	}

	sequence := make([]uint, N)

	for i := range sequence {
		sequence[i] = uint(i)
	}

	graph := gograph.New[uint]()

	for len(prufer) > 0 {

		var smallest uint
		var index int
		for i, id := range sequence {

			if !slices.Contains(prufer, id) {
				smallest = id
				index = i
				break
			}
		}

		graph.AddEdge(gograph.NewVertex(prufer[0]), gograph.NewVertex(smallest))

		sequence = slices.Delete(sequence, index, index+1)
		prufer = prufer[1:]
	}

	graph.AddEdge(gograph.NewVertex(sequence[0]), gograph.NewVertex(sequence[1]))

	return graph, nil
}

func PruferCodeFromGraph(graph gograph.Graph[uint]) ([]uint, error) {

	// ensure our graph is *probably* a tree
	if graph.Size() < 2 {
		return nil, &NonTreeError{}
	}

	if graph.Size() != graph.Order()-1 {
		return nil, &NonTreeError{}
	}

	verticies := graph.GetAllVertices()

	// sort all of the verticies
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

		var smallest_leaf *int
		var neighbor uint

		// look for the leaf (1 neighbor) with the smallest label
		for index, vertex := range verticies {

			// we shall see this counts as a 'leaf' for a directed graph
			if vertex.InDegree() == 1 {
				neighbor = vertex.Neighbors()[0].Label()
				smallest_leaf = &index

				break
			}
		}

		// if we fail to find a leaf our graph is not a tree
		if smallest_leaf == nil {
			return nil, &NonTreeError{}
		}

		pruferCode[i] = neighbor

		graph.RemoveVertices(verticies[*smallest_leaf])
		verticies = slices.Delete(verticies, *smallest_leaf, *smallest_leaf+1)

	}

	return pruferCode, nil
}
