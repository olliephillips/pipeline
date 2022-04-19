# Pipeline

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/olliephillips/pipeline?style=for-the-badge)
![GitHub Workflow Status (branch)](https://img.shields.io/github/workflow/status/olliephillips/pipeline/Unit%20Test/master?label=Tests%20(master)&style=for-the-badge)
![GitHub Release Date](https://img.shields.io/github/release-date/olliephillips/pipeline?style=for-the-badge)
![GitHub commits since latest release (by date)](https://img.shields.io/github/commits-since/olliephillips/pipeline/latest?style=for-the-badge)


A pipeline for tests and artifact build.

### Unit Tests

Tests are run on all pushes and pull requests to main/master

### Releases
Create a tag, push it. A github action will create the release using goreleaser.

```bash
git tag -a v0.0.1 -m "Test release"
git push origin v0.0.1
```

### Badges


### Todo
- Main branch
- Tests workflow
- Badges
- Announce (Goreleaser)
