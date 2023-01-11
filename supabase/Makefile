.PHONY: fmt
fmt:
	@deno fmt

include .env
.PHONY: slack
slack:	
	@deno run --allow-all functions/slack/index.ts
