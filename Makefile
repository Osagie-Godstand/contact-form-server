build-api:
	@go build -o bin/api ./cmd/sendit/

run: build-api
	@./bin/api

clean: 
	@rm -rf bin

