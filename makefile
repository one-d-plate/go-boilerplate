.PHONY: rename-module table migrate rollback

table:
	mkdir -p migrations
	goose -dir migrations create $(name) sql

migrate:
	goose -dir migrations postgres "user=sso_user password=sso_pass dbname=sso_db host=localhost sslmode=disable" up

rollback:
	goose -dir migrations postgres "user=sso_user password=sso_pass dbname=sso_db host=localhost sslmode=disable" down

rename-module:
ifndef NEW_MODULE
	$(error ❌ Harap masukkan nama module baru, misalnya: make rename-module NEW_MODULE=github.com/username/my-new-service)
endif
	@bash scripts/rename_module.sh $(NEW_MODULE)