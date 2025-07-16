# Silkworm

Silkworm is a WordPress plugin update ticket creation tool. It's meant to bridge the gap between [Platypus](https://github.com/farghul/platypus.git) and [Bowerbird](https://github.com/farghul/bowerbird.git), reading the output from *Platypus* and supplying *Bowerbird* with information recieved and tickets created.

![Silkworm](cocoons.webp)

## Prerequisite

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

A variety of json files to enable ticket creation, authorized querying, and changelog gathering (see `jsons` folder for reference).

## Build

Before building the application, change the values of the `temp`, `tokens`, `repos`, and `config` constants to reflect your environment:

``` go
temp     string = "/data/automation/temp/"
tokens   string = "/data/automation/tokens/"
repos    string = "/data/automation/bitbucket/"
config   string = "desso-automation-conf/atlassian/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac

``` console
go build -o [name] .
```

### Linux

``` console
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

``` console
[program] [flag]
```

## Options

``` console
-h, --help      Help Information
-r, --run       Run Program
-v, --version   Display Program Version
```

## Example: 

``` console
silkworm -r
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/silkworm/blob/main/LICENSE.md "Unlicense Yourself, Set Your Code Free") and is part of the Public Domain.
