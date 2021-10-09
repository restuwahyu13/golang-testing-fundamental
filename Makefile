test:
	@gotest -v ./...

gen:
ifdef f:
	@mockgen -source=./gomock/entity/${f}.go -destination=./gomock/mocks/${name}.go -package=mocks
endif

gmk:
	@gotest -v ./gomock/...