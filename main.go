package main

import (
	"context"
	"go-microservices/determinant"
	"go-microservices/eigen"
	"go-microservices/inverse"
	"go-microservices/multiplication"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var determinantService determinant.DeterminantService
	determinantService = determinant.NewService()
	determinantService = determinant.NewLoggingService(log.With(logger, "Component:", "Determinant"), determinantService)

	var eigenValueService eigen.EigenValueService
	eigenValueService = eigen.NewService()
	eigenValueService = eigen.NewLoggingService(log.With(logger, "Component:", "EigenValue"), eigenValueService)

	var inverseService inverse.InverseService
	inverseService = inverse.NewService()
	inverseService = inverse.NewLoggingService(log.With(logger, "Component:", "Inverse Operation"), inverseService)

	proxy := []string{"http://localhost:8080/inverse"}
	var multiplicationService multiplication.MultiplicationService
	multiplicationService = multiplication.NewService()
	multiplicationService = multiplication.ProxyingMiddleware(context.Background(), proxy, logger)(multiplicationService)
	multiplicationService = multiplication.NewLoggingService(log.With(logger, "Component:", "Singular Value Decomposition"), multiplicationService)

	determinantHandler := httptransport.NewServer(
		determinant.MakeUppercaseEndpoint(determinantService),
		determinant.DecodeDeterminantRequest,
		determinant.EncodeResponse,
	)

	eigenValueHandler := httptransport.NewServer(
		eigen.MakeEigenValueEndpoint(eigenValueService),
		eigen.DecodeEigenValueRequest,
		eigen.EncodeEigenValueResponse,
	)

	inverseHandler := httptransport.NewServer(
		inverse.MakeInverseEndpoint(inverseService),
		inverse.DecodeInverseRequest,
		inverse.EncodeInverseResponse,
	)

	multiplicationHandler := httptransport.NewServer(
		multiplication.MakeMultiplicationEndpoint(multiplicationService),
		multiplication.DecodeMultiplicationRequest,
		multiplication.EncodeMultiplicationResponse,
	)

	http.Handle("/determinant", determinantHandler)
	http.Handle("/eigen", eigenValueHandler)
	http.Handle("/inverse", inverseHandler)
	http.Handle("/multiplication", multiplicationHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
