build:
	docker build -t dev:5000/registry:latest .
push:
	docker push dev:5000/registry:latest
all: build push