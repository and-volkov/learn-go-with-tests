package main_test

import (
	"fmt"
	"testing"

	"github.com/and-volkov/go-specs-greet/adapters"
	"github.com/and-volkov/go-specs-greet/adapters/grpcserver"
	"github.com/and-volkov/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
