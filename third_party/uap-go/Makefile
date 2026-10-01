include ../../Makefile.Common

# Only this nested owner distinguishes anchored originals from owned source.
# The common first-party header checker still checks every owned Go/shell file.
ALL_SRC_AND_SHELL_CMD := GOOS="" GOARCH="" $(GOCMD) run ./internal/attribution -owned-source

ALL_SRC_AND_DOC_CMD := GOOS="" GOARCH="" $(GOCMD) run ./internal/attribution -owned-docs

.PHONY: checkattribution
checkattribution:
	GOOS="" GOARCH="" $(GOCMD) run ./internal/attribution

checklicense addlicense misspell misspell-correction: checkattribution
