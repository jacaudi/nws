# Changelog

## [0.2.0](https://github.com/jacaudi/nws/compare/v0.1.0...v0.2.0) (2026-04-26)

* feat!: introduce *Client API for v0.1.0 ([#8](https://github.com/jacaudi/nws/issues/8)) ([d4308b7](https://github.com/jacaudi/nws/commit/d4308b7d212063905582b16a154688130bd2f141)), closes [#3](https://github.com/jacaudi/nws/issues/3) [#4](https://github.com/jacaudi/nws/issues/4) [#7](https://github.com/jacaudi/nws/issues/7) [#6](https://github.com/jacaudi/nws/issues/6) [#2](https://github.com/jacaudi/nws/issues/2) [#4](https://github.com/jacaudi/nws/issues/4) [#7](https://github.com/jacaudi/nws/issues/7) [#6](https://github.com/jacaudi/nws/issues/6) [#2](https://github.com/jacaudi/nws/issues/2) [#2](https://github.com/jacaudi/nws/issues/2) [#3](https://github.com/jacaudi/nws/issues/3) [#4](https://github.com/jacaudi/nws/issues/4) [#6](https://github.com/jacaudi/nws/issues/6) [#7](https://github.com/jacaudi/nws/issues/7) [#5](https://github.com/jacaudi/nws/issues/5)


### BREAKING CHANGES

* nws.Config (type) and Config.SetUserAgent /
Config.SetUnits (methods) are removed. Use nws.NewClient and the
Functional Options (WithUserAgent, WithUnits, etc.) instead.
The package-level wrapper functions are unchanged.

## [0.1.0](https://github.com/jacaudi/nws/compare/v0.0.4...v0.1.0) (2026-04-26)

### Bug Fixes

* **ci:** satisfy linters and skip live-API test ([7321541](https://github.com/jacaudi/nws/commit/7321541f6d7918ff3138b8bc46f9bcc9709fe113))
* **radar:** accept string or number for resolutionVersion ([2fd352c](https://github.com/jacaudi/nws/commit/2fd352c41d776cc970c45d2019f412f2840ebbb4))


### Features

* add cli example for active alerts ([13b3cf8](https://github.com/jacaudi/nws/commit/13b3cf8e5c1f411db11c1b2b99b0eb7fab220cc4))
