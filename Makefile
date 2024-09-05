running:
	nodemon --exec go run main.go --signal SIGTERM
runvite:
	cd frontend && yarn dev
build:
	docker-compose up -d --build