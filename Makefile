#!make
include .env
export $(shell sed 's/=.*//' .env)

run: mod db app

gen:
	@rm -rf ./cmd
	@rm -rf ./internal
	@rm -rf ./pkg
	@mkdir -p cmd
	@mkdir -p internal/config
	@mkdir -p internal/repositories
	@mkdir -p internal/handler
	@mkdir -p internal/services
	@mkdir -p internal/requests
	@mkdir -p internal/app
	@mkdir -p internal/app
	@mkdir -p pkg
	@swagger generate server -f ./docs/swagger.yaml -C ./swagger-templates/default-server.yml \
		-T ./swagger-templates/templates \
		--name example
	@swagger generate client -f ./docs/swagger.yaml -C ./swagger-templates/default-client.yml \
		-T ./swagger-templates/templates --tags=pet --name=pet
	@swagger generate client -f ./docs/swagger.yaml -C ./swagger-templates/default-client.yml \
    		-T ./swagger-templates/templates --tags=store --name=store
	@swagger generate model -m pet -f ./docs/swagger.yaml -C ./swagger-templates/default-model.yml \
    		-T ./swagger-templates/templates
	@swagger generate client -f ./docs/swagger.yaml -C ./swagger-templates/default-client.yml \
			 -T ./swagger-templates/templates --tags=user --name=user

mod:
	go mod vendor

db:
	docker-compose up -d --force-recreate --no-deps --build postgresql

app:
	docker-compose up -d --force-recreate --no-deps --build ms-go

swag:
	@swag init -g ./cmd/server/main.go  --parseDependency --parseInternal