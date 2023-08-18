package routes

import (
	"receipt-processor/services"

	"github.com/gin-gonic/gin"
)

func ReceiptRoute(router *gin.Engine) {
	router.POST("/receipts/process", services.ProcessReceipt)
	router.GET("/receipts/:id/points", services.GetPoints)
}
