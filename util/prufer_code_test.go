package util

import (
	"errors"
	"slices"
	"testing"

	"github.com/dq1Mango/gograph"
)

func graphAndCode() (gograph.Graph[uint], []uint) {
	graph := gograph.New[uint]()

	edges := []struct{ u, v uint }{{0, 1}, {1, 2}, {2, 3}, {2, 4}}

	// graph.AddEdge(gograph.NewVertex(0), gograph.NewVertex(1))
	// graph.AddEdge(gograph.NewVertex(1), gograph.NewVertex(2))
	// graph.AddEdge(gograph.NewVertex(2), gograph.NewVertex(3))
	// graph.AddEdge(gograph.NewVertex(2), gograph.NewVertex(4))

	for _, edge := range edges {
		graph.AddEdge(gograph.NewVertex(edge.u), gograph.NewVertex(edge.v))
	}

	return graph, []uint{1, 2, 2}

}

func TestPruferCodeFromGraph(t *testing.T) {

	graph, expected := graphAndCode()

	pruferCode, err := PruferCodeFromGraph(graph.Clone())

	if err != nil {
		t.Fatal(err)
	}

	if !slices.Equal(pruferCode, expected) {
		t.Errorf("Error: expected: %v got: %v", expected, pruferCode)
	}

	graph.AddEdge(gograph.NewVertex(uint(3)), gograph.NewVertex(uint(4)))

	_, err = PruferCodeFromGraph(graph)

	if !errors.Is(err, &NonTreeError{}) {
		t.Errorf("Expected error: %T, got: %T", &NonTreeError{}, err)
	}
}

func TestGraphFromPruferCode(t *testing.T) {
	expected, pruferCode := graphAndCode()

	graph, err := GraphFromPruferCode(pruferCode...)

	if err != nil {
		t.Fatal(err)
	}

	if expected.Size() != graph.Size() {
		t.Errorf("generated graph of wrong size")
	}
	if expected.Order() != graph.Order() {
		t.Errorf("generated graph of wrong order")
	}

	for _, edge := range graph.AllEdges() {
		if !expected.ContainsEdge(
			gograph.NewVertex(edge.Source().Label()),
			gograph.NewVertex(edge.Destination().Label()),
		) {
			t.Errorf("graphs not equal")
		}
	}

	pruferCode[2] = uint(len(pruferCode)) + 2

	if _, err := GraphFromPruferCode(pruferCode...); !errors.Is(err, &InvalidPruferCodeError{}) {
		t.Errorf("Expected error: %T, got: %T", &InvalidPruferCodeError{}, err)
	}

}
