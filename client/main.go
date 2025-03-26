package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-crud-go/proto/generated"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewItemServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create Item
	item := &pb.CreateRequest{Id: "1", Name: "Laptop"}
	created, err := client.CreateItem(ctx, item)
	if err != nil {
		log.Fatalf("CreateItem failed: %v", err)
	}
	fmt.Println("Created Item:", created)

	// Read Item
	itemRes, err := client.ReadItem(ctx, &pb.ReadRequest{Id: "1"})
	if err != nil {
		log.Fatalf("ReadItem failed: %v", err)
	}
	fmt.Println("Retrieved Item:", itemRes)

	// Update Item
	updatedItem := &pb.UpdateRequest{Id: "1", Name: "Updated Laptop"}
	updatedRes, err := client.UpdateItem(ctx, updatedItem)
	if err != nil {
		log.Fatalf("UpdateItem failed: %v", err)
	}
	fmt.Println("Updated Item:", updatedRes)

	// List Items
	listRes, err := client.ListItems(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("ListItems failed: %v", err)
	}
	fmt.Println("All Items:", listRes)

	// Delete Item
	_, err = client.DeleteItem(ctx, &pb.DeleteRequest{Id: "1"})
	if err != nil {
		log.Fatalf("DeleteItem failed: %v", err)
	}
	fmt.Println("Item Deleted")
}
