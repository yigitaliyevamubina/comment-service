package main

import (
	"comment-service/config"
	pb "comment-service/genproto/comment_service"
	"comment-service/pkg/db"
	"comment-service/pkg/logger"
	"comment-service/service"
	"google.golang.org/grpc"
	"net"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "comment-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, _, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal("sql connection error", logger.Error(err))
	}

	commentService := service.NewCommentService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("cannot listen", logger.Error(err))
	}

	server := grpc.NewServer()
	pb.RegisterCommentServiceServer(server, commentService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := server.Serve(lis); err != nil {
		log.Fatal("server cannot serve", logger.Error(err))
	}
}
