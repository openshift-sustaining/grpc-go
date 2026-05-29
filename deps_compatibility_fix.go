//go:build tools

// This file ensures dependencies are pinned to minimum compatible versions.
// The build constraint ensures this file is ignored during normal builds.
package grpc

import (
	// The otel module introduced backwards-incompatible changes prior to version 1.28.0,
	// so we include them here to ensure the minimum compatible version is used.
	// This is a common issue when used in conjunction with the kubernetes instrumentation, which also depends on otel.
	_ "go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	_ "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	_ "go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	_ "go.opentelemetry.io/otel/metric"
	_ "go.opentelemetry.io/otel/sdk"
	_ "go.opentelemetry.io/otel/trace"
)
