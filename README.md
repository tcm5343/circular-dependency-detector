#  circular-dependency-detector
## What?
A GitHub action, containerized and written in Go, to detect simple cycles or topological generations in directed multi-graphs.

## Usage


## Contributing
[Task](https://taskfile.dev/) shall be used in the CI/CD pipelines and for local development. Run `task --list-all` for the list of tasks. For local development, create a `.env` file at the root of the repository to modify your config. The only supported format for the input graph file environment variable (`INPUT_FILE`) is an adjacency list which follows the format used by [NetworkX](https://networkx.org/documentation/stable/reference/readwrite/adjlist.html#).

```shell
user@machine:~/dev/circular-dependency-detector$ cat ./.env 
INPUT_FILE=testing/data/adj_list_no_cycle.txt  # no spaces for now, defaults to {i don't know yet}
```

## Todo (ordered)
* Use graph generated in a workflow and pass it to the container
* Write unit tests
* Error handling
* Logging
* Return 2d slice of nodes instead of integers from topological generations
* Consider edge cases involving multi-graphs in topological generations
* Allow the input file to define nodes as strings instead of just integers (use a set and map)
* Exit early if cycles are identified
* Input other NetworkX output formats

## Musical Acknowledgements
Bob Dylan - Early Mornin' Rain</br>
