run:
	@echo "  >  Running...\t\t"
	go run ./cmd/client
	
dockerrun:
	@echo "  >  Running docker...\t\t"
	docker build -t gotest .
	docker run -v /var/run/docker.sock:/var/run/docker.sock gotest
