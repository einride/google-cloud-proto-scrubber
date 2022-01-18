# Code generated by go.einride.tech/mage-tools. DO NOT EDIT.
# To learn more, see .mage/magefile.go and https://github.com/einride/mage-tools.

.DEFAULT_GOAL := all

magefile := .mage/tools/bin/magefile

$(magefile): .mage/go.mod .mage/*.go
	@cd .mage && go run go.einride.tech/mage-tools/cmd/build

.PHONY: clean-mage-tools
clean-mage-tools:
	@git clean -fdx .mage/tools

.PHONY: all
all: $(magefile)
	@$(magefile) all

.PHONY: convco-check
convco-check: $(magefile)
ifndef rev
	$(error missing argument rev="...")
endif
	@$(magefile) convcoCheck $(rev)

.PHONY: format-markdown
format-markdown: $(magefile)
	@$(magefile) formatMarkdown

.PHONY: format-yaml
format-yaml: $(magefile)
	@$(magefile) formatYaml

.PHONY: git-verify-no-diff
git-verify-no-diff: $(magefile)
	@$(magefile) gitVerifyNoDiff

.PHONY: go-mod-tidy
go-mod-tidy: $(magefile)
	@$(magefile) goModTidy

.PHONY: go-test
go-test: $(magefile)
	@$(magefile) goTest

.PHONY: golangci-lint
golangci-lint: $(magefile)
	@$(magefile) golangciLint

.PHONY: goreview
goreview: $(magefile)
	@$(magefile) goreview
