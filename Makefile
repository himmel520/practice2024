deep:
	dep-tree entropy cmd/main.go
.PHONY: deep

run:
	docker compose -f compose.yml up --build --no-log-prefix --attach websocket_server
.PHONY: run