package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zp857/util/stack"
	"go.uber.org/zap"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

const (
	debugTemplate = "[Recovery from panic] error %v\nrequest %v\nstack %v"
	template      = "[Recovery from panic] error %v\nrequest %v\n"
)

func GinRecovery(logger *zap.SugaredLogger, debug bool, skip int) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}
				if debug {
					stackBytes := stack.GetStack(skip)
					logger.Errorf(debugTemplate, err, string(httpRequest), string(stackBytes))
				} else {
					logger.Errorf(template, err, string(httpRequest))
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
