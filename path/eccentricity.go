package path

import (
	"math"

	"github.com/hmdsefi/gograph"
)

// Essentially just wrapper on dijkstras, but i want it anyway
func Eccentricity[T comparable](graph gograph.Graph[T], vertex *gograph.Vertex[T]) float64 {

	distances := Dijkstra(graph, vertex.Label())

	max_distance := 0.0
	for _, distance := range distances {

		if distance == math.MaxFloat64 {
			continue
		}

		max_distance = max(distance, max_distance)
	}

	return max_distance
}
