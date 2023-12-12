package rescue

import (
	"context"
	"github.com/zp857/util/stack"
	"go.uber.org/zap"
)

// Recover is used with defer to do cleanup on panics.
// Use it like:
//
//	defer Recover(func() {})
func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		zap.L().Sugar().Errorf("%+v\n%s", p, stack.GetStack(3))
	}
}

// RecoverCtx is used with defer to do cleanup on panics.
func RecoverCtx(ctx context.Context, cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}

	if p := recover(); p != nil {
		zap.L().Sugar().Errorf("%+v\n%s", p, stack.GetStack(3))
	}
}
