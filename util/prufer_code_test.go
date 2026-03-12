package util

import (
	"slices"
	"testing"

	"github.com/dq1Mango/gograph"
)

func TestPruferCodeFromGraph(t *testing.T) {

	graph := gograph.New[uint]()

	edges := []struct{ u, v uint }{{0, 1}, {1, 2}, {2, 3}, {2, 4}}

	// graph.AddEdge(gograph.NewVertex(0), gograph.NewVertex(1))
	// graph.AddEdge(gograph.NewVertex(1), gograph.NewVertex(2))
	// graph.AddEdge(gograph.NewVertex(2), gograph.NewVertex(3))
	// graph.AddEdge(gograph.NewVertex(2), gograph.NewVertex(4))

	for _, edge := range edges {
		graph.AddEdge(gograph.NewVertex(edge.u), gograph.NewVertex(edge.v))
	}

	pruferCode, err := PruferCodeFromGraph(graph)

	if err != nil {
		t.Fatal(err)
	}

	expected := []uint{1, 2, 2}

	if !slices.Equal(pruferCode, expected) {
		t.Errorf("Error: expected: %v got: %v", expected, pruferCode)
	}
}
