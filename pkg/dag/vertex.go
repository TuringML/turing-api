package dag

import (
	"fmt"

	"github.com/goombaio/orderedset"
)

// Vertex type implements a vertex of a Directed Acyclic graph or DAG.
type Vertex struct {
	ID       string
	Visited  bool
	Value    interface{}
	Parents  *orderedset.OrderedSet
	Children *orderedset.OrderedSet
}

// NewVertex creates a new vertex.
func NewVertex(id string, value interface{}) *Vertex {
	v := &Vertex{
		ID:       id,
		Visited:  false,
		Parents:  orderedset.NewOrderedSet(),
		Children: orderedset.NewOrderedSet(),
		Value:    value,
	}

	return v
}

// Degree return the number of parents and children of the vertex
func (v *Vertex) Degree() int {
	return v.Parents.Size() + v.Children.Size()
}

// InDegree return the number of parents of the vertex or the number of edges
// entering on it.
func (v *Vertex) InDegree() int {
	return v.Parents.Size()
}

// OutDegree return the number of children of the vertex or the number of edges
// leaving it.
func (v *Vertex) OutDegree() int {
	return v.Children.Size()
}

// HasVisited determine if the node has been visited already or not
func (v *Vertex) HasVisited() bool {
	return v.Visited
}

// String implements stringer interface and prints an string representation
// of this instance.
func (v *Vertex) String() string {
	result := fmt.Sprintf("ID: %s - Parents: %d - Children: %d - Value: %v\n", v.ID, v.Parents.Size(), v.Children.Size(), v.Value)

	return result
}
