package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var defaultLogger logger

func init() {
	zapCfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(zap.InfoLevel),
		Development:       false,
		DisableCaller:     true,
		DisableStacktrace: true,
		Sampling:          &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:          "json",
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",
			TimeKey:    "ts",
			EncodeTime: EpochTimeEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	zapLogger, err := zapCfg.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(zapLogger)

	// Initial default logger
	defaultLogger = newLogger()
}

// Errorf logs an error.
func Errorf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Errorf(ctx, template, args...)
}

// Warnf logs an error.
func Warnf(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Warnf(ctx, template, args...)
}

// Infof logs an information.
func Infof(ctx context.Context, template string, args ...interface{}) {
	defaultLogger.Infof(ctx, template, args...)
}

// CtxLog logs with context.
func CtxLog(ctx context.Context, args ...interface{}) {
	defaultLogger.CtxLog(ctx, args...)
}

// Debug logs debug message.
func Debug(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

// Log logs normal message.
func Log(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

// EpochTimeEncoder serializes a time.Time to a floating-point number of seconds
// since the Unix epoch.
func EpochTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(t.Unix())
}
