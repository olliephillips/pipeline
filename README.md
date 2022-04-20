# Pipeline

## A repo/pipeline template for tests, code quality/security review and artifact build

Building a template repo using test & build tools with Github Actions. Optimised for Go.

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/olliephillips/pipeline?style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/olliephillips/pipeline/Build?style=flat-square)
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/olliephillips/pipeline/Unit%20Test/master?label=tests&style=flat-square)
![GitHub Release Date](https://img.shields.io/github/release-date/olliephillips/pipeline?style=flat-square)
![GitHub commits since latest release (by date)](https://img.shields.io/github/commits-since/olliephillips/pipeline/latest?style=flat-square)
[![DeepSource](https://deepsource.io/gh/olliephillips/pipeline.svg/?label=active+issues&token=uYY_4Kwjq9MnjT7TzykEyv-J)](https://deepsource.io/gh/olliephillips/pipeline/?ref=repository-badge)
![GitHub](https://img.shields.io/github/license/olliephillips/pipeline?label=license&style=flat-square)

### Unit Tests

Tests are run on 'push' and 'pull request' to main/master.

### Build

Software is built on 'push' and 'pull request' to main/master. Will it build!?

### Code Quality & Security

Code quality, security and test coverage is monitored using [DeepSource](https://deepsource.io).

### Releases

Releases are built for target OS using GoReleaser. Create a tag, push it.
Latest release is [at this generic link here](https://github.com/olliephillips/pipeline/releases/latest)

This is typical is the git create tag & push tag syntax.

```bash
git tag -a v0.0.1 -m "Test release"
git push origin v0.0.1
```

### Branch protection

Within Github, we've configured a rule to prevent direct push to master.  As an admin my commits don't hit that protection.

### Badges
Various Github & Deepsource badges are being used. Mostly done via [Sheilds.io](https://shields.io)

### Announcements
New releases will be announced to Twitter (requires Twitter's elevated developer access).

### Todo
- Main branch

### License
Licensed under [Mozilla Public License 2.0](LICENSE)
