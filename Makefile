start:
	go build
	./cushon-technical-solution -port=8080

start-db:
	docker-compose up -d

stop-db:
	docker-compose down