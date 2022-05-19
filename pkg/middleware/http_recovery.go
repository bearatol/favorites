package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/bearatol/favorites/pkg/logger"
)

// HTTPRecover middleware
func HTTPRecover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				logger.Log().With("stack_trace", string(debug.Stack())).Error(p)

				w.WriteHeader(http.StatusInternalServerError)
				buf := bytes.NewBuffer([]byte(fmt.Sprintf("recovered from panic: %v\n", p)))
				_, _ = buf.Write(debug.Stack())
				_, _ = w.Write(buf.Bytes())
			}
		}()
		next.ServeHTTP(w, r)
	})
}
