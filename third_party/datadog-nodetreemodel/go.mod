module github.com/DataDog/datadog-agent/pkg/config/nodetreemodel

go 1.24.0

// Internal deps fix version
replace github.com/spf13/cast => github.com/DataDog/cast v1.8.0

require (
	github.com/DataDog/datadog-agent/pkg/config/model v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/config/viperconfig v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/util/log v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/go-viper/mapstructure/v2 v2.4.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/spf13/cast v1.10.0
	github.com/stretchr/testify v1.12.1
	go.uber.org/atomic v1.11.0
	golang.org/x/exp v0.0.0-20251009144603-d2f985daa21b
)

require go.yaml.in/yaml/v3 v3.0.5 // indirect

require (
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/version v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/viper v1.14.1-0.20251008075154-b33ffa9792d9 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/magiconair/properties v1.8.10 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	go.yaml.in/yaml/v2 v2.4.3
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/time v0.14.0 // indirect
)

// This section was automatically added by 'dda inv modules.add-all-replace' command, do not edit manually

replace github.com/DataDog/datadog-agent/pkg/util/scrubber => ../datadog-scrubber

replace github.com/DataDog/viper => ../datadog-viper
