include ../../Makefile.Common

ATTRIBUTION_CMD := GOOS="" GOARCH="" $(GOCMD) run ../../internal/buildscripts/thirdparty-attribution/main.go -policy-sha 2f893b59cafc1ff751aab1fc6d32e9cbee0569bbba6d710265e437615a7a9d6c
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
