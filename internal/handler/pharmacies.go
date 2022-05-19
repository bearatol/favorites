package handler

import (
	"context"
	"fmt"

	gw "github.com/bearatol/favorites/proto/favorites/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/bearatol/favorites/pkg/middleware"
)

func (h *Handler) Pharmacies(ctx context.Context, data *gw.RequestData) (*emptypb.Empty, error) {
	userID, ok := ctx.Value(middleware.ContextUserID).(int64)
	if !ok || userID <= 0 {
		return nil, newHandlerError(codes.Unauthenticated, "user do not have correct id", nil).Err()
	}
	if data.ID <= 0 {
		return nil, newHandlerError(codes.Canceled, "the id of element is missing", nil).Err()
	}

	switch data.ACTION {
	case actionADD:
		if err := h.serv.Pharmacies.SetPharmacy(ctx, userID, data.ID); err != nil {
			return nil, newHandlerError(codes.Aborted, "add pharmacy", err).Err()
		}
	case actionDEL:
		if err := h.serv.Pharmacies.DeletePharmacy(ctx, userID, data.ID); err != nil {
			return nil, newHandlerError(codes.Aborted, "delete pharmacy", err).Err()
		}
	default:
		return nil, newHandlerError(codes.InvalidArgument, "pharmacy: method doesn't exist", nil).Err()
	}

	return &emptypb.Empty{}, newHandlerError(codes.OK, fmt.Sprintf("pharmacy: %+v, user id: %d", data, userID), nil).Err()
}
