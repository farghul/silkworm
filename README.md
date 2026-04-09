# Silkworm

Silkworm is a WordPress plugin update ticket creation tool. It's meant to bridge the gap between [Platypus](https://github.com/farghul/platypus.git) and [Bowerbird](https://github.com/farghul/bowerbird.git), reading the output from *Platypus* and supplying *Bowerbird* with information recieved and tickets created.

![Silkworm](cocoons.webp)

## 📚 Prerequisite

The [Go Programming Language](https://go.dev "Build simple, secure, scalable systems with Go") installed to enable building executables from source code.

A variety of json files to enable ticket creation, authorized querying, and changelog gathering (see `jsons` folder for reference).

## 🚧 Build

Before building the application, change the values of these constants to reflect your environment:

``` go
meta     string = "/data/automation/jsons/"
temp     string = "/data/automation/temp/"
tokens   string = "/data/automation/tokens/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac

``` zsh
go build -o [name] .
```

### Linux

``` zsh
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## 🏃 Run

``` zsh
silkworm -r
```

## 🎏 Available Flags

| Command               | Action                      |
|:----------------------|:----------------------------|
|    `-h, --help`       |   Help information          |
|    `-r, --run`        |   Run program               |
|    `-v, --version`    |   Display program version   |

## 🎫 License

Code is distributed under [The Unlicense](https://github.com/farghul/silkworm/blob/main/LICENSE.md "Unlicense Yourself, Set Your Code Free") and is part of the Public Domain.
