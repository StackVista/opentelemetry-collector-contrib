include ../../Makefile.Common

ATTRIBUTION_CMD := GOOS="" GOARCH="" $(GOCMD) run ../../internal/buildscripts/thirdparty-attribution/main.go -policy-sha 59e621591c6cc9e2a2097c9a5c2567e4b44de72b294e0c6d0f963fc0aace95ba
ALL_SRC_AND_SHELL_CMD := $(ATTRIBUTION_CMD) -owned-source
ALL_SRC_AND_DOC_CMD := $(ATTRIBUTION_CMD) -owned-docs

.PHONY: checkattribution
checkattribution:
	$(ATTRIBUTION_CMD) -git-revision HEAD

checklicense addlicense misspell misspell-correction: checkattribution

# Original source is immutable; format only owned Go source.
fmt: checkattribution $(GOFUMPT) $(GOIMPORTS)
	@files=$$($(ATTRIBUTION_CMD) -owned-source | awk '/\.go$$/'); \
	if [ -n "$$files" ]; then $(GOFUMPT) -l -w -extra $$files; $(GOIMPORTS) -w $$files; fi

generate: checkattribution fmt
	@:
