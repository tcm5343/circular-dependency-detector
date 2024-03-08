## Todo

* In a workflow, create a input graph file, and execute the action with it
* Determine if project structure is idiomatic
* Determine if `.env` or CLI args are better for configuration
* Build out CI/CD pipeline:
    * Add end to end testing of the program
    * Add linting to the pipeline
    * Add unit tests to the pipeline
* `TopologicalGenerationsOf` should return a 2D slice of strings
* Build out `.env` options
    * `LOG_LEVEL` as defined [here](https://pkg.go.dev/golang.org/x/exp/slog#Level), should take both a string or an int value
* Consider edge cases involving multi-graphs in topological generations
* Input other NetworkX output formats
* Python scripts:
    * To generate input graph files using NetworkX for testing
    * Bring the existing visualizer back into a working state
