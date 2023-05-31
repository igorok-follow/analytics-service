package app

import (
	"context"
	"github.com/igorok-follow/analytics-service/app/endpoint"
	"github.com/igorok-follow/analytics-service/app/repository"
	"github.com/igorok-follow/analytics-service/app/service"
	context_middleware "github.com/igorok-follow/analytics-service/middleware/context"
	recovery_middleware "github.com/igorok-follow/analytics-service/middleware/recovery"
	"github.com/igorok-follow/analytics-service/tools/event_handler"
	"github.com/igorok-follow/analytics-service/tools/logger"
	"github.com/igorok-follow/analytics-service/tools/tracing"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"github.com/igorok-follow/analytics-service/config"
	"github.com/igorok-follow/analytics-service/extra/api"
)

func Run(config *config.Config) {
	conn := repository.NewConnection(config.Database.Uri)
	err := conn.Open()
	if err != nil {
		logger.FatalError("Configure connection error", err)
	}

	defer conn.DB.Close()

	logger.Debug("Connected to database")

	repositories := repository.NewRepositoryContainer(conn.DB)

	tracer, err := tracing.InitTracer(config.Tracing.JaegerUri, config.Server.Name)
	if err != nil {
		logger.FatalError("tracer init error", err)
	}

	eventHandler := event_handler.NewEventHandler(3000, config.Analytics.ApiKey)
	eventHandler.Run()

	deps := &service.Dependencies{
		EventHandler: eventHandler,
		Tracer:       tracer,
		Repositories: repositories,
	}
	services := service.NewServices(deps)
	endpoints := endpoint.NewEndpointContainer(services, deps)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			recovery_middleware.UnaryServerInterceptor(),
			context_middleware.UnaryServerInterceptor(),
		)),
	)
	api.RegisterEventServer(s, endpoints.EventEndpoint)

	logger.Info("Starting "+config.Server.Name+"...",
		logger.String("host", config.Server.Host),
		logger.String("port", config.Server.Port))

	l, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		logger.FatalError("listen tcp error", err)
	}

	go func() {
		logger.FatalError("serve error", s.Serve(l))
	}()

	gwconn, err := grpc.DialContext(
		context.Background(),
		"localhost"+config.Server.Port,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		logger.FatalError("Failed to dial server", err)
	}

	gwmux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   false,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "X-Remote-Address", "X-Requested-With", "Authorization"},
		AllowCredentials: true,
	}).Handler(gwmux)
	err = api.RegisterEventHandler(context.Background(), gwmux, gwconn)
	if err != nil {
		logger.FatalError("Failed to register gateway", err)
	}

	gwServer := &http.Server{
		Addr:    config.Gateway.Port,
		Handler: gwmux,
	}

	logger.Info("Serving gRPC-Gateway...",
		logger.String("host", config.Gateway.Host),
		logger.String("port", config.Gateway.Port))
	logger.FatalError("gRPC-Gateway serving error", http.ListenAndServe(gwServer.Addr, handler))
}
