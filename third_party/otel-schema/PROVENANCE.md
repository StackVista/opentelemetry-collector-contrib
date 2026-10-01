# Temporary otel-schema owning-source backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Logical identity: `go.opentelemetry.io/otel/schema`. Exact checksum-verified selected distribution
`v0.0.13`; module sum `h1:gf5AhGzU3V4Ll3xGO+D+Eg2VTMqhz2JlPFr5JpVC/3Q=`; go.mod sum `h1:Ge9lCwrk+B2oYsO2+Tv1R97c+1psVHFs5/4oRTI0OeY=`.
Original module ZIP SHA-256: `6e7bc8ccc25541295dbcb0134ff6afaaec4783d933fc67c0614a824e77c99700`.
Complete 32-file upstream distribution, licenses, copyright, tests
and fixtures are retained and anchored in `UPSTREAM-SHA256SUMS`.
Only matching maintained parser-major imports and required module metadata differ.
No caller API, custom YAML interface, configuration fixture or toolchain contract
is removed. Exact allowed source/metadata hashes are pinned by the native guard;
license/copyright assets have no maintenance exception. Copied upstream CI files
are inactive evidence, not enabled repository workflows.

Every independent consumer and actual generated distribution explicitly selects
this owning source. The native attribution guard fails on missing, changed or
unaccounted assets; owned source retains first-party header/spelling checks.
Original source stays byte-exact through formatting/generation. Remove this
backport after a compatible maintained owning release passes retained upstream
and actual consumer/distribution checks. Product adoption is independent; this
candidate does not establish release or deployment qualification.

Upstream test dependencies use this fork's already selected compatible Testify
1.12.1 line, so independent retained tests do not restore archived parser code.
Go minima remain governed by actual dependencies; no product toolchain bump
is introduced. Original toolchain directives are removed by the existing native
tidy helper without changing the supported minimum.

`UPSTREAM-MODULE.json` records immutable Git origin, original Go checksums and
module ZIP digest without local cache paths. Complete license/source byte
inventories and exact maintenance deltas are enforced by the native guard.
