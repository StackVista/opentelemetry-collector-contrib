# Remaining owner recovery qualification

Tracking: https://github.com/StackVista/stackstate/issues/717

The candidate remains draft PR4 on yaml-maintained-v3. Recovery retained all signed
checkpoints; no product source/access, stable release or deployment was changed.
The original owning decisions and source/checksums remain in
OWNER-QUALIFICATION.md and each UPSTREAM-MODULE.json/UPSTREAM-SHA256SUMS.

A fresh Git archive at 782192818a764812cf3bf54ddbff2539dbb5826e contains all
5,726 original files from the nine controller-reviewed official distributions,
plus all 244 original Scaleway files. The original ZIP SHA-256 values and Git
executable modes of the restored Aerospike/AthenZ/Prometheus distributions were
rechecked. Exact maintenance hashes are enforced separately; original attribution
assets have no exception. Deletion controls cover tools/, dot-YAML, logs, cnf,
configuration JSON and legitimate upstream filenames with spaces.

Native CI36884236939 checked merge8065e7a73c3af900e992d4245dabb69480bdc724.
Its parents are main3d796458171de0f6dbf3d715aa8639428cfa5247 and candidate782192;
its tree3dc55b1e263735db8c914f7c53c39b277eb2fd52 is identical to the candidate.
Setup, CheckApi, crosslink, Collector-version and Linux other lint passed.
Contrib cross-compilation passed all twelve configured platform/architecture
combinations. These jobs do not build the separate testbed/supervisor executables.
The Linux Contrib artifact SHA-256 is
16908a8fe1b405c9a0657f83164d72d4e83628f0216a9da93f8e65b1ab7d1cab.
Its Go metadata selects all reviewed owning replacements and maintained YAML
v2.4.3/v3.0.5; no archived YAML/ghodss module is linked.

CI exposed remaining consumer tidy drift, including two semconv requirements
whose effective version was already selected by the fork's graph. This recovery
records generated tidy output and checks repeat fixed points without changing
the selected component generation or relaxing gates.

Original AthenZ Windows syscall.Stat_t/Stdin type-check failures reproduce before
the parser change. Aerospike's full upstream suite needs localhost:3000 and fails
without that service. Neither is suppressed or misreported as a passing suite.
Vulnerability, integration and action-startup failures remain separate. Complete
owner upstream/caller races and nine original-parser comparisons reported in the
preceding recovery are retained; they are not presented as newly rerun here.

Reproducible serial Linux executable commands (GOWORK=off for independent roots):

```
make genotelcontribcol genoteltestbedcol
GOWORK=off GOMAXPROCS=1 GOMEMLIMIT=3GiB GOGC=30 GOFLAGS=-p=1 CGO_ENABLED=0 \
  go -C cmd/otelcontribcol build -mod=readonly -trimpath -o ../../bin/otelcontribcol_linux_amd64 .
GOWORK=off GOMAXPROCS=1 GOMEMLIMIT=3GiB GOGC=30 GOFLAGS=-p=1 CGO_ENABLED=0 \
  go -C cmd/oteltestbedcol build -mod=readonly -trimpath -o ../../bin/oteltestbedcol_linux_amd64 .
GOWORK=off GOMAXPROCS=1 GOMEMLIMIT=3GiB GOGC=30 GOFLAGS=-p=1 CGO_ENABLED=0 \
  go -C cmd/opampsupervisor build -mod=readonly -trimpath -o /tmp/opampsupervisor .
go version -m <binary>
go tool nm <binary>
sha256sum <binary>
```

The final session result records actual completion, source heads, artifact hashes,
module/symbol graphs and local OTLP UserAgent fixtures; a command recipe alone is
not a passed check. Prior killed builds are failures, not evidence of completion.
Product collectors at aa335e5f69b563bf582f846f8fc3ec31ff9e8ec2 still require separate
builder adoption and binary/image qualification by the controller/product owner.

All three serial Linux executable builds at clean signed 782192 passed with zero
OOM events. The standalone Prometheus owner additionally selects the reviewed
Swag facade0.25.1 and five patched utilities0.27.1, with the required Go1.25
minimum; its earlier isolated graph still selected Swag0.23 archived YAML.
Every affected independent consumer records the necessary minimum. Both full
native tidy passes passed; final fixed-point/binary checks follow this checkpoint.
