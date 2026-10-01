# Temporary datadog-logs-config owning-source backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Logical identity: `github.com/DataDog/datadog-agent/comp/logs/agent/config`. Exact checksum-verified selected distribution
`v0.73.0-devel.0.20251030121902-cd89eab046d6`; module sum `h1:8LTxkJdRLdcZKeMk96Hpu5X1qR6h7maoiopNUzGxbd8=`; go.mod sum `h1:IVuf9cLGibnYuHnthQNcYoJk18tHUw/eL2d/IWcilR0=`.
Original module ZIP SHA-256: `16acb3f9277a4f2378a136fd6517eba717441cd627eaa9fe4dc25f1108f80a89`.
Complete 20-file upstream distribution, licenses, copyright, tests
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

Original monorepo-relative sibling replacements are removed because the complete
published independent module distribution does not include those siblings.
Published owning requirements remain; explicit consumer selections do not rely
on dependency replacements propagating. Upstream immutable source is
`cd89eab046d6`; no DataDog generation/API upgrade is applied.

Native tests/lint select upstream `test` helpers via their documented build tag.
Production builds do not set this tag. Original source bytes are unchanged.

Independent published requirements use the actual selected-generation support
modules from the Contrib distribution, checksum-verified at cd89eab046d6.
This replaces monorepo-only relative selections and stale/unpublished direct
requirements without upgrading the DataDog API generation. Exact dependency
versions and checksums are in the owned go.mod/go.sum and pinned policy.

Upstream test dependencies use this fork's already selected compatible Testify
1.12.1 line, so independent retained tests do not restore archived parser code.
Go minima remain governed by actual dependencies; no product toolchain bump
is introduced. Original toolchain directives are removed by the existing native
tidy helper without changing the supported minimum.

`UPSTREAM-MODULE.json` records immutable Git origin, original Go checksums and
module ZIP digest without local cache paths. Complete license/source byte
inventories and exact maintenance deltas are enforced by the native guard.
