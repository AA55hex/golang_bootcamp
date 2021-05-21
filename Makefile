include server/configs.env
RUNPATH=server/

go-get: 
	@echo "-----Installing dependencies-----"
	@cd $(RUNPATH) && go mod tidy
	@echo "-----Verifying dependencies-----"
	@cd $(RUNPATH) && go mod verify

go-test:
	@echo "-----Running docker-compose (detach mode)-----"
	@docker-compose up -d
	@echo "-----Running tests-----"
	@docker exec -it go_server go test -coverprofile cover.out ./...
	@echo "-----Removing docker-compose-----"
	@docker-compose down

go-run:
	@echo "-----Running docker-compose (detach mode)-----"
	@docker-compose up

curl-test:
	@echo "-----Running POST requests-----"
	@curl -XPOST -H "Content-type: application/json" -d '{ "name": "curl_book_0", "price": 9999, "genre": 1, "amount": 9999 }' 'localhost:3000/books/new'
	@curl -XPOST -H "Content-type: application/json" -d '{ "name": "curl_book_1", "price": 9999, "genre": 1, "amount": 9999 }' 'localhost:3000/books/new'
	@curl -XPOST -H "Content-type: application/json" -d '{ "name": "curl_book_2", "price": 9999, "genre": 1, "amount": 9999 }' 'localhost:3000/books/new'

clear:
	@echo "-----Removing docker-compose-----"	
	@docker-compose down