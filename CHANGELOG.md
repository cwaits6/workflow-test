# Changelog

## [0.10.0](https://github.com/cwaits6/workflow-test/compare/v0.9.0...v0.10.0) (2026-08-02)


### Features

* probe sticky release-as behaviour ([#21](https://github.com/cwaits6/workflow-test/issues/21)) ([2ac6f95](https://github.com/cwaits6/workflow-test/commit/2ac6f95a69a0a161278d6489c3df27c7d53f2c6b))


### CI/CD

* unpin release-as after the forced 0.9.0 release ([#23](https://github.com/cwaits6/workflow-test/issues/23)) ([280ad9a](https://github.com/cwaits6/workflow-test/commit/280ad9a093014bbbb4f075cf9339a25933d1cdf9))

## [0.9.0](https://github.com/cwaits6/workflow-test/compare/v0.4.0...v0.9.0) (2026-08-02)


### Features

* add a second feature line ([#18](https://github.com/cwaits6/workflow-test/issues/18)) ([15de499](https://github.com/cwaits6/workflow-test/commit/15de4998c318276ef0d1126bdf95aac4ebdc46f1))


### CI/CD

* pin the next release to 0.9.0 via release-as ([#20](https://github.com/cwaits6/workflow-test/issues/20)) ([41df0aa](https://github.com/cwaits6/workflow-test/commit/41df0aad95192733301f1c67458c14d19248c617))

## [0.4.0](https://github.com/cwaits6/workflow-test/compare/v0.3.1...v0.4.0) (2026-08-02)


### ⚠ BREAKING CHANGES

* the greeting line is gone and the version banner prefix changed from workflow-test to workflow-test/v2.

### Features

* rename the binary output prefix and drop the greeting line ([#16](https://github.com/cwaits6/workflow-test/issues/16)) ([ef91189](https://github.com/cwaits6/workflow-test/commit/ef911899527abb9591cdb67e6e2e3d280eb1bc33))

## [0.3.0](https://github.com/cwaits6/workflow-test/compare/v0.2.0...v0.3.0) (2026-08-02)


### Features

* add a validation greeting line ([#12](https://github.com/cwaits6/workflow-test/issues/12)) ([c4f2ac1](https://github.com/cwaits6/workflow-test/commit/c4f2ac18396b9531d0fcc4f353285d019ede47b8))

## [0.3.1](https://github.com/cwaits6/workflow-test/compare/v0.3.0...v0.3.1) (2026-08-02)


### CI/CD

* normalise changelog headings and add config overrides ([#14](https://github.com/cwaits6/workflow-test/issues/14)) ([35a27e3](https://github.com/cwaits6/workflow-test/commit/35a27e3976f4a310fc8955123153c28d4dc0647b))

## [0.2.0](https://github.com/cwaits6/workflow-test/compare/v0.1.6...v0.2.0) (2026-03-24)


### Bug Fixes

* pass App secrets to shared release workflow ([#8](https://github.com/cwaits6/workflow-test/issues/8)) ([f736785](https://github.com/cwaits6/workflow-test/commit/f7367856026d1ce14b6a22493f7fe375cc945474))


### Features

* migrate to GoReleaser for binary builds ([ceaf819](https://github.com/cwaits6/workflow-test/commit/ceaf819ac62d4a9ff5b08ca0b9ee7287a2624fbd))

## [0.1.6](https://github.com/cwaits6/workflow-test/compare/v0.1.5...v0.1.6) (2026-03-16)


### Bug Fixes

* re-trigger release after removing branch protection ([7f9f3a1](https://github.com/cwaits6/workflow-test/commit/7f9f3a13f2a2eacab7d2f3dbbbfcc4006c15ddd3))
* test trivy sbom dependency submission ([#5](https://github.com/cwaits6/workflow-test/issues/5)) ([5d337da](https://github.com/cwaits6/workflow-test/commit/5d337da845313ba619122f0fc352a140c703e24c))
* validate release pipeline with branch protection ([#6](https://github.com/cwaits6/workflow-test/issues/6)) ([e581123](https://github.com/cwaits6/workflow-test/commit/e58112326df87c6fe386ce38d1c32eb72b4562a8))

## [0.1.5](https://github.com/cwaits6/workflow-test/compare/v0.1.4...v0.1.5) (2026-03-16)


### Bug Fixes

* validate updated action versions ([c07b0c2](https://github.com/cwaits6/workflow-test/commit/c07b0c29f1a2bd577cacea4026de763a8b62c7d1))

## [0.1.4](https://github.com/cwaits6/workflow-test/compare/v0.1.3...v0.1.4) (2026-03-16)


### Bug Fixes

* improve version output formatting ([21e0bba](https://github.com/cwaits6/workflow-test/commit/21e0bbadf743679893f346ba1008695ebbba39d8))

## [0.1.3](https://github.com/cwaits6/workflow-test/compare/v0.1.2...v0.1.3) (2026-03-16)


### Bug Fixes

* use correct Wolfi package name go-1.26 ([#4](https://github.com/cwaits6/workflow-test/issues/4)) ([7a8802f](https://github.com/cwaits6/workflow-test/commit/7a8802f4d01641c8330982a3b5a9e66e5a9b08a1))

## [0.1.2](https://github.com/cwaits6/workflow-test/compare/v0.1.1...v0.1.2) (2026-03-16)


### Bug Fixes

* trigger container build on release and push to main ([#3](https://github.com/cwaits6/workflow-test/issues/3)) ([0af8b08](https://github.com/cwaits6/workflow-test/commit/0af8b087124a9ed40cdf66de868400fd0f89a78c))

## [0.1.1](https://github.com/cwaits6/workflow-test/compare/v0.1.0...v0.1.1) (2026-03-16)


### Bug Fixes

* point shared workflow callers at fix branch ([#2](https://github.com/cwaits6/workflow-test/issues/2)) ([07c04e9](https://github.com/cwaits6/workflow-test/commit/07c04e9845bfd1a6db1f379badd820e98a518a0d))

## [0.1.0](https://github.com/cwaits6/workflow-test/compare/v0.0.0...v0.1.0) (2026-03-16)


### Features

* add shared workflow callers and Dockerfile ([#1](https://github.com/cwaits6/workflow-test/issues/1)) ([f1eeed3](https://github.com/cwaits6/workflow-test/commit/f1eeed3111fd61f1f269f58491b46d3ef0a0ae5a))
