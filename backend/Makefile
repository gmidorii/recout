BIN=recout-backend
WIRE_BIN=wire

.PHONY: build
build: wire
	go build -o ./$(BIN)

.PHONY: deploy
deploy:
	gcloud app deploy --quiet

.PHONY: local
local:
	dev_appserver.py app.yaml

.PHONY: wire
wire:
	$(WIRE_BIN) gen github.com/gmidorii/recout/backend/injector
