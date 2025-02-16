package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/BassemHalim/memeDB/memeService/src/server"
	pb "github.com/BassemHalim/memeDB/proto/memeService"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func buildDBConnString() string {
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "5432")
	user := getEnvOrDefault("DB_USER", "postgres")
	password := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbname := getEnvOrDefault("DB_NAME", "memedb")

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)
}

func initDB() (*sql.DB, error) {

	db, err := sql.Open("postgres", buildDBConnString())
	if err != nil {
		return nil, fmt.Errorf("Error connecting to the database: %v", err)
	}
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging the database: %v", err)
	}

	return db, err
}

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("Service", "MEME_SERVICE")
	db, err := initDB()
	if err != nil {
		log.Error("Failed to connect to the database", "ERROR", err)
		os.Exit(1)
	}
	defer db.Close()

	serverPort := getEnvOrDefault("SERVER_PORT", "50051")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", serverPort))
	if err != nil {
		log.Error("Failed to listen to TCP", "PORT", serverPort)
	}
	log.Info("Server started on port", "Port", serverPort)
	s := grpc.NewServer()
	pb.RegisterMemeServiceServer(s, server.NewMemeServer(db, log))

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
