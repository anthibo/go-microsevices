package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"log-service/logs"
	"net"

	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}
	err := l.Models.Insert(logEntry)
	if err != nil {
		fmt.Println(err)
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	// return response
	res := &logs.LogResponse{Result: "logged"}
	return res, nil
}

func (app *Config) gRPCListen() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", gRpcPort))
	if err != nil {
		log.Fatalf("failed to listen for gRPC server: %v", err)
	}
	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC server starting on port: %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen for gRPC server: %v", err)
	}
}
