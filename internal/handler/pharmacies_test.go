package handler

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"

	"github.com/bearatol/favorites/internal/service"
	mock_service "github.com/bearatol/favorites/internal/service/mocks"
	"github.com/bearatol/favorites/pkg/middleware"
	gw "github.com/bearatol/favorites/proto/favorites/gen"
)

func TestHandlerPharmacies(t *testing.T) {
	type mockBehavior func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64)

	type args struct {
		ctx context.Context
		req *gw.RequestData
	}
	fmt.Println("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq1")
	tests := []struct {
		name         string
		args         args
		mockBehavior mockBehavior
		expectation  error
	}{
		{
			name: "OK set",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: 1234, ACTION: actionADD},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {
				r.EXPECT().SetPharmacy(ctx, user, pharmacy).Return(nil)
			},
			expectation: newHandlerError(codes.OK, "", nil).Err(),
		},
		{
			name: "second add",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: 1234, ACTION: actionADD},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {
				r.EXPECT().SetPharmacy(ctx, user, pharmacy).Return(nil)
			},
			expectation: newHandlerError(codes.OK, "", nil).Err(),
		},
		{
			name: "OK delete",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: 1234, ACTION: actionDEL},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {
				r.EXPECT().DeletePharmacy(ctx, user, pharmacy).Return(nil)
			},
			expectation: newHandlerError(codes.OK, "", nil).Err(),
		},
		{
			name: "second delete",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: 1234, ACTION: actionDEL},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {
				r.EXPECT().DeletePharmacy(ctx, user, pharmacy).Return(nil)
			},
			expectation: newHandlerError(codes.OK, "", nil).Err(),
		},
		{
			name: "method doesn't exist",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: 1234, ACTION: "NO_METHOD"},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {},
			expectation:  newHandlerError(codes.InvalidArgument, "pharmacy: method doesn't exist", nil).Err(),
		},
		{
			name: "method is empty",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: 1234},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {},
			expectation:  newHandlerError(codes.InvalidArgument, "pharmacy: method doesn't exist", nil).Err(),
		},
		{
			name: "ID is empty",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ACTION: actionADD},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {},
			expectation:  newHandlerError(codes.Canceled, "the id of element is missing", nil).Err(),
		},
		{
			name: "ID < 0",
			args: args{
				ctx: context.WithValue(context.Background(), middleware.ContextUserID, int64(123)),
				req: &gw.RequestData{ID: -1, ACTION: actionADD},
			},
			mockBehavior: func(r *mock_service.MockPharmacies, ctx context.Context, user, pharmacy int64) {},
			expectation:  newHandlerError(codes.Canceled, "the id of element is missing", nil).Err(),
		},
	}
	fmt.Println("qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockPharmacies(c)
			userID, ok := test.args.ctx.Value(middleware.ContextUserID).(int64)
			if !ok || userID <= 0 {
				t.Error("user id doesn't valid")
				return
			}
			test.mockBehavior(repo, test.args.ctx, userID, test.args.req.ID)
			service := &service.Service{Pharmacies: repo}
			h := NewHandler(service)
			data := &gw.RequestData{
				ID:     test.args.req.ID,
				ACTION: test.args.req.ACTION,
			}

			_, err := h.Pharmacies(test.args.ctx, data)
			assert.ErrorIs(t, err, test.expectation)
		})
	}
}
