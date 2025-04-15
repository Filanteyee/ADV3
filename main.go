package main

import (
	"ADV3/handler"
	"ADV3/proto/gen"
	"ADV3/proto/orderpb"
	"ADV3/repository"
	"ADV3/usecase"
	"context"
	"log"
	"net"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	db := client.Database("orders_db")
	repo := repository.NewOrderRepository(db)

	// Establish gRPC connection to ProductService
	productConn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to ProductService: %v", err)
	}
	defer productConn.Close()
	productClient := gen.NewProductServiceClient(productConn)

	uc := usecase.NewOrderUsecase(repo, productClient)

	// Start gRPC server
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		orderpb.RegisterOrderServiceServer(grpcServer, handler.NewGRPCOrderServer(uc))

		log.Println("✅ gRPC server running on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	r := gin.Default()
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/static/orders.html")
	})

	h := handler.NewOrderHandler(uc)
	r.POST("/orders", h.CreateOrder)
	r.GET("/orders", h.GetOrders)
	r.GET("/orders/:id", h.GetOrderByID)
	r.PATCH("/orders/:id", h.UpdateOrderStatus)
	r.DELETE("/orders/:id", h.DeleteOrder)
	r.PATCH("/orders/:id/cancel", h.CancelOrder)

	log.Println("✅ REST server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}
}
