# Temporary athenz owning-source backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Logical identity: `github.com/AthenZ/athenz`. Exact checksum-verified selected distribution
`v1.12.13`; module sum `h1:OhZNqZsoBXNrKBJobeUUEirPDnwt0HRo4kQMIO1UwwQ=`; go.mod sum `h1:XXDXXgaQzXaBXnJX6x/bH4yF6eon2lkyzQZ0z/dxprE=`.
Original module ZIP SHA-256: `3859bec0d5f6837dd266d6d3dd1dac83c201317ee8e30df279cb43837ca9634c`.
Complete 3758-file upstream distribution, licenses, copyright, tests
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

Scoped Git attributes retain original asset bytes, including line endings.
