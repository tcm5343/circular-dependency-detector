## Todo

* Docker
    * volume mount the code?

* CI/CD
    * add outputs from the action
    * add asserts to e2e tests
    * lint

* Build out `.env` options
    * `LOG_LEVEL` as defined [here](https://pkg.go.dev/golang.org/x/exp/slog#Level), should take both a string or an int value

* Features:
    * Input DOT formatted strings
    * Output graph as PNG
    * `TopologicalGenerationsOf` should return a 2D slice of strings
