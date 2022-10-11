
cover:
	./scripts/coverage.sh

gmod:
	@rm -f go.*;
	./scripts/update_go_mod.sh
