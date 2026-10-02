# Final test-owner CI classification

Tracking: https://github.com/StackVista/stackstate/issues/717

Reviewed terminal source: 406198ac282266cda0fa837f87cebb894cca5bcd.
Native build-and-test36913581937, arm36913581920, scoped36913581910 and
Kubernetes e2e36913581820 all failed. Failures are classified from actual logs:

* checks110543986909: multimod uses direct Go invocations and bypasses VCR's
  module-mode native recipe; original incomplete vendor metadata causes
  inconsistent vendoring. Verification now explicitly uses readonly module
  mode; inventories and original vendor assets remain intact. The subsequent
  native check exposed missing registry entries for the two foreign owners;
  these follow the existing excluded-module entries without assigning Collector
  release versions.
* Linux other lint110543988739: Scaleway's original gomoddirectives gate rejects
  the new local VCR replacement. Select the public checksum-proven immutable
  da917a producer, with exactly the reviewed owner tree and APIs. An owned lint
  config retains all original rules, allowing only the guarded VCR identity.
  Local/unrelated replacements remain rejected; original lint bytes stay exact.
* stable/oldstable other110543988766/110543989257, arm other110544132556 and
  scoped110542278015: selected Ginkgo now rejects common Go parallel/count
  flags. Aerospike native tests retain race/timeout/tag flags, and repeat whole
  invocations twice/three times with count1, without Go parallel. Other modules'
  flags and repeats remain unchanged. No packages/specs are dropped.
* These same other/arm jobs also fail original Aerospike TestMain when connecting
  to localhost:3000 without a server. No server provisioning or pass is claimed.
* Windows other110543988570: original AthenZ syscall.Stat_t and syscall.Stdin
  type errors. Original inventories identify unchanged offending source.
* receiver-0 integration110543987504: Elasticsearch7.16.3 container lacks mapped
  port9200; integration startup never reaches its expected endpoint.
* exporter-3 integration110543987532: RabbitMQ latest rejects deprecated
  transient_nonexcl_queues with Exception541 INTERNAL_ERROR.
* Kubernetes1.30/1.23 jobs110546894835/110546894860: ClusterRBAC and
  NamespacedRBACNoPodIP conditions time out; package exceeds600s.
* Fifteen govuln groups fail on actual selected vulnerable dependencies including
  OTel/grpc, AWS eventstream/CloudWatch, expr and golang.org/x libraries. These gates stay required. A YAML
  backport or unchanged component source is not proof those findings are safe.
* lint/unittest/integration/Kubernetes aggregate jobs propagate failed children;
  they are not separate source defects.

Integration/RBAC source and their go.mod selections are unchanged by the
406198 test-owner increment. This establishes unrelated current holds, not a
claim that every production-baseline job has passed. Golden, telemetrygen,
workflow lint, merge-freeze and Prometheus compliance passed at406198.
Changelog/shellcheck/check-links/Project Tidy startup failures remain policy
holds; no access/allowlist settings or workflows are changed.

Three production executable/OTLP proofs remain at9912672. No production rebuild,
new release or product pin rollout is performed for this classification.
The final session result reports corrected source, real targeted controls and
replacement native CI status. Pending or failed native jobs are never passes.

October2 recovery inspected terminal bbfee340b0 run36917514485 and its arm,
scoped, e2e and telemetrygen siblings. Native checks110557154841 now reach the
missing foreign-owner registry error; Linux other110557156426 rejects remote
VCR replacements too. Both recorded integration corrections are completed in
c975ba529d, preserving original lint bytes and requiring the checksum-anchored
public VCR identity. The original exact go.mod hash also rejects pin changes.

Scoped110555305465 and arm110556235644 have no Ginkgo flag errors; local
subpackages pass and root Aerospike fails localhost:3000. Stable/oldstable other
110557156496/110557156847 fail the same server requirement. Windows other
110557156299 retains original AthenZ syscall errors. RabbitMQ110557155705,
Elasticsearch110557155426 reports repeated scraping errors and receives zero
resources where four are expected, rather than the earlier missing-port failure.
The fifteen govuln groups remain failing gates.

Telemetrygen110555021449 fails fetching unchanged Docker28.3.3 from Go proxy
with HTTP/2 INTERNAL_ERROR: external download failure, not a compiled-source
failure. Supervisor110559677866 HUPReload receives config, sends SIGHUP, then
agent exits1; its five-second remote-config-start condition expires. Source and
module metadata are unchanged by this test-owner increment, but a passing
baseline was not established here: keep an unresolved timing/runtime hold.
Kubernetes110560205263/110560205297 retain RBAC receive-condition timeouts.
Aggregate failures propagate failed children; startup failures without runnable
jobs do not establish a specific policy root cause. Exact replacement-head
native status and targeted controls are returned in the session result.

Native CI36972828772 job110730514864 passes at signedc975ba529d, including
CheckApi, crosslink, tidy/generation/coverage and multimod verification. The
original VCR vendor metadata remains unchanged, and all foreign modules are
registered. Local native multimod verification also passes after recovery.

October2 targeted native Scaleway lint passes with original source/attribution,
license and spelling gates. The owned config differs from original only by the
single VCR identity setting. Isolated gomoddirectives controls accept that remote
identity and reject local/unrelated replacements. The real attribution guard
rejects an altered VCR immutable pin before lint; standalone module verification
passes. Earlier full parser/caller/race/omission suites are preserved receipts,
not rerun claims. Only interrupted metadata/lint integration was requalified.
