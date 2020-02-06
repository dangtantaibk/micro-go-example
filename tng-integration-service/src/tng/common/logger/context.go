package logger

import (
	"context"
	"time"
)

type contextKey string

// String returns contextKey as string.
func (ck contextKey) String() string {
	return string(ck)
}

// Context key definition.
// RqIDCtxKey is request ID context key.
// RqClientIPCtxKey is request client IP context key.
// ExecTimeCtxKey is execute time context key.
var (
	RqIDCtxKey       = contextKey("request_id")
	RqClientIPCtxKey = contextKey("client_ip")
	RqExecTimeCtxKey = contextKey("exec_time")
	RqURICtxKey      = contextKey("request_uri")
	RqUserAgent      = contextKey("user_agent")
)

func GetRqUserAgent(ctx context.Context) string {
	return getStringFromCtx(ctx, RqUserAgent)
}

// GetRqIDFromCtx get request ID in context and returns as string.
func GetRqIDFromCtx(ctx context.Context) string {
	return getStringFromCtx(ctx, RqIDCtxKey)
}

// GetRqClientIPFromCtx get client IP in context and returns as string.
func GetRqClientIPFromCtx(ctx context.Context) string {
	return getStringFromCtx(ctx, RqClientIPCtxKey)
}

// GetRqExecTimeFromCtx get exec time in context and returns as string.
func GetRqExecTimeFromCtx(ctx context.Context) string {
	if ctx != nil {
		if val, ok := ctx.Value(RqExecTimeCtxKey).(time.Time); ok {
			elapsedTime := time.Since(val)
			return elapsedTime.String()
		}
	}
	return ""
}

// GetRqURIFromCtx get exec time in context and returns as string.
func GetRqURIFromCtx(ctx context.Context) string {
	return getStringFromCtx(ctx, RqURICtxKey)
}

// getStringFromCtx returns value string in context.
func getStringFromCtx(ctx context.Context, key contextKey) string {
	if ctx != nil {
		if val, ok := ctx.Value(key).(string); ok {
			return val
		}
	}
	return ""
}
