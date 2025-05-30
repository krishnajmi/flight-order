package flight

import (
	"context"
	"errors"
)

func (ts *topologicalSorter) Sort(ctx context.Context, graph map[string]string) ([]string, error) {
	visited := make(map[string]bool)
	var stack []string

	// Get all unique locations
	locations := make(map[string]bool)
	for src, dst := range graph {
		locations[src] = true
		locations[dst] = true
	}

	// Perform DFS from each unvisited location
	for loc := range locations {
		if !visited[loc] {
			ts.topologicalSortUtil(ctx, loc, graph, visited, &stack)
		}
	}

	return stack, nil
}

// validateJournies checks if the input journies are valid
func validateJournies(ctx context.Context, journies [][]string) error {
	if len(journies) == 0 {
		return nil
	}

	for _, journey := range journies {
		if len(journey) != 2 {
			return errors.New("invalid journey format: each journey must have source and destination")
		}
	}
	return nil
}

// topologicalSortUtil performs DFS and builds topological ordering
func (ts *topologicalSorter) topologicalSortUtil(ctx context.Context, current string, graph map[string]string, visited map[string]bool, stack *[]string) {
	visited[current] = true

	if next, exists := graph[current]; exists && !visited[next] {
		ts.topologicalSortUtil(ctx, next, graph, visited, stack)
	}

	*stack = append([]string{current}, *stack...)
}
