# AoC2020
AoC2020 Solves - in Go! This year the solves will most likely be late!

## Structure
The repository is structured as follwos:

- [pkg](./pkg) directory:
  - One package, named Day##, for the ## day of the month (e.g. [Day01](./pkg/Day01))
  - Every package presents the same interface: one exported method `func Solve(input string) (int,int)` which provides the answers for part 1 and part 2
    - (note, Day08 returns, at the moment, `(int64,int64)`, but that will change as soon as I care)
  - Auxiliary packages for visualization purposes and to gather common functionalities (or those that I believe will become shared in the future)
- [cmd](./cmd) directory:
  - [solver.go](./cmd/solver.go) is the main entrance point, it can be compiled into an executable that can either execute a specific day, or present the execution time of all days
  - [solver_test.go](./cmd/solver_test.go) provides unit testing to verify that refactoring doesn't break old days
    - *NOTE*: these are NOT unit tests on the examples provided in the puzzle. I fill them after solving the day to make sure I do not break old days when refactoring.


## Execution
To execute unit tests, simply run `go test ./...` from the `root` directory (note: most packages don't have tests so far, so you might just care about running it inside `cmd`)

To execute the solver, go into the `cmd` directory and build the executable: `go build solver.go`. This will create an executable called `solver`, which can be either used to solve a single day or to time the execution times of all days (note: visualization is also possible, but enabled only for few days- nonexisting or broken for the others)

### To execute a single day:

`./solver --day XX`

Example:

```
 ./solver --day 12
2020/12/12 21:36:34 input.ReadInput took 57.068µs
2020/12/12 21:36:34 Day12.Solve took 59.761µs
Part 1: 1424, Part 2:63447
```

### To test the execution time

```
2020/12/12 21:37:15 input.ReadInput took 42.736µs
2020/12/12 21:37:15 input.ReadInput took 44.901µs
2020/12/12 21:37:15 input.ReadInput took 34.208µs
2020/12/12 21:37:15 input.ReadInput took 57.248µs
2020/12/12 21:37:15 input.ReadInput took 49.373µs
2020/12/12 21:37:15 input.ReadInput took 49.267µs
2020/12/12 21:37:15 input.ReadInput took 95.57µs
2020/12/12 21:37:15 input.ReadInput took 41.62µs
2020/12/12 21:37:15 input.ReadInput took 42.685µs
2020/12/12 21:37:15 input.ReadInput took 42.908µs
2020/12/12 21:37:15 input.ReadInput took 47.453µs
2020/12/12 21:37:15 input.ReadInput took 53.19µs
2020/12/12 21:37:15 input.ReadInput took 39.559µs
2020/12/12 21:37:15 input.ReadInput took 38.081µs
2020/12/12 21:37:15 input.ReadInput took 29.059µs
2020/12/12 21:37:15 input.ReadInput took 25.396µs
2020/12/12 21:37:15 input.ReadInput took 31.637µs
2020/12/12 21:37:15 input.ReadInput took 35.344µs
2020/12/12 21:37:15 input.ReadInput took 30.806µs
2020/12/12 21:37:15 input.ReadInput took 25.045µs
2020/12/12 21:37:15 input.ReadInput took 22.865µs
2020/12/12 21:37:15 input.ReadInput took 23.192µs
2020/12/12 21:37:15 input.ReadInput took 28.213µs
2020/12/12 21:37:15 input.ReadInput took 38.441µs
2020/12/12 21:37:15 input.ReadInput took 40.461µs
2020/12/12 21:37:15 main.readAllInputs took 1.304148ms
2020/12/12 21:37:15 Day01.Solve took 33.287µs
2020/12/12 21:37:15 Day02.Solve took 249.415µs
2020/12/12 21:37:15 Day03.Solve took 37.676µs
2020/12/12 21:37:15 Day04.Solve took 572.614µs
2020/12/12 21:37:15 Day05.Solve took 529.631µs
2020/12/12 21:37:15 Day06.Solve took 1.093418ms
2020/12/12 21:37:15 Day07.Solve took 3.9658ms
2020/12/12 21:37:15 Day08.Solve took 15.741141ms
2020/12/12 21:37:15 Day09.Solve took 2.347936ms
2020/12/12 21:37:15 Day10.Solve took 19.757µs
2020/12/12 21:37:15 Day11.Solve took 41.333085ms
2020/12/12 21:37:15 Day12.Solve took 41.835µs
2020/12/12 21:37:15 Day13.Solve took 48ns
2020/12/12 21:37:15 Day14.Solve took 39ns
2020/12/12 21:37:15 Day15.Solve took 40ns
2020/12/12 21:37:15 Day16.Solve took 39ns
2020/12/12 21:37:15 Day17.Solve took 39ns
2020/12/12 21:37:15 Day18.Solve took 38ns
2020/12/12 21:37:15 Day19.Solve took 39ns
2020/12/12 21:37:15 Day20.Solve took 39ns
2020/12/12 21:37:15 Day21.Solve took 38ns
2020/12/12 21:37:15 Day22.Solve took 39ns
2020/12/12 21:37:15 Day23.Solve took 38ns
2020/12/12 21:37:15 Day24.Solve took 38ns
2020/12/12 21:37:15 Day25.Solve took 41ns
2020/12/12 21:37:15 main.timeSolves took 67.457204ms
```

This will report the execution times of all the file reading and then of every single day (note: it automatically executes up to day25, even if only the stub is present, but the impact on execution time is negligible, less than `50ns` for each day)

#### Inputs
The program expects input files to be in an `/inputs` folder, each day named as `X.txt` (e.g. `1.txt`, `15.txt`, please note the absence of a leading zero here)
