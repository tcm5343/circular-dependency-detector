#  circular-dependency-detector
## What?
A GitHub action, containerized and written in Go, to detect simple cycles or topological generations in directed multi-graphs.

## Usage


## Contributing
For local development, creating a `.env` file at the root of the repository to modify your config.
```text
INPUT_FILE="testing/data/adj_list_no_cycle.txt"  # no spaces in path for now, defaults to ... idk
```

[Task](https://taskfile.dev/) shall be used in the CI/CD pipelines and for local development. Run `task --list-all` for an update to date list of tasks.

## Todo
* Return 2d slice of nodes instead of integers from topological generations
* Consider edge cases involving multi-graphs in topological generations
* Allow the input file to define nodes as strings instead of just integers (use a set and map)
* Input the NetworkX adjacency list format
* Write unit tests

## Musical Acknowledgements
Bob Dylan - Early Mornin' Rain
Galt MacDermot - Ripped Open By Metal Explosion
