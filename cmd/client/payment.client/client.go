package main

import (
	grpcclient "payment/payment.client/grpc.client"
	"payment/payment.client/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	_, conn := grpcclient.GetGrpcClientInstance()
	defer conn.Close()

	r := gin.Default()
	routes.AppRoutes(r)
	r.Run(":8080")

}

// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	pb "project1/proto"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"google.golang.org/grpc"
// )

// var (
// 	mongoclient *mongo.Client
// 	ctx         context.Context
// 	server      *gin.Engine
// )

// func main() {
// 	r := gin.Default()

// 	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
// 	if err != nil {
// 		fmt.Println("111")
// 		log.Fatalf("Failed to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewPaymentServiceClient(conn)

// 	// response, err := client.CreateCustomer(context.Background(), &pb.Customer{CustomerId: 12,Balance:3000 })
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to call SayHello: %v", err)
// 	// }

// 	// fmt.Printf("Response: %s\n", response)

// 	// response1, err := client.CreatePayment(context.Background(), &pb.PaymentDetails{
// 	// 	Cardno:          12348,
// 	// 	Cvvverified:     1237,
// 	// 	Balance:         200,
// 	// })
// 	// if err != nil {
// 	// 	log.Fatalf("Failed to call SayHello: %v", err)
// 	// }

// 	// fmt.Printf("Response: %s\n", response1)
// 	fmt.Println("1234")
// 	r.POST("/createpayment", func(c *gin.Context) {
// 		var request pb.PaymentDetails
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			fmt.Println("12")
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		response, err := client.CreatePayment(c.Request.Context(), &request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})
// 	r.Run(":8080")
// }
