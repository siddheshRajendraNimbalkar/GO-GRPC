package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	porto "github.com/siddheshRajendraNimbalkar/GO-GRPC/protoc"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := porto.NewNameClient(conn)

	router := gin.Default()

	router.POST("/get-full-name", func(c *gin.Context) {
		var req struct {
			FirstName  string `json:"first_name"`
			MiddleName string `json:"middle_name"`
			LastName   string `json:"last_name"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		grpcReq := &porto.NameRequest{
			FirstName:  req.FirstName,
			MiddleName: req.MiddleName,
			LastName:   req.LastName,
		}

		grpcRes, err := client.NameReply(context.Background(), grpcReq)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"full_name": grpcRes.FullName})
	})

	log.Println("Gin server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start Gin server: %v", err)
	}
}
