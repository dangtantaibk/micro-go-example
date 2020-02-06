package logger

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Ensure logger implemented Logger.
var _ Logger = &logger{}

// Logger interface.
type Logger interface {
	Errorf(ctx context.Context, template string, args ...interface{})
	Warnf(ctx context.Context, template string, args ...interface{})
	Infof(ctx context.Context, template string, args ...interface{})
	CtxLog(ctx context.Context, args ...interface{})
}

type logger struct {
	errors map[interface{}][]string // errors
	warns  map[interface{}][]string // warns
	infos  map[interface{}][]string // infos
	locker sync.RWMutex             // for locking when read and write
}

func newLogger() logger {
	return logger{
		errors: make(map[interface{}][]string),
		warns:  make(map[interface{}][]string),
		infos:  make(map[interface{}][]string),
	}
}

func (l *logger) Errorf(ctx context.Context, template string, args ...interface{}) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if rqID := GetRqIDFromCtx(ctx); len(rqID) > 0 {
		var (
			messages = l.errors[ctx]
			message  = fmt.Sprintf(template, args...)
		)
		l.errors[ctx] = append(messages, message)
	}
}

func (l *logger) Warnf(ctx context.Context, template string, args ...interface{}) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if rqID := GetRqIDFromCtx(ctx); len(rqID) > 0 {
		var (
			messages = l.warns[ctx]
			message  = fmt.Sprintf(template, args...)
		)
		l.warns[ctx] = append(messages, message)
	}
}

func (l *logger) Infof(ctx context.Context, template string, args ...interface{}) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if rqID := GetRqIDFromCtx(ctx); len(rqID) > 0 {
		var (
			messages = l.infos[ctx]
			message  = fmt.Sprintf(template, args...)
		)
		l.infos[ctx] = append(messages, message)
	}
}

func (l *logger) CtxLog(ctx context.Context, args ...interface{}) {
	l.locker.Lock()
	defer l.locker.Unlock()

	if rqID := GetRqIDFromCtx(ctx); len(rqID) > 0 {
		var (
			sugaredLogger = *zap.S()
			code          int
			appName       string
		)
		if len(args) > 0 {
			code = args[0].(int)
		}
		if len(args) > 1 {
			appName = args[1].(string)
		}
		sugaredLogger = *sugaredLogger.With(
			zap.String(RqIDCtxKey.String(), rqID),
			zap.String(RqURICtxKey.String(), GetRqURIFromCtx(ctx)),
			zap.String(RqClientIPCtxKey.String(), GetRqClientIPFromCtx(ctx)),
			zap.String(RqExecTimeCtxKey.String(), GetRqExecTimeFromCtx(ctx)),
			zap.Int("status_code", code),
			zap.String("app_name", appName),
		)
		if messages, ok := l.errors[ctx]; ok {
			sugaredLogger = *sugaredLogger.With(zap.Strings("errors", messages))
			delete(l.errors, ctx)
		}
		if messages, ok := l.warns[ctx]; ok {
			sugaredLogger = *sugaredLogger.With(zap.Strings("warns", messages))
			delete(l.warns, ctx)
		}
		if messages, ok := l.infos[ctx]; ok {
			sugaredLogger = *sugaredLogger.With(zap.Strings("infos", messages))
			delete(l.infos, ctx)
		}
		sugaredLogger.Info()
	}
}
