package kafka

import (
	"context"
)

type HandlerFunc func(ctx context.Context, msg []byte)
