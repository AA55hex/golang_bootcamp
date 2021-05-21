include server/configs.env
RUNPATH=server/

go-get: 
	@echo "-----Installing dependencies-----"
	@cd $(RUNPATH) && go mod tidy
	@echo "-----Verifying dependencies-----"
	@cd $(RUNPATH) && go mod verify

go-test: clear
	@echo "-----Running docker-compose (detach mode)-----"
	@docker-compose up -d
	@echo "-----Running tests-----"
	@docker exec -it go_server go test -coverprofile cover.out ./...
	@echo "-----Removing docker-compose-----"
	@docker-compose down

go-run:
	@echo "-----Running docker-compose (detach mode)-----"
	@docker-compose up

clear:
	@echo "-----Removing docker-compose-----"	
	@docker-compose down