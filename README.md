# Gofermart

Cumulative loyalty system.

[![go report card](https://goreportcard.com/badge/github.com/mobypolo/ya-41-56go?style=flat-square)](https://goreportcard.com/report/github.com/mobypolo/ya-41-56go)
[![test status](https://github.com/mobypolo/ya-41-56go/workflows/gophermart/badge.svg?branch=main "test status")](https://github.com/mobypolo/ya-41-56go/actions)

## Getting Started

Dependencies:

* Go `1.23`
* Docker for PostgreSQL
* Linux or macOS platform

### Startup

To build a gophermart-server, run in the terminal:

```bash
make # in the root directory of the project
```

#### Gophermart

To startup the gophermart (server), run in the terminal:

```bash
./server # in the root directory of the project
```

##### Sartup flags & envs

* `-d` | `DATABASE_DSN` – PostgresSQL DSN
* `-a` | `ADDRESS` – HTTP server address (default `localhost:8080`)

## Development

### Updating the template

To be able to receive updates for autotests and other parts of the template, run the command:

```bash
git remote add -m master template https://github.com/yandex-praktikum/go-musthave-group-diploma-tpl.git
```

To update the autotest code, run the command:

```bash
git fetch template && git checkout template/master .github
```

Then add the changes to your repository.
