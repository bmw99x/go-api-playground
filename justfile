# The default command
default: build run

build_swagger:
    swag init

# Install all the dependencies
deps:
    go mod download

# Build the Go app
build:
    go build -o app ./cmd

# Run the Go app
run:
	go run ./cmd/main.go

# Build and Run the Docker container
docker-up:
    docker-compose up --build

# Stop the Docker container
docker-down:
    docker-compose down


# Test the Go app
test:
    cd tests && go test -v
