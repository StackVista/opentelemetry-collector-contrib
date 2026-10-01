module github.com/open-telemetry/opentelemetry-collector-contrib

// NOTE:
// This go.mod is NOT used to build any official binary.
// To see the builder manifests used for official binaries,
// check https://github.com/open-telemetry/opentelemetry-collector-releases
//
// For the OpenTelemetry Collector Contrib distribution specifically, see
// https://github.com/open-telemetry/opentelemetry-collector-releases/tree/main/distributions/otelcol-contrib

go 1.24

retract (
	v0.76.2
	v0.76.1
	v0.65.0
	v0.37.0 // Contains dependencies on v0.36.0 components, which should have been updated to v0.37.0.
)

replace github.com/ua-parser/uap-go => ./third_party/uap-go

replace github.com/wk8/go-ordered-map/v2 => github.com/StackVista/stackstate-agent/third_party/go-ordered-map/v2 v2.0.0-20261001130553-d6ac30666e85

replace github.com/scaleway/scaleway-sdk-go => ./third_party/scaleway-sdk-go

replace go.opentelemetry.io/otel/schema => ./third_party/otel-schema

replace github.com/aerospike/aerospike-client-go/v8 => ./third_party/aerospike-client-go

replace github.com/DataDog/viper => ./third_party/datadog-viper

replace github.com/AthenZ/athenz => ./third_party/athenz

replace github.com/DataDog/datadog-agent/pkg/util/scrubber => ./third_party/datadog-scrubber

replace github.com/DataDog/datadog-agent/pkg/config/nodetreemodel => ./third_party/datadog-nodetreemodel

replace github.com/DataDog/datadog-agent/pkg/config/setup => ./third_party/datadog-setup

replace github.com/DataDog/datadog-agent/comp/logs/agent/config => ./third_party/datadog-logs-config

replace github.com/prometheus/prometheus => ./third_party/prometheus
