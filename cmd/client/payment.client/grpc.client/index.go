package grpcclient

import (
	pb "e_commerce_site/ecommerce_proto/payment_proto"
	"log"
	"sync"

	"google.golang.org/grpc"
)

var once sync.Once

type GrpcClient pb.PaymentServiceClient

var (
	instance GrpcClient
)

func GetGrpcClientInstance() (GrpcClient, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	once.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		instance = pb.NewPaymentServiceClient(conn)
	})

	return instance, conn
}
