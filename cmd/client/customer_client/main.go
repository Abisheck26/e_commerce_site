package main

import (

	//_ "github.com/kishorens18/ecommerce/cmd/client/docs"
	"e_commerce_site/cmd/client/customer_client/routes"
	grpcclient "e_commerce_site/cmd/client/grpcClient"

	"github.com/gin-gonic/gin"
)

// @title Documenting API (E-Commerce Webstie)
// @version 1
// @Description Buy Anything In Our Webstite

// @contact.name PAYMENT
// @contact.url http://localhost:8081/swagger/index.html
// @contact.name INVENTORY
// @contact.url http://localhost:8083/swagger/index.html#/
// @contact.email rohith.s@netxd.com
// @contact.phone 1234567787

// @securityDefinitions.apikey bearer
// @in header
// @name Authorization

// @host localhost:8080
// @BasePath /api/v1
func main() {
	_, conn := grpcclient.GetGrpcCustomerServiceClient()
	defer conn.Close()
	r := gin.Default()
	routes.AppRoutes(r)
	r.Run(":8080")

}
