This day was solved very differently from usual.

Instead of coding the solution, I went for using YACC to define the simple language (only multiplication and addition) with custom order of precedences.
This in turn generates two different go files that solve separately part1 and part2. The [Day18/main.go](./main.go) file simply calls both and returns the input

Both the yacc files and the generated ones are used.
The execution time comptued by the solver with `--time` option does not take into account the time needed to generate the go file as that is seen as the final solution. In the future I might add a flag for it.
