package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	adapters "github.com/and-volkov/go-specs-greet/adapters"
	"github.com/and-volkov/go-specs-greet/adapters/httpserver"
	"github.com/and-volkov/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}
