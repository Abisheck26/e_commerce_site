package grpcclient

import (
	"log"
	"sync"

	 pb "e_commerce_site/ecommerce_proto/payment_proto"
	"google.golang.org/grpc"
)

var paymentOnce sync.Once

type GrpcPaymentServiceClient pb.PaymentServiceClient

var (
	paymentInstance GrpcPaymentServiceClient
)

func GetGrpcPaymnentService() (GrpcPaymentServiceClient,*grpc.ClientConn) {
	var conn *grpc.ClientConn
	paymentOnce.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		paymentInstance = pb.NewPaymentServiceClient(conn)
	})

	return paymentInstance,conn
}
