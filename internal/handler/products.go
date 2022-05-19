package handler

import (
	"context"
	"fmt"

	gw "github.com/bearatol/favorites/proto/favorites/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bearatol/favorites/pkg/middleware"
)

func (h *Handler) Products(ctx context.Context, data *gw.RequestData) (*emptypb.Empty, error) {
	userID, ok := ctx.Value(middleware.ContextUserID).(int64)
	if !ok || userID <= 0 {
		return nil, newHandlerError(codes.Unauthenticated, "user do not have correct id", nil).Err()
	}
	if data.ID <= 0 {
		return nil, newHandlerError(codes.Canceled, "the id of element is missing", nil).Err()
	}

	switch data.ACTION {
	case actionADD:
		if err := h.serv.Products.SetProduct(ctx, userID, data.ID); err != nil {
			return nil, newHandlerError(codes.Aborted, "add product", err).Err()
		}
	case actionDEL:
		if err := h.serv.Products.DeleteProduct(ctx, userID, data.ID); err != nil {
			return nil, newHandlerError(codes.Aborted, "delete product", err).Err()
		}
	default:
		return nil, newHandlerError(codes.InvalidArgument, "product: method doesn't exist", nil).Err()
	}

	return &emptypb.Empty{}, newHandlerError(codes.OK, fmt.Sprintf("product: %+v, user id: %d", data, userID), nil).Err()
}
