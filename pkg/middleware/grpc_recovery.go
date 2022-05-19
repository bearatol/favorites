package middleware

import (
	"context"
	"fmt"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bearatol/favorites/pkg/logger"
)

// GRPCRecover interceptor.
func GRPCRecover(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			errMsg := fmt.Sprintf("recovered from panic: %v", p)
			err = status.Errorf(codes.Internal, errMsg)
			logger.Log().With(
				"stack_trace", string(debug.Stack()),
				"panic", true,
			).Error(errMsg)
		}
	}()
	return handler(ctx, req)
}
