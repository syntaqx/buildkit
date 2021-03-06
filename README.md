<img src="assets/og-social-preview.jpg">

[![Go Report Card](https://goreportcard.com/badge/github.com/syntaqx/buildkit)](https://goreportcard.com/report/github.com/syntaqx/buildkit)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/syntaqx/buildkit)](https://pkg.go.dev/github.com/syntaqx/buildkit)
[![codecov](https://codecov.io/gh/syntaqx/buildkit/branch/main/graph/badge.svg?token=5aj7H1Xrvz)](https://codecov.io/gh/syntaqx/buildkit)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/syntaqx/buildkit?sort=semver)](https://hub.docker.com/r/syntaqx/buildkit)

BuildKit is a backend boilerplate that defines an opinionated API that
implements reusable core services for common entities.

## Service endpoints

- [API](http://localhost:8080/)
- [Jaeger UI](http://localhost:16686/)

## Docker build specifications

| Event          | Ref                    | Commit SHA | Docker Tags                       |
|----------------|------------------------|------------|-----------------------------------|
| `schedule`     | `refs/heads/main`      | `45f132a`  | `sha-45f132a`, `nightly`          |
| `push`         | `refs/heads/main`      | `cf20257`  | `sha-cf20257`, `edge`             |
| `push`         | `refs/heads/my/branch` | `a5df687`  | `sha-a5df687`, `my-branch`        |
| `push tag`     | `refs/tags/v1.2.3`     | `bf4565b`  | `sha-bf4565b`, `v1.2.3`, `latest` |

## License

[MIT]: https://syntaqx.mit-license.org

`buildkit` is open source software released under the [MIT license][MIT].
