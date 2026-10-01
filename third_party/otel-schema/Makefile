include ../../Makefile.Common

ATTRIBUTION_CMD := GOOS="" GOARCH="" $(GOCMD) run ../../internal/buildscripts/thirdparty-attribution/main.go -policy-sha c0c82d90d44d0605b5eb1619e30f91d1a6c353465a8c582aec00d9cd05ca4e68
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
