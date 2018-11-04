runserver:
	go build -o calculator ./server
	./calculator

clireq:
	go run client/main.go ${op} ${x} ${y}
