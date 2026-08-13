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
that alias three ways — an `int32` ID, a `string` ID and a local struct ID —
to watch the parameter survive being renamed across a package boundary.

[`user`](./user) holds the generic type itself. It's deliberately dull: the
interesting behaviour is in the aliasing, not in anything `User` does.

## Requirements

- [Go](https://go.dev/dl/) 1.24+, which is the release that allows this. On
  1.23 the alias is a compile error, not a warning
- [Bun](https://bun.sh) for the JS tooling: Biome, Prettier, markdownlint-cli2,
  commitlint and lefthook
- [Docker](https://www.docker.com/), used by lefthook to run hadolint against
  the Dockerfiles
- [yamllint](https://yamllint.readthedocs.io),
  [actionlint](https://github.com/rhysd/actionlint),
  [typos](https://github.com/crate-ci/typos),
  [GoReleaser](https://goreleaser.com) and [Vale](https://vale.sh), run by the
  git hooks:

  ```shell
  go install github.com/rhysd/actionlint/cmd/actionlint@v1.7.12
  go install github.com/goreleaser/goreleaser/v2@v2.17.1
  go install github.com/errata-ai/vale/v3/cmd/vale@v3.17.1  # needs Go 1.25.7+
  pipx install yamllint==1.38.0
  cargo install typos-cli --version 1.49.0
  vale sync   # fetches the style packages Vale needs
  ```

  These five are optional. `bun install` doesn't provide them, so the hook
  skips any that aren't on your `PATH` rather than failing the push. You can
  clone and contribute without installing all five; CI runs them
  unconditionally, so the check isn't bypassed, it just lands later.

## Usage

```shell
go run .
```

```text
Hi, my name is Peter Integer and my ID is 1
Hi, my name is Peter String and my ID is a
Hi, my name is Peter Struct and my ID is {a}
```

## Development

Clone the repo and install dependencies:

```shell
bun install
```

That also installs the [lefthook](https://lefthook.dev) git hooks
(`pre-commit`, `pre-push`, `commit-msg`) defined in
[`lefthook.yml`](./lefthook.yml). `pre-commit` fixes what a tool can settle on
its own and restages it. `pre-push` never writes — it runs the same checks in
report mode over the whole tree.

### Building

```shell
go build -v ./...
```

### Testing

```shell
go test -cover ./...
```

### Linting

```shell
golangci-lint run                        # Go
bun biome:lint .                         # JSON formatting and key ordering
bun prettier:lint "**/*.{md,yml,yaml}"   # Markdown and YAML layout
bun markdown:lint                        # Markdown structure
yamllint --strict .                      # YAML style
actionlint                               # GitHub Actions workflows
typos                                    # spelling, everywhere
goreleaser check                         # release config
vale --glob='!{node_modules/**,styles/**,dist/**,CHANGELOG.md}' .
```

[Biome](https://biomejs.dev) is configured in [`biome.json`](./biome.json) and
owns every file type it supports, including the JSON key sorting that used to
come from `prettier-plugin-sort-json`. That lives on as Biome's `useSortedKeys`
assist, switched off for `package.json` so its conventional key order survives.

[Prettier](https://prettier.io) fills the two gaps Biome leaves: Markdown and
YAML. It owns table alignment, so don't line a table up by hand — it will just
redo it. `proseWrap` is `preserve`, so your line breaks stay where you put
them. The `pre-commit` hook runs Prettier first and markdownlint second, and
the markdownlint rules with an opinion about layout (MD004, MD007, MD012,
MD049, MD050) are off in
[`.markdownlint-cli2.jsonc`](./.markdownlint-cli2.jsonc), so the two can't undo
each other. What Prettier doesn't touch is in
[`.prettierignore`](./.prettierignore).

Prose gets a tier above layout. [Vale](https://vale.sh) is the style one —
house voice, weasel words, wordiness — configured in [`.vale.ini`](./.vale.ini)
with the Google and proselint packages. `vale sync` fetches those into
`styles/`, which is why `styles/Google/` and `styles/proselint/` are ignored by
git while the House vocabulary beside them is committed. Product names and
jargon go in that vocabulary rather than in `<!-- vale off -->` comments
scattered through the prose, and `Vale.Terms` then holds the document to the
spelling it records.

Vale mostly warns. Style advice that blocks a merge teaches people
`--no-verify`, which costs more than it catches. Errors still fail, and a term
spelled against the vocabulary is an error.

Above Vale sits the tier that has right answers.
[ltex-cli-plus](https://github.com/ltex-plus/ltex-ls-plus) wraps LanguageTool
for grammar, spelling, and punctuation. It fails the build where Vale only
warns, and it runs in CI only, in [`prose.yml`](./.github/workflows/prose.yml):
it's a ~300 MB download shipping its own Java runtime, which is more than a
commit should wait on. Configure it in [`.ltex.json`](./.ltex.json), where
`PASSIVE_VOICE` is off because Vale already flags passive voice and two tools
underlining the same sentence is how a team learns to ignore both.

YAML gets the same split as Markdown: Prettier lays it out and
[yamllint](https://yamllint.readthedocs.io) judges what Prettier produced (see
[`.yamllint.yml`](./.yamllint.yml)). yamllint wants two spaces before an inline
comment and Prettier collapses that to exactly one, so `min-spaces-from-content`
stays lowered to 1. Insist on two and the formatter takes the second one
straight back off.

Dockerfiles are linted with
[hadolint](https://github.com/hadolint/hadolint) through `docker compose` (see
[`docker-compose.yml`](./docker-compose.yml)).

[gitleaks](https://github.com/gitleaks/gitleaks) scans the full history for
committed secrets. It runs in CI only — it needs an unshallow clone, which is
too slow for a git hook.

[typos](https://github.com/crate-ci/typos) skips `CHANGELOG.md` (see
[`.typos.toml`](./.typos.toml)). release-please writes that file out of commit
subjects, and a misspelling in a released commit is history rather than a
defect.

### Commit messages

Commits follow [Conventional Commits](https://www.conventionalcommits.org/),
enforced by the `commit-msg` hook and again in CI. Use the prompt rather than
writing them by hand:

```shell
bun commit
```

## Releases

The [release-please](https://github.com/googleapis/release-please) action opens
release pull requests from the Conventional Commits history. Merging one tags
the release and cuts it, and a second job in the same workflow runs
[GoReleaser](https://goreleaser.com/) (see
[`.goreleaser.yaml`](./.goreleaser.yaml)) to hang the binaries on it. Both live
in [`release-please.yml`](./.github/workflows/release-please.yml): a tag pushed
by an action doesn't reliably start a workflow of its own, so a separate on-tag
workflow is how you end up with a release and no binaries.

That workflow also takes a manual run with a tag as input, for when a release
ends up tagged but empty. release-please creates the release and then reports
it, so a failure between the two leaves GoReleaser gated on an output that
never arrived.

The action runs in manifest mode:
[`release-please-config.json`](./release-please-config.json) says how to
release and
[`.release-please-manifest.json`](./.release-please-manifest.json) holds the
version we're on. That manifest is the file to correct by hand when a release
goes wrong. Don't move the tag.

**Pull requests are rebased, never merged.** GitHub writes the pull request
title into the body of a merge commit, and release-please reads conventional
commits out of a commit body as well as its subject, so every change landed by
a merge commit reached the changelog twice. Merge commits are off in the
repository settings, and no combination of GitHub's merge commit settings
leaves that body empty.

## Dependency updates

[Dependabot](./.github/dependabot.yml) keeps Go, Bun, Docker, and GitHub Actions
dependencies current. Pull requests that pass CI merge automatically (see
[`auto-merge-dependabot.yml`](./.github/workflows/auto-merge-dependabot.yml)).

## License

[GPL-3.0-only](./LICENSE)
