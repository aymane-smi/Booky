run_compose_build:
	sudo docker-compose up -d --build
run_compose:
	sudo docker-compose up -d
test_server:
	go run main.go
build_binary:
	go build -o booky && ./book