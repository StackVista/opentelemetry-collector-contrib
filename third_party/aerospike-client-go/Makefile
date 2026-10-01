include ../../Makefile.Common

ATTRIBUTION_CMD := GOOS="" GOARCH="" $(GOCMD) run ../../internal/buildscripts/thirdparty-attribution/main.go -policy-sha 9a993ccd7cc622805e2870e821811c5921852a3c2a730c29bd7ac44ae723f8fc
ALL_SRC_AND_SHELL_CMD := $(ATTRIBUTION_CMD) -owned-source
ALL_SRC_AND_DOC_CMD := $(ATTRIBUTION_CMD) -owned-docs

.PHONY: checkattribution
checkattribution:
	$(ATTRIBUTION_CMD)

checklicense addlicense misspell misspell-correction: checkattribution

# Original source is immutable; format only owned Go source.
fmt: checkattribution $(GOFUMPT) $(GOIMPORTS)
	@files=$$($(ATTRIBUTION_CMD) -owned-source | awk '/\.go$$/'); \
	if [ -n "$$files" ]; then $(GOFUMPT) -l -w -extra $$files; $(GOIMPORTS) -w $$files; fi

generate: checkattribution fmt
	@:

# Ginkgo rejects Go's parallel/count flags; repeat full invocations instead.
GOTEST_OPT := $(GOTEST_RACE_OPT) -timeout $(GOTEST_TIMEOUT) --tags=$(GO_BUILD_TAGS)
test-twice: $(GOTESTSUM)
	@set -e; for iteration in 1 2; do \
		$(GOTESTSUM) --packages="./..." -- $(GOTEST_OPT) -count=1; \
	done
