package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "grpc-crud-go/proto/generated"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterItemServiceServer(grpcServer, &server{items: make(map[string]*pb.Item)})

	fmt.Println("Server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type server struct {
	pb.UnimplementedItemServiceServer
	items map[string]*pb.Item
	mu    sync.Mutex
}

// Create Item
func (s *server) CreateItem(ctx context.Context, req *pb.CreateRequest) (*pb.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item := &pb.Item{Id: req.Id, Name: req.Name}
	s.items[req.Id] = item
	return item, nil
}

// Read Item
func (s *server) ReadItem(ctx context.Context, req *pb.ReadRequest) (*pb.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, exists := s.items[req.Id]
	if !exists {
		return nil, fmt.Errorf("Item not found")
	}
	return item, nil
}

// Update Item
func (s *server) UpdateItem(ctx context.Context, req *pb.UpdateRequest) (*pb.Item, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, exists := s.items[req.Id]
	if !exists {
		return nil, fmt.Errorf("Item not found")
	}
	item.Name = req.Name
	return item, nil
}

// Delete Item
func (s *server) DeleteItem(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.items, req.Id)
	return &pb.Empty{}, nil
}

// List Items
func (s *server) ListItems(ctx context.Context, req *pb.Empty) (*pb.ItemList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	var itemList pb.ItemList
	for _, item := range s.items {
		itemList.Items = append(itemList.Items, *item)
	}
	return &itemList, nil
}
