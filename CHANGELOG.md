# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.1.0] - 2019-02-06
### Added
- Add a tail newline to output files.
- Introduce [Go modules (vgo)](https://github.com/golang/go/wiki/Modules).

### Fixed
- Use `golang.org/x/cmd/goyacc` instead of `go tool yacc`. It is removed since Go 1.8.
- Meta-fix: re-name the git tag for the version 1.0.0 (v1.0 -> v1.0.0).

## 1.0.0 - 2016-11-18
### Added
- The first release

[Unreleased]: https://github.com/nekketsuuu/lambda2pi/compare/v1.1.0...HEAD
[1.1.0]: https://github.com/nekketsuuu/lambda2pi/compare/v1.0.0...v1.1.0
