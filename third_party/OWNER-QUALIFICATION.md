# Maintained YAML owner qualification

Tracking: https://github.com/StackVista/stackstate/issues/717

The fork's actual generated Contrib and testbed distributions determine runtime
consumption. Component presence is not SUSE product adoption. The reviewed
UAParser source and attribution gate remain in this candidate's ancestry.

| Logical owner | Retained source | Owning-release decision |
| --- | --- | --- |
| wk8/go-ordered-map/v2 | v2.1.8, shared Agent d6ac30666e85 | Reuse reviewed public addressable owner; no duplicate source |
| scaleway-sdk-go | beta35, 1fc11916a520 | Private Toolbox source has no qualified Contrib CI path; authorized public original-only backport |
| DataDog/viper | v1.14.1 pseudo, b33ffa9792d9 | Shared v1.14.0 changes selected mapstructure/HCL contracts; maintained v1.15.1 changes nil override lookup, so retain exact selected source |
| AthenZ/athenz | v1.12.13 | First maintained owning release v1.12.43 requires Go 1.26.2; retain supported selected source |
| otel/schema | v0.0.13 | Latest v0.0.19 still imports archived v3; exact matching-major backport |
| aerospike-client-go/v8 | v8.4.0 | Latest v8.9.0 still imports archived v3; preserve config/custom YAML interfaces |
| prometheus/prometheus | v0.307.1 | Maintained v0.310.0 selects Collector component/consumer/pdata/processor v1.51.0 and YAML v4 prerelease; avoid a generation-wide API/dependency change |
| DataDog scrubber, nodetreemodel, setup, logs config | selected cd89eab046d6006660af4709889e5cebc65806b7 | Preserve selected generation; published support versions match actual Contrib graph |

Each local owner contains original module identity, complete distribution,
UPSTREAM-MODULE.json, original byte inventory and exact permitted maintenance
hashes. Original licenses/copyright cannot be waived. Native integration checks
owned Go/shell headers and documents. All first-party selectors/gates remain;
foreign owners are registered for native checks and explicitly tidied in
third_party/tidy-order.txt before crosslink's first-party order.

Root and every requiring independent module explicitly select the owners.
Generated OCB configs include both actual local module identities and root's
reviewed remote versioned replacements. No dependency replacement is assumed to
propagate. Original old logical owner versions with maintained replacements are
not archived runtime implementations. Version prefixes of commit-derived nested
module pseudo-versions do not downgrade their checksum-anchored original source.

The three inherited internal/datadog invalid tags now select authoritative
published modules from the same cd89 generation; PUBLISHED-MODULES.md records
checksums. Remaining unrelated vulnerability, RabbitMQ, Elasticsearch, Kubernetes
RBAC, Windows scoped tests and action startup failures are not suppressed.
Aerospike's root upstream suite requires a live localhost server: relevant config
fixtures are tested separately without manufacturing a server pass. DataDog
retained suites require upstream's test build tag; production does not use it.

Product builders require explicit adoption of independently reviewed source,
actual binaries/images and delivery verification. No fork result establishes
released/deployed product completion. Remove temporary backports after compatible
maintained owning releases pass original and actual consumer/distribution checks.
