package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/BassemHalim/memeDB/memeService/internal/utils"
	"github.com/BassemHalim/memeDB/memeService/internal/db"
	"github.com/BassemHalim/memeDB/memeService/internal/server"
	pb "github.com/BassemHalim/memeDB/proto/memeService"

	"google.golang.org/grpc"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("Service", "MEME_SERVICE")
	db, err := db.New()
	if err != nil {
		log.Error("Failed to connect to the database", "ERROR", err)
		os.Exit(1)
	}
	defer db.Close()

	serverPort := utils.GetEnvOrDefault("SERVER_PORT", "50051")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Error("Failed to listen to TCP", "PORT", serverPort)
	}
	log.Info("Server started on port", "Port", serverPort)
	s := grpc.NewServer()
	pb.RegisterMemeServiceServer(s, server.New(db, log))

	c := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Info("Shutting down gracefully...")
		s.GracefulStop()
		done <- true
	}()

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Error("Failed to serve", "ERROR", err)
			done <- true
		}
	}()

	<-done
	log.Info("Server stopped")

}
