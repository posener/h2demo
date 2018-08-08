run-server:
	go run ./h2-server.go

run-client:
	go run ./h2-client.go

run-client-http1:
	go run ./h2-client.go -version 1

run-h2conn-server:
	go run ./h2conn-server.go

run-h2conn-client:
	go run ./h2conn-client.go

certs:
	openssl req -newkey rsa:2048 -nodes -keyout server.key -x509 -days 365 -out server.crt
