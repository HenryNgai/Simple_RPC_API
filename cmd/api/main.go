package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/HenryNgai/SIMPLE_RPC_API/proto/aggregator"
	"google.golang.org/grpc"
)

// Server implements the ContentAggregator gRPC service
// Think we need to add this in order to ensure new methods in service do not break functinality
// Backwards compatibility? TODO - Review
type Server struct {
	pb.UnimplementedContentAggregatorServer
}

// GetContent handles gRPC requests to fetch aggregated content
// ctx for managing lifetime of API calls, Database calls, etc
// req contains info sent by client
// Types defined in proto file
func (s *Server) GetContent(ctx context.Context, req *pb.ContentRequest) (*pb.ContentResponse, error) {
	log.Printf("Received GetContent request for user_id: %s, categories: %v", req.UserId, req.Categories)

	// Simulated content aggregation (replace this with your actual logic)
	aggregatedItems := []*pb.ContentItem{
		{
			Title:       "Tech News Today",
			Description: "Latest updates in technology.",
			Source:      "TechCrunch",
			Url:         "https://techcrunch.com/article",
			PublishedAt: time.Now().Format(time.RFC3339),
		},
		{
			Title:       "Sports Highlights",
			Description: "Top moments in sports.",
			Source:      "ESPN",
			Url:         "https://espn.com/article",
			PublishedAt: time.Now().Add(-time.Hour).Format(time.RFC3339),
		},
	}

	// Return the aggregated content
	return &pb.ContentResponse{Items: aggregatedItems}, nil
}

func main() {
	// Create a TCP listener on port 50051
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the ContentAggregator service (associates the grpcServer with a ContentAggregator service)
	pb.RegisterContentAggregatorServer(grpcServer, &Server{})

	log.Println("gRPC server is running on port 50051...")

	// Start serving gRPC requests
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
