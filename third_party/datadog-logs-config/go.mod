module github.com/DataDog/datadog-agent/comp/logs/agent/config

go 1.24.0

require (
	github.com/DataDog/datadog-agent/comp/core/config v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/config/model v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/config/setup v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/config/structure v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/config/utils v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/logs/types v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/util/log v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/util/pointer v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.73.0-devel.0.20251030121902-cd89eab046d6
	github.com/stretchr/testify v1.12.1
	go.uber.org/atomic v1.11.0
)

require go.yaml.in/yaml/v2 v2.4.3 // indirect

require (
	github.com/DataDog/datadog-agent/comp/api/api/def v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/comp/core/flare/builder v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/comp/core/flare/types v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/comp/core/secrets/def v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/comp/core/secrets/noop-impl v0.0.0-20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/comp/core/status v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/comp/def v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/collector/check/defaults v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/config/create v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/config/env v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/config/mock v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/config/nodetreemodel v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/config/teeconfig v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/config/viperconfig v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/fips v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/template v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/executable v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/filesystem v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/fxutil v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/option v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/system v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/system/socket v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/util/winutil v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/datadog-agent/pkg/version v0.73.0-devel.0.20251030121902-cd89eab046d6 // indirect
	github.com/DataDog/viper v1.14.1-0.20251008075154-b33ffa9792d9 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/ebitengine/purego v0.9.0 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/gofrs/flock v0.12.1 // indirect
	github.com/hectane/go-acl v0.0.0-20230122075934-ca0b05cb1adb // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/lufia/plan9stats v0.0.0-20250317134145-8bc96cf8fc35 // indirect
	github.com/magiconair/properties v1.8.10 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/shirou/gopsutil/v4 v4.25.9 // indirect
	github.com/spf13/cast v1.10.0 // indirect
	github.com/spf13/cobra v1.10.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
	github.com/tklauser/go-sysconf v0.3.15 // indirect
	github.com/tklauser/numcpus v0.10.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.uber.org/dig v1.19.0 // indirect
	go.uber.org/fx v1.24.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	go.yaml.in/yaml/v3 v3.0.5
	golang.org/x/exp v0.0.0-20251009144603-d2f985daa21b // indirect
	golang.org/x/sys v0.37.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	golang.org/x/time v0.14.0 // indirect
)

// This section was automatically added by 'dda inv modules.add-all-replace' command, do not edit manually

replace github.com/DataDog/datadog-agent/pkg/config/nodetreemodel => ../datadog-nodetreemodel

replace github.com/DataDog/datadog-agent/pkg/config/setup => ../datadog-setup

replace github.com/DataDog/datadog-agent/pkg/util/scrubber => ../datadog-scrubber

replace github.com/DataDog/viper => ../datadog-viper
