# generictypealiasesdemo

[![Build](https://github.com/alrayyes/generictypealiasesdemo/actions/workflows/build.yml/badge.svg)](https://github.com/alrayyes/generictypealiasesdemo/actions/workflows/build.yml)
[![Tests](https://github.com/alrayyes/generictypealiasesdemo/actions/workflows/test.yml/badge.svg)](https://github.com/alrayyes/generictypealiasesdemo/actions/workflows/test.yml)
[![Codecov](https://codecov.io/gh/alrayyes/generictypealiasesdemo/graph/badge.svg?token=LMBZHSBSSD)](https://codecov.io/gh/alrayyes/generictypealiasesdemo)
[![Go Reference](https://pkg.go.dev/badge/github.com/alrayyes/generictypealiasesdemo.svg)](https://pkg.go.dev/github.com/alrayyes/generictypealiasesdemo)

A little demo to learn about [generic type aliases](https://github.com/golang/go/issues/46477).

## What it demonstrates

Before Go 1.24, a type alias couldn't take type parameters. You could alias a
generic type only by fixing its arguments — `type intUser = user.User[int]` —
so an alias never carried the generality of the type it named.

Go 1.24 lifted that. [`main.go`](./main.go) is the whole point in one line:

```go
type newUser[T any] = user.User[T]
```

`newUser` is not a new type. It's another name for `user.User`, parameter and
all, so `newUser[int32]` and `user.User[int32]` are the same type to the
compiler and the methods come along unchanged. The rest of the program builds
that alias three ways — with an integer ID, a string ID and a local struct ID
— to watch the parameter survive being renamed across a package boundary.

[`user`](./user) holds the generic type itself. It's deliberately dull: the
interesting behaviour is in the aliasing, not in anything `User` does.

## Requirements

[Go](https://go.dev/dl/) 1.24+, which is the release that allows this. On 1.23
the alias is a compile error, not a warning.

Working on the demo needs more than running it does — that list, and everything
else about changing it, is in [CONTRIBUTING.md](CONTRIBUTING.md).

## Usage

```shell
go run .
```

```text
Hi, my name is Peter Integer and my ID is 1
Hi, my name is Peter String and my ID is a
Hi, my name is Peter Struct and my ID is {a}
```

## Contributing

Everything about working on this — the tooling, the git hooks, what each linter
is for, and how a release is cut — is in [CONTRIBUTING.md](CONTRIBUTING.md).
Short version: `bun install`, branch, and commit under
[Conventional Commits](https://www.conventionalcommits.org/), because those
subjects pick the next version number.

## Licence

[GPL-3.0-only](./LICENSE)
