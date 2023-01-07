.PHONY: fmt
fmt:
	@deno fmt

include .env
.PHONY: slack
slack:	
	@deno run --allow-all --unstable functions/slack/index.ts
