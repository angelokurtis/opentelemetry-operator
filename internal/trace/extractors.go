package trace

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"runtime"
	"strings"
)

func StartSpanFromContext(ctx context.Context) (trace.Span, context.Context) {
	provider := otel.GetTracerProvider()
	tracer := provider.Tracer("opentelemetry-operator")
	function, line := getCaller()
	ctx, span := tracer.Start(ctx, function, trace.WithAttributes(attribute.String("debug.line", line)))
	return span, ctx
}

func getCaller() (string, string) {
	pc, filename, line, _ := runtime.Caller(2)
	l := fmt.Sprintf("%s:%d", filename, line)
	f := runtime.FuncForPC(pc).Name()
	return strings.ReplaceAll(f, "github.com/open-telemetry/opentelemetry-operator/", ""),
		strings.ReplaceAll(l, "/home/kurtis/wrkspc/github.com/angelokurtis/opentelemetry-operator/", "")
}
