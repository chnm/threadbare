# Application 
# ------------------------------------------------------------
.PHONY : build, up, down, test

build :
	docker compose --profile threadbare build 

up :
	docker compose --profile api --force-recreate --detach 

down :
	docker compose --profile db --profile threadbare down

test : 
	go test -v ./...

# Database
# ------------------------------------------------------------
.PHONY : migration, db-up, db-down

migration :
	@read -p "What is the slug for the migration? " migration;\
	migrate create -dir common/db/migrations -ext sql -seq $$migration

db-up :
	@echo "Migrating to current version of database"
	migrate -database "$(DBSTR_LOCAL)" -path common/db/migrations up

db-down :
	migrate -database "$(DBSTR_LOCAL)" -path common/db/migrations down 1

# Deploy 
# ------------------------------------------------------------
.PHONY : deploy 

deploy : export THREADBARE_VERSION=release
deploy :
	docker compose --profile threadbare build --parallel
	docker push ghcr.io/hepplerj/threadbare-crawler:release