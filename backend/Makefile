BIN=recout-backend

.PHONY: build
build:
	go build -o ./$(BIN)

.PHONY: deploy
deploy:
	gcloud app deploy --quiet

.PHONY: local
local:
	dev_appserver.py app.yaml
