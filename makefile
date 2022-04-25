build:
	docker build -t fury-core-go-template:dev .

run:
	docker run -p 8080:8080 fury-core-go-template:dev

run-db:
	docker-compose up -d db

clean:
	docker-compose rm -s -f
