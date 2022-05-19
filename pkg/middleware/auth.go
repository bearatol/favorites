package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	GRPCookieHeader = "grpcgateway-cookie"
	CookieHeader    = "Cookie"
	CookieJWTKey    = "APP.token"
	JWTUserIDKey    = "fUserId"
)

func parseCookies(rawCookies string) []*http.Cookie {
	header := http.Header{}
	header.Add(CookieHeader, rawCookies)
	request := http.Request{Header: header}

	return request.Cookies()
}

func parseToken(c, jwtKey string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(c, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func getTokenFromCookie(cookies []*http.Cookie) (string, error) {
	for _, c := range cookies {
		if c.Name == CookieJWTKey {
			return c.Value, nil
		}
	}
	return "", status.Errorf(codes.Unauthenticated, "Request unauthenticated with "+CookieJWTKey)
}

func userClaimFromToken(s *jwt.Token) int64 {
	m, ok := s.Claims.(jwt.MapClaims)
	if !ok {
		return 0
	}

	if s.Claims.Valid() != nil {
		return 0
	}

	switch v := m[JWTUserIDKey].(type) {
	case float64:
		return int64(v)
	case json.Number:
		i, _ := v.Int64()
		return i
	}
	return 0
}

func Auth(jwtKey string) func(inCtx context.Context) (outCtx context.Context, err error) {
	return func(inCtx context.Context) (outCtx context.Context, err error) {
		rawToken, err := grpc_auth.AuthFromMD(inCtx, "bearer")
		if rawToken == "" {
			rawCookies := metautils.ExtractIncoming(inCtx).Get(GRPCookieHeader)
			rawToken, err = getTokenFromCookie(parseCookies(rawCookies))
		}

		if err != nil {
			return inCtx, err
		}

		// парсим jwt-token
		tokenInfo, err := parseToken(rawToken, jwtKey)
		if err != nil {
			return inCtx, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
		}

		grpc_ctxtags.Extract(inCtx).Set("auth.sub", tokenInfo.Claims)
		outCtx = context.WithValue(inCtx, ContextUserID, userClaimFromToken(tokenInfo))
		return outCtx, nil
	}
}
