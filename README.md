# Silkworm

Silkworm is a Jira ticket creation tool for WordPress plugin updates. It's meant to bridge the gap between [Platypus](https://github.com/farghul/platypus.git) and [Bowerbird](https://github.com/farghul/bowerbird.git), reading the output from *Platypus* and supplying *Bowerbird* with information via newly created tickets.

![Silkworm](cocoons.webp)

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

A `jsons/jira.json` file containing your API URL and Basic token to enable ticket creation:

``` json
{
    "base": "Jira Issue base URL",
    "path": "Path to Silkworm application",
    "source": "Location of the updates.txt file",
    "token": "Email:Jira API Token combination with Base 64 Encoding"
}
```

## Build

From the root folder containing `main.go`, use the command that matches your environment:

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
-c, --create    Run the main program
-h, --help      Help Information
-v, --version   Display Program Version
```

## Example: 

``` console
silkworm -c
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/silkworm/blob/main/LICENSE.md) and is part of the Public Domain.