package controller

import (
	"fmt"
	"net/http"
	"payment/constants"
	"payment/payment.server/controllers"

	grpcclient "payment/payment.client/grpc.client"
	pb "payment/proto"

	"github.com/gin-gonic/gin"
)

func HandlerCreatePayment(c *gin.Context) {

	grpcClient, _ := grpcclient.GetGrpcClientInstance()
	var res pb.PaymentDetails
	token := c.GetHeader("Authorization")
	
	result1, err1 := controllers.ExtractCustomerID(token, constants.SecretKey)
	//var request pb.PaymentDetails
	res.CustomerID=result1
	fmt.Println(res.CustomerID)
	fmt.Println(err1)


	
	if err := c.ShouldBindJSON(&res); err != nil {
		fmt.Println("12")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	fmt.Println(result1)
	response, err := grpcClient.CreatePayment(c.Request.Context(), &res)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": response})
}
