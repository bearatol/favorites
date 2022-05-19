package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/bearatol/favorites/internal/handler"
	"github.com/bearatol/favorites/internal/repository"
	"github.com/bearatol/favorites/internal/service"
	"github.com/bearatol/favorites/pkg/config"
	"github.com/bearatol/favorites/pkg/logger"
	"github.com/bearatol/favorites/pkg/middleware"
	gw "github.com/bearatol/favorites/proto/favorites/gen"
)

func main() {
	confFile := flag.String("config_file", "config.local.yml", "config file")
	confDir := flag.String("config_dir", "configs", "config directory")
	flag.Parse()

	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	defer func(lz *zap.SugaredLogger) {
		if err := lz.Sync(); err != nil {
			log.Fatal(err)
		}
	}(logger.Log())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := config.InitConfig(*confDir, *confFile, "yml"); err != nil {
		logger.Log().Fatal(err)
	}

	db, err := repository.NewPostgresDB(ctx)
	if err != nil {
		logger.Log().Fatal(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	serv := service.NewService(repo)
	app := handler.NewHandler(serv)

	serverGRPC := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.GRPCRecover,
			grpc_auth.UnaryServerInterceptor(middleware.Auth(config.Conf().Farmacy.JWTKey)),
		),
		grpc.ChainStreamInterceptor(
			grpc_recovery.StreamServerInterceptor(),
			grpc_auth.StreamServerInterceptor(middleware.Auth(config.Conf().Farmacy.JWTKey)),
		),
	)
	gw.RegisterFaivouritesServer(serverGRPC, app)

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = gw.RegisterFaivouritesHandlerFromEndpoint(ctx, mux, config.Conf().GRPC.URL, opts)
	if err != nil {
		logger.Log().Fatal(err)
	}

	//start grpc
	{
		l, err := net.Listen("tcp", config.Conf().GRPC.URL)
		if err != nil {
			logger.Log().With("error in listening", config.Conf().GRPC.URL).Fatal(err)
		}
		logger.Log().Infof("Start GRPC server: %s", config.Conf().GRPC.URL)
		go func() {
			logger.Log().Fatal(serverGRPC.Serve(l))
		}()
	}
	//start http
	{
		serverHTTP := &http.Server{
			Addr:    config.Conf().HTTP.URL,
			Handler: middleware.HTTPRecover(mux),
		}
		l, err := net.Listen("tcp", config.Conf().HTTP.URL)
		if err != nil {
			logger.Log().Fatal(err)
		}
		logger.Log().Infof("Start HTTP server http://%s", config.Conf().HTTP.URL)
		logger.Log().Fatal(serverHTTP.Serve(l))
	}
}
