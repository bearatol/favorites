package handler

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/bearatol/favorites/internal/service"
	"github.com/bearatol/favorites/pkg/logger"
	gw "github.com/bearatol/favorites/proto/favorites/gen"
)

const (
	actionADD = "ADD"
	actionDEL = "DELETE"
)

type handlerError struct {
	code codes.Code
	msg  string
	err  error
}

func (e *handlerError) Err() error {
	if e.code == 0 {
		logger.Log().Infof("code: %d, msg: %s", e.code, e.msg)
		return status.Errorf(e.code, "%s", e.msg)
	}
	logger.Log().With(
		"code", e.code,
		"msg", e.msg,
	).Error(e.err)
	return status.Errorf(e.code, "%s, error: [%v]", e.msg, e.err)
}

func newHandlerError(code codes.Code, msg string, err error) *handlerError {
	return &handlerError{
		code: code,
		msg:  msg,
		err:  err,
	}
}

type Handler struct {
	gw.UnimplementedFaivouritesServer
	serv *service.Service
}

func NewHandler(serv *service.Service) *Handler {
	return &Handler{
		serv: serv,
	}
}
