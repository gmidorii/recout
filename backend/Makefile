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
	dev_appserver.py app.yaml --clear_datastore

.PHONY: wire
wire:
	$(WIRE_BIN) gen github.com/gmidorii/recout/backend/injector

.PHONY: test
test:
	go test -v ./...

.PHONY: mock
mock:
	#rm ./app/mock/*
	mockgen -source=./infra/repository/abstract.go -destination=./app/mock/repository.go -package=mock
	mockgen -source=./infra/pixela/client.go -destination=./app/mock/pixela_client.go -package=mock
