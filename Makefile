.PHONY: new-migration migrate-up migrate-down sqlc-generate

MIGRATIONS_DIR = "db/migrations"

new-migration:
	@read -p "Enter migration file name: " MIGRATION_NAME; \
	if [ -z "$$MIGRATION_NAME" ]; then \
		echo "Migration file name cannot be empty. Aborting..."; \
		exit 1; \
	fi; \
	migrate create -ext sql -dir "$(MIGRATIONS_DIR)" $$MIGRATION_NAME;

