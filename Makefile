run_compose_build:
	sudo docker-compose up -d --build
run_compose:
	sudo docker-compose up -d
down_compose:
	sudo docker-compose down
test_server:
	go run main.go
build_binary:
	go build -o booky && ./book