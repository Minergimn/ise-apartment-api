package logger

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type ctxKey struct{}

var attachedLoggerKey = &ctxKey{}

var globalLogger *zap.SugaredLogger

func fromContext(ctx context.Context) *zap.SugaredLogger {
	var result *zap.SugaredLogger
	if attachedLogger, ok := ctx.Value(attachedLoggerKey).(*zap.SugaredLogger); ok {
		result = attachedLogger
	} else {
		result = globalLogger
	}

	jaegerSpan := opentracing.SpanFromContext(ctx)
	if jaegerSpan != nil {
		if spanCtx, ok := opentracing.SpanFromContext(ctx).Context().(jaeger.SpanContext); ok {
			result = result.With("trace-id", spanCtx.TraceID())
		}
	}

	return result
}

//ErrorKV comment for linter
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Error(append([]interface{}{message}, kvs...)...)
}
//WarnKV comment for linter
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Warn(append([]interface{}{message}, kvs...)...)
}

//InfoKV comment for linter
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Info(append([]interface{}{message}, kvs...)...)
}

//DebugKV comment for linter
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Debug(append([]interface{}{message}, kvs...)...)
}

//FatalKV comment for linter
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	fromContext(ctx).Fatal(append([]interface{}{message}, kvs...)...)
}

//AttachLogger comment for linter
func AttachLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, attachedLoggerKey, logger)
}

//CloneWithLevel comment for linter
func CloneWithLevel(ctx context.Context, newLevel zapcore.Level) *zap.SugaredLogger {
	return fromContext(ctx).
		Desugar().
		WithOptions(WithLevel(newLevel)).
		Sugar()
}

//SetLogger comment for linter
func SetLogger(newLogger *zap.SugaredLogger) {
	globalLogger = newLogger
}

func init() {
	notSugaredLogger, err := zap.NewProduction()
	if err != nil {
		log.Panic(err)
	}

	globalLogger = notSugaredLogger.Sugar()
}