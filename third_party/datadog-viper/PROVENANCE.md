# Temporary datadog-viper owning-source backport

Tracking: https://github.com/StackVista/stackstate/issues/717

Logical identity: `github.com/DataDog/viper`. Exact checksum-verified selected distribution
`v1.14.1-0.20251008075154-b33ffa9792d9`; module sum `h1:uX4ZqokylxBI67r+AN09CiFY+s8Frmf33YZlyqIMBc4=`; go.mod sum `h1:QGomve/3EbYfi58jADS97U2OKfsxqh2pWemuT0azbdk=`.
Original module ZIP SHA-256: `de2a7acf30264518464ecb5396ce7aa4350bb69e669a25a80db9bd7244b075c0`.
Complete 18-file upstream distribution, licenses, copyright, tests
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

Original upstream Makefile is retained byte-for-byte; owned GNUmakefile
provides native integration without replacing that original asset.
