.PHONY: aqua
aqua: # export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
	@go run github.com/aquaproj/aqua-installer@latest --aqua-version v1.23.1

.PHONY: tool
tool:
	@aqua i

.PHONY: fmt
fmt:
	@yamlfmt
