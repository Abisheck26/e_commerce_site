package routes

import (
	controller "payment/payment.client/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {
	r.POST("/createpayment", controller.HandlerCreatePayment)
}
