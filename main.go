package main

import (
	"go-microservices/determinant"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	var svc determinant.DeterminantService
	svc = determinant.NewService()

	determinantHandler := httptransport.NewServer(
		determinant.MakeUppercaseEndpoint(svc),
		determinant.DecodeDeterminantRequest,
		determinant.EncodeResponse,
	)

	http.Handle("/determinant", determinantHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
