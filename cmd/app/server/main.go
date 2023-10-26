package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/97vaibhav/Vessel_tracking/cmd/vesselpb"
	"github.com/97vaibhav/Vessel_tracking/internal/domain/service"
	"github.com/97vaibhav/Vessel_tracking/internal/handler"
	"github.com/97vaibhav/Vessel_tracking/internal/infrastructure/datastore"
	"github.com/97vaibhav/Vessel_tracking/utils"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	zapConf := utils.NewZapConfig()
	logger, err := zapConf.Build()
	if err != nil {
		panic(err)
	}
	config, err := utils.LoadConfig()
	if err != nil {
		logger.Panic("failed to load config", zap.Error(err))
	}
	if config.Debug {
		logger.Info("Service RUN on DEBUG mode")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Tokyo")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil && config.Debug {
		logger.Panic("cant open database connection", zap.Error(err))
	}
	defer dbConn.Close()

	vesselRepo := datastore.NewMysqlVesselRepository(dbConn)
	vesselUsecase := service.NewVesselService(vesselRepo)
	list, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		logger.Panic("SOMETHING UNEXPECTED HAPPEN", zap.Error(err))
	}
	srv := handler.NewVesselServer(vesselUsecase)
	server := grpc.NewServer()
	vesselpb.RegisterVesselServiceServer(server, srv)
	reflection.Register(server)
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		<-ctx.Done()
		logger.Info("stoping gRPC server...")
		server.GracefulStop()
		logger.Info("done server.GracefulStop")
	}()

	fmt.Println("Server Run at ", config.ServerAddress)

	err = server.Serve(list)
	if err != nil {
		logger.Panic("Unexpected Error", zap.Error(err))
		server.GracefulStop()
	}
}
