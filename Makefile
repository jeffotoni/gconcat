
cover:
	./scripts/coverage.sh

gmod:
	@rm -f go.*;
	./scripts/update_go_mod.sh

test: 
	go test -race -v ./...
	go test -v -tags musl -covermode atomic -coverprofile=coverage.out ./...	
