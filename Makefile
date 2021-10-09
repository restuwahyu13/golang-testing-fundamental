############################
## Golang Command Testing
############################
test:
	@gotest -v ./...

####################################
## Golang Command Testing Coverage
####################################

coverage: line1.a cover.a line2.a execute.a

line1.a:
	@echo "========================================"
	@echo "============ Testing Report ============"
	@echo "========================================"
	@echo ""

cover.a:
	@gotest -v -coverprofile=coverage.out ./...

line2.a:
	@echo ""
	@echo "========================================"
	@echo "======= Coverage Testing Report ========"
	@echo "========================================"
	@echo ""

execute.a:
	@go tool cover -func=coverage.out

# gen:
# ifdef f:
# 	@mockgen -source=./gomock/entity/${f}.go -destination=./gomock/mocks/${name}.go -package=mocks
# endif

# gmk:
# 	@gotest -v ./gomock/...