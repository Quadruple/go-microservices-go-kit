package svd

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"golang.org/x/time/rate"

	"github.com/sony/gobreaker"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
)

func ProxyingMiddleware(ctx context.Context, instances []string, logger log.Logger) SvdServiceMiddleware {
	if len(instances) == 0 {
		logger.Log("proxy_to", "none")
		return func(next SvdService) SvdService { return next }
	}

	var (
		qps         = 100
		maxAttempts = 3
		maxTime     = 250 * time.Millisecond
	)

	var (
		instanceList = instances
		endpointer   sd.FixedEndpointer
	)
	logger.Log("proxy_to", fmt.Sprint(instanceList))
	for _, instance := range instanceList {
		var e endpoint.Endpoint
		e = makeSvdProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), qps))(e)
		endpointer = append(endpointer, e)
	}

	balancer := lb.NewRoundRobin(endpointer)
	retry := lb.Retry(maxAttempts, maxTime, balancer)

	return func(next SvdService) SvdService {
		return proxymw{ctx, next, retry}
	}
}

type proxymw struct {
	ctx  context.Context
	next SvdService
	svd  endpoint.Endpoint
}

func (mw proxymw) GetSingularValueDecomposition(matrix [][]float64) ([][]float64, error) {
	response, err := mw.svd(mw.ctx, svdRequest{Matrix: matrix})
	if err != nil {
		return nil, err
	}

	resp := response.(svdResponse)
	if resp.Err != "" {
		return resp.SvdMatrix, errors.New(resp.Err)
	}
	return mw.next.GetSingularValueDecomposition(matrix)
}

func makeSvdProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/svd"
	}

	return httptransport.NewClient(
		"GET",
		u,
		EncodeSvdRequest,
		DecodeSvdResponse,
	).Endpoint()
}
