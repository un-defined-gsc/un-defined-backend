.PHONY: help

dev.migrate.up:
	@echo "Migrating up..."
	@goose -dir migrations postgres "host=localhost password=api user=api dbname=api  sslmode=disable" up

dev.migrate.reset:
	@echo "Migrating down..."
	@goose -dir migrations postgres "host=localhost password=api user=api dbname=api  sslmode=disable" reset

dev.swagger.init:
	@echo "Generating swagger..."
	@swag init  --parseVendor  -d . -g /cmd/api-service/main.go 


help:
	@echo "Usage: make <target>"
	@echo ""
	@echo "Targets:"
	@echo "  dev.migrate.up      Migrate up"
	@echo "  dev.migrate.reset   Migrate down"
	@echo "  dev.swagger.init    Generate swagger"
	@echo "  help                Show this help"