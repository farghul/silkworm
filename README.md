# Silkworm

Silkworm is a WordPress plugin update ticket creation tool. It's meant to bridge the gap between [Platypus](https://github.com/farghul/platypus.git) and [Bowerbird](https://github.com/farghul/bowerbird.git), reading the output from *Platypus* and triggering *Bowerbird* with information recieved and tickets created.

![Silkworm](cocoons.webp)

## Prerequisite

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

A `env.json` file containing your API URL and Basic token to enable ticket creation:

``` json
{
  "cloud": "Jira Issue API base URL, ex. https://<domain>/rest/api/latest/",
  "token": "Jira Basic Token",
  "source": "Path to the updates.txt file produced by Platypus",
  "programs": "Root folder for Silkworm and its dependencies"
}
```

A `ticket.json` file containing your API URL and Basic token to enable ticket creation:

``` json
{
  "fields": {
    "assignee": {
      "accountId": "Account ID of user assigned to ticket"
    },
    "issuetype": {
      "self": "IssueType identity url",
      "id": "IssueType ID",
      "name": "IssueType name"
    },
    "labels": [
      "Tag 1",
      "Tag 2"
    ],
    "reporter": {
      "self": "Reporters API identity url",
      "accountId": "Reporters account ID",
      "emailAddress": "Reporters email address"
    },
    "project": {
      "self": "Project identity url",
      "id": "Project ID",
      "key": "Project short name",
      "name": "Project long name",
      "projectTypeKey": "Project Type name"
    },
    "description": "A longer and more detailed explanation of the purpose and goals of the ticket.",
    "summary": "Short, descriptive title for the new ticket.",
    "priority": {
      "self": "Priority identity url",
      "name": "Priority name",
      "id": "Priority ID"
    }
  }
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