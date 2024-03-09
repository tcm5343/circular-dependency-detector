#  circular-dependency-detector
## What?
A GitHub action, containerized and written in Go, to detect simple cycles or topological generations in directed multi-graphs. 

This project was first implemented at the University of Texas at Austin as a group project with [tcm5343](https://github.com/tcm5343), [pabs159](https://github.com/pabs159), and [rory-tatum](https://github.com/rory-tatum). A use case was identified that applied to our industry experience.

## Usage

Below is an example GitHub Action job which utilizes the circular dependency detector. Simply write you graph input file to `/testing/data` and then pass that into the detector. Reminder that the checkout location in an Action is `/github/workspace`.

```yaml
jobs:
  cdd:
    runs-on: ubuntu-latest
    steps:
      - name: create graph input file
        run: |
          echo -e "1 2\n2 3\n3 1" > some-input-graph.txt

      - name: circular-dependency-detector
        uses: tcm5343/circular-dependency-detector@master
        with:
          adjacency_list_path: /github/workspace/some-input-graph.txt
```


## Contributing

### Software

* [Task](https://taskfile.dev/) shall be used in the CI/CD pipelines and for local development for orchestration. Run `task` for the list of tasks. 
* [Podman](https://podman.io/) is used to manage containers in the`Taskfile`.
* [golangci-lint](https://golangci-lint.run/) is used for lint.

### Local Configuration 

For local development, create a `.env` file at the root of the repository to use as your config. The only supported format for the input graph file environment variable (`INPUT_FILE`) is an adjacency list which follows the format used by [NetworkX](https://networkx.org/documentation/stable/reference/readwrite/adjlist.html#). An `.env-example` file exists at the root of the project.

```shell
user@machine:~/dev/circular-dependency-detector$ cat ./.env 
INPUT_FILE=testing/data/adj_list_no_cycle.txt  # no spaces for now, defaults to {i don't know yet}
```

### Roadmap

Future work is identified in [TODO.md](TODO.md).

## Musical Acknowledgements

Bob Dylan - Early Mornin' Rain</br>
