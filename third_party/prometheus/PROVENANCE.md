# Temporary prometheus owning-source backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Logical identity: `github.com/prometheus/prometheus`. Exact checksum-verified selected distribution
`v0.307.1`; module sum `h1:Hh3kRMFn+xpQGLe/bR6qpUfW4GXQO0spuYeY7f2JZs4=`; go.mod sum `h1:/7YQG/jOLg7ktxGritmdkZvezE1fa6aWDj0MGDIZvcY=`.
Original module ZIP SHA-256: `4799ef4948765e648d99ad62af90d0ebbeca62fb3ada506167d5d19c80f8bc34`.
Complete 1457-file upstream distribution, licenses, copyright, tests
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

The original upstream Makefile is retained byte-for-byte. Owned GNUmakefile
provides native integration without replacing the original build artifact.

Upstream test dependencies use this fork's already selected compatible Testify
1.12.1 line, so independent retained tests do not restore archived parser code.
Go minima remain governed by actual dependencies; no product toolchain bump
is introduced. Original toolchain directives are removed by the existing native
tidy helper without changing the supported minimum.

Scoped Git attributes retain original asset bytes, including line endings.

`UPSTREAM-MODULE.json` records immutable Git origin, original Go checksums and
module ZIP digest without local cache paths. Complete license/source byte
inventories and exact maintenance deltas are enforced by the native guard.
