MIGRATIONS_PATH = ./cmd/migrate/migrations
MIGRATE=migrate

.PHONY:migrate-create
migration:
	$(MIGRATE) create -seq -ext sql -dir $(MIGRATIONS_PATH) $(name)

.PHONY:migrate-up
migrate-up:
	$(MIGRATE) -path=$(MIGRATIONS_PATH)	-database="$(DB_ADDR)" up 
.PHONY:force
force:
	$(MIGRATE) -path=$(MIGRATIONS_PATH) -database="$(DB_ADDR)" force $(version)
.PHONY:migrate-down
migrate-down:
	$(MIGRATE) -path=$(MIGRATIONS_PATH)	-database="$(DB_ADDR)" down