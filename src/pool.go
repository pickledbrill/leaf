package main

// Pool is the logical isolation for each workspace
type Pool struct {
	Name   string
	TaskID string
}

// createPool creates a "pool"
func createPool(name string) (*Pool, error) {
	pool := Pool{
		Name: name,
	}
	return &pool, nil
}
