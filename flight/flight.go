package flight

import (
	"context"
	"errors"
)

func (flightJourney *flightJourney) GetJourneyOrder(ctx context.Context, journies [][]string) ([]string, error) {
	if err := validateJournies(ctx, journies); err != nil {
		return nil, err
	}

	if len(journies) == 0 {
		return nil, nil
	}

	// Build graph
	graph := make(map[string]string)
	for _, journey := range journies {
		source, destination := journey[0], journey[1]
		if _, exists := graph[source]; exists {
			return nil, errors.New("invalid journey sequence: location appears multiple times as source")
		}
		graph[source] = destination
	}

	sorter := newSortService()
	// Get ordered path using topological sort
	result, err := sorter.Sort(ctx, graph)
	if err != nil {
		return nil, err
	}

	// Validate the path length matches input journies
	if len(result) != len(journies)+1 {
		return nil, errors.New("invalid journey sequence: disconnected paths")
	}

	// Validate the path follows the graph edges
	for i := 0; i < len(result)-1; i++ {
		if graph[result[i]] != result[i+1] {
			return nil, errors.New("invalid journey sequence: no valid path exists")
		}
	}

	return result, nil
}
