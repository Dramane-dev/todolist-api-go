build-run:
	docker build -t simplytodo-go .
	docker-compose up -d

build:
	docker build -t simplytodo-go .