package main

import (
	"go-microservices/determinant"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var svc determinant.DeterminantService
	svc = determinant.NewService()
	svc = determinant.NewLoggingService(log.With(logger, "Component:", "Determinant"), svc)

	determinantHandler := httptransport.NewServer(
		determinant.MakeUppercaseEndpoint(svc),
		determinant.DecodeDeterminantRequest,
		determinant.EncodeResponse,
	)

	http.Handle("/determinant", determinantHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
