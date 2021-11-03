package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	pb "github.com/ozonmp/ise-apartment-api/pkg/ise-apartment-api"
)

var (
	httpTotalRequests = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_microservice_requests_total",
		Help: "The total number of incoming HTTP requests",
	})
)

func createGatewayServer(grpcAddr, gatewayAddr string, allowedOrigins []string, swaggerDocPath string) *http.Server {
	// Create a client connection to the gRPC Server we just started.
	// This is where the gRPC-Gateway proxies the requests.
	conn, err := grpc.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(opentracing.GlobalTracer()),
			),
		),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to dial server")
	}

	mux := runtime.NewServeMux()
	if err := pb.RegisterIseApartmentApiServiceHandler(context.Background(), mux, conn); err != nil {
		log.Fatal().Err(err).Msg("Failed registration handler")
	}

	mux.HandlePath("GET", "/swagger.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		swaggerBytes, err := os.ReadFile(swaggerDocPath)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		w.Write(swaggerBytes)
	})

	gatewayServer := &http.Server{
		Addr:    gatewayAddr,
		Handler: cors(tracingWrapper(swaggerWrapper(mux)), allowedOrigins),
	}

	return gatewayServer
}

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		httpTotalRequests.Inc()
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || errors.Is(err, opentracing.ErrSpanContextNotFound) {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				"ServeHTTP",
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}

func cors(h http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		providedOrigin := r.Header.Get("Origin")
		matches := false
		for _, allowedOrigin := range allowedOrigins {
			if providedOrigin == allowedOrigin {
				matches = true
				break
			}
		}

		if matches {
			w.Header().Set("Access-Control-Allow-Origin", providedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

func swaggerWrapper(h http.Handler) http.Handler {
	docsServer := http.FileServer(http.Dir("./swagger/dist"))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swagger" || strings.HasPrefix(r.URL.Path, "/swagger/") {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/swagger")
			docsServer.ServeHTTP(w, r)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
