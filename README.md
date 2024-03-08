#  circular-dependency-detector
## What?
A GitHub action, containerized and written in Go, to detect simple cycles or topological generations in directed multi-graphs. 

This project was first implemented at the University of Texas at Austin as a group project with [tcm5343](https://github.com/tcm5343), [pabs159](https://github.com/pabs159), and [rory-tatum](https://github.com/rory-tatum). A use case was identified that applied to our industry experience.

## Usage

Below is an example GitHub Action job which utilizes the circular dependency detector. Simply write you graph input file to `/testing/data` and then pass that into the detector.

```yaml
jobs:
  cdd:
    runs-on: ubuntu-latest
    steps:
      - name: checkout cdd
        uses: actions/checkout@v4
        with:
          repository: ${{ github.repository }}
          ref: ${{ github.GITHUB_REF_NAME }}  # this should be a working tag

      - name: install task
        uses: arduino/setup-task@v2
        with:
          repo-token: ${{ secrets.GITHUB_TOKEN }}
      
      - name: create graph input file
        run: |
          echo -e "1 2\n2 3\n3 1" > testing/data/some-input-graph.txt

      - name: circular dependency detector
        run: |
          task INPUT_FILE="testing/data/some-input-graph.txt" run
```


## Contributing

### Software

* [Task](https://taskfile.dev/) shall be used in the CI/CD pipelines and for local development for orchestration. Run `task` for the list of tasks. 
* [Podman](https://podman.io/) is used to manage containers in the`Taskfile`.
* [golangci-lint](https://golangci-lint.run/) is used for lint.

### Configuration 

For local development, create a `.env` file at the root of the repository to modify your config. The only supported format for the input graph file environment variable (`INPUT_FILE`) is an adjacency list which follows the format used by [NetworkX](https://networkx.org/documentation/stable/reference/readwrite/adjlist.html#). An `.env-example` file exists at the root of the project.

```shell
user@machine:~/dev/circular-dependency-detector$ cat ./.env 
INPUT_FILE=testing/data/adj_list_no_cycle.txt  # no spaces for now, defaults to {i don't know yet}
```

## Musical Acknowledgements

Bob Dylan - Early Mornin' Rain</br>
