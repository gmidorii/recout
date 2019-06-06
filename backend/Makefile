BIN=recout-backend
WIRE_BIN=wire
PACK_FILE=front.pb

.PHONY: build
build: wire
	go build -o ./$(BIN)

.PHONY: generate
generate:
	npm --prefix ../front/ run generate

.PHONY: pack
pack:
	brbundle pack -f $(PACK_FILE) ../front/dist/

.PHONY: deploy
deploy: generate pack
	gcloud app deploy --quiet

.PHONY: local
local: pack
	dev_appserver.py app.yaml

.PHONY: clear-local
local: pack
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
