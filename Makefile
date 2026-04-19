target := fractalflame

build:
	@mkdir -p build
	@echo "Building $(target)"
	@go build -o ./build/${TARGET} ./cmd/${TARGET}

test:
	gotestsum