module github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/metrics

go 1.24.0

require (
	github.com/stretchr/testify v1.12.1
	go.opentelemetry.io/otel v1.38.0
	go.uber.org/goleak v1.3.0
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
