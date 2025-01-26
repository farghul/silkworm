# Silkworm

Silkworm is a WordPress plugin update ticket creation tool. It's meant to bridge the gap between [Platypus](https://github.com/farghul/platypus.git) and [Bowerbird](https://github.com/farghul/bowerbird.git), reading the output from *Platypus* and supplying *Bowerbird* with information recieved and tickets created.

![Silkworm](cocoons.webp)

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

A variety of json files to enable ticket creation, authorized querying, and changelog gathering (see `jsons` folder for reference).

## Build

Before building the application, change the value of the `assets` constant to reflect your environment:

``` go
assets string = "/data/scripts/automation/assets/"
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
-v, --version   Display Program Version
```

## Example: 

``` console
silkworm
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/silkworm/blob/main/LICENSE.md) and is part of the Public Domain.