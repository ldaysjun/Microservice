package handlers

import (
	"TimeService/svc"
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"os"
	//stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func LogMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		w := log.NewSyncWriter(os.Stderr)
		logger := log.NewLogfmtLogger(w)
		logger.Log("req",request)
		resp,err := next(ctx,request)
		return resp,err
	}
}
//
//func InstrumentingMiddleware(next endpoint.Endpoint) endpoint.Endpoint  {
//	var dur metrics.Histogram = prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
//		Namespace: "myservice",
//		Subsystem: "api",
//		Name:     "request_duration_seconds",
//		Help:     "Total time spent serving requests.",
//	}, []string{})
//	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
//		defer func(begin time.Time) { dur.Observe(time.Since(begin).Seconds()) }(time.Now())
//		resp,err := next(ctx,request)
//		return resp,err
//	}
//}

func WrapEnpoints(in svc.Endpoints) svc.Endpoints{
	//in.WrapAllExcept(LogMiddleware)
	//in.WrapAllExcept(InstrumentingMiddleware)
	return in
}