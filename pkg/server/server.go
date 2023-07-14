package server

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/tim-pipi/cloudwego-api-gateway/internal/config"
	"github.com/tim-pipi/cloudwego-api-gateway/pkg/clientpool"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/obs-opentelemetry/logging/logrus"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
)

// Starts the API Gateway server
func Start(svcConfig *config.ServiceConfig) {
	hlog.SetLogger(hertzlogrus.NewLogger())
	// Set log level based on environment variable
	hlog.SetLevel(hlog.LevelDebug)

	f, err := os.OpenFile(svcConfig.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Write to stderr and log file
	mw := io.MultiWriter(os.Stderr, f)
	hlog.SetOutput(mw)
	serviceName := "cloudwego-api-gateway"

	// Observability
	allowMetrics := os.Getenv("ALLOW_METRICS")
	if allowMetrics == "1" {
		p := provider.NewOpenTelemetryProvider(
			provider.WithServiceName(serviceName),
			// Support setting ExportEndpoint via environment variables: OTEL_EXPORTER_OTLP_ENDPOINT
			provider.WithExportEndpoint(":4317"),
			provider.WithInsecure(),
		)
		defer p.Shutdown(context.Background())
	}

	tracer, cfg := hertztracing.NewServerTracer()
	h := server.Default(
		server.WithHostPorts(":8080"),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	// Status code middleware
	h.Use(func(c context.Context, ctx *app.RequestContext) {
		ctx.Next(c)
		hlog.CtxDebugf(c, "Request status code: %d", ctx.Response.StatusCode())
	})

	// Sample /ping route
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, "Hello, CloudWeGo!")
	})

	cp := clientpool.NewClientPool(svcConfig.IDLDir, svcConfig.EtcdAddr)
	h.Any("/:ServiceName/:ServiceMethod", func(c context.Context, ctx *app.RequestContext) {
		// Check that JSON is valid
		var req interface{}
		err := json.Unmarshal(ctx.Request.BodyBytes(), &req)

		// TODO - return proper HTTP status codes based on RPC server error
		if err != nil {
			hlog.Error("Invalid JSON: ", err.Error())
			ctx.Error(err)
			errorJSON := map[string]interface{}{
				"error": ctx.Errors.Errors(),
			}

			ctx.JSON(consts.StatusBadRequest, errorJSON)
			return
		}

		cp.Call(c, ctx)
	})

	h.Spin()
}
