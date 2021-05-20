include server/configs.env
RUNPATH=server/
MODULE=github.com/AA55hex/golang_bootcamp/server

go-get: 
	@echo "-----Installing dependencies-----"
	@cd $(RUNPATH) && go mod tidy
	@echo "-----Verifying dependencies-----"
	@cd $(RUNPATH) && go mod verify

go-test: clear
	@echo "-----Building database docker-----"
	@docker build -t go_test_database -f Dockerfile.database .
	@echo "-----Running docker-----"
	@docker run -d --network-alias $(MYSQL_HOST) --name bootcamp_container go_test_database
	@echo "-----Running tests-----"
	@cd $(RUNPATH) && go test ./...
	@echo "-----Removing container-----"
	@docker rm -f bootcamp_container

go-run:
	@docker-compose up

clear:
	@docker rm -f bootcamp_container