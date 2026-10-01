include ../../Makefile.Common

ATTRIBUTION_CMD := GOOS="" GOARCH="" $(GOCMD) run ../../internal/buildscripts/thirdparty-attribution/main.go -policy-sha 26199b8076817aa3cd38185d9a89b3732a3ae868e68c1a536a1d8eaacc2e0ffa
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
