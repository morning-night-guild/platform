.PHONY: tool
tool:
	@aqua i

.PHONY: lint
lint:
	@buf lint

.PHONY: format
format:
	@buf format -w

.PHONY: gen
gen:
	@buf generate

.PHONY: doc
doc:
	@rm -rf doc && mkdir -p doc/api && protoc --doc_out=./doc/api --doc_opt=markdown,README.md "$(shell find . -name "*.proto")"
