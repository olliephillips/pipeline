# Pipeline

## A pipeline for tests, code quality/security review and artifact build

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/olliephillips/pipeline?style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/olliephillips/pipeline/Build?style=flat-square)
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/olliephillips/pipeline/Unit%20Test/master?label=tests&style=flat-square)
![GitHub Release Date](https://img.shields.io/github/release-date/olliephillips/pipeline?style=flat-square)
![GitHub commits since latest release (by date)](https://img.shields.io/github/commits-since/olliephillips/pipeline/latest?style=flat-square)
[![DeepSource](https://deepsource.io/gh/olliephillips/pipeline.svg/?label=active+issues&token=uYY_4Kwjq9MnjT7TzykEyv-J)](https://deepsource.io/gh/olliephillips/pipeline/?ref=repository-badge)
![GitHub](https://img.shields.io/github/license/olliephillips/pipeline?label=license&style=flat-square)

### Unit Tests

Tests are run on all pushes and pull requests to main/master.

### Build

The software is built on each push or pull request to main/master.

### Code Quality & Security

Code quality, security and test coverage is monitored using DeepSource.

### Releases

Releases are built for target OS using GoReleaser. Create a tag, push it. 
A github action will create the release using goreleaser.

```bash
git tag -a v0.0.1 -m "Test release"
git push origin v0.0.1
```

### Badges
Various Github & Deepsource badges are being used

### Todo
- Main branch
- Announce (Goreleaser)

### License
Licensed under [Mozilla Public License 2.0](LICENSE)
