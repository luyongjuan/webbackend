run_app:
	docker build -t operation_app .
	--docker stop operation
	--docker rm operation
	docker run -p 8080:8080 --name operation -d operation_app

build:
	docker build -t traffic_operation .

run:
	POSTGRES_URL=172.17.0.2 \
	POSTGRES_USER=postgres \
    	POSTGRES_PASSWORD=postgres \
    	go run cmd/app/app.go

