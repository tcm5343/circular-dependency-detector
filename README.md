#  circular-dependency-detector
[![Go Report Card](https://goreportcard.com/badge/github.com/tcm5343/circular-dependency-detector)](https://goreportcard.com/report/github.com/tcm5343/circular-dependency-detector)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
## What?
A GitHub action, containerized and written in Go, to detect simple cycles or topological generations in directed multi-graphs. While this repo has utility, it is not novel and serves primarily pedagogical purposes. 

This project was first implemented at the University of Texas at Austin as a group project with [tcm5343](https://github.com/tcm5343), [pabs159](https://github.com/pabs159), and [rory-tatum](https://github.com/rory-tatum). A use case was identified that applied to our industry experience.

## Usage


## Contributing
[Task](https://taskfile.dev/) shall be used in the CI/CD pipelines and for local development. Run `task --list-all` for the list of tasks. Lint shall be performed using [golangci-lint](https://golangci-lint.run/).

For local development, create a `.env` file at the root of the repository to modify your config. The only supported format for the input graph file environment variable (`INPUT_FILE`) is an adjacency list which follows the format used by [NetworkX](https://networkx.org/documentation/stable/reference/readwrite/adjlist.html#). An `.env-example` file exists at the root of the project.

```shell
user@machine:~/dev/circular-dependency-detector$ cat ./.env 
INPUT_FILE=testing/data/adj_list_no_cycle.txt  # no spaces for now, defaults to {i don't know yet}
```

## Todo
* In a workflow, create a input graph file, and execute the action with it
* Determine if project structure is idiomatic
* Determine if `.env` or CLI args are better for configuration
* Create build pipeline:
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

## Musical Acknowledgements
Bob Dylan - Early Mornin' Rain</br>
