package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"receipt-processor/models"
	"receipt-processor/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	return gin.Default()
}

func getReceipt() models.Receipt {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
		Total: "35.35",
	}

	return receipt
}

func TestProcessReceipt(t *testing.T) {
	router := SetUpRouter()
	router.POST("/receipts/process", services.ProcessReceipt)

	jsonReceipt, _ := json.Marshal(getReceipt())
	request, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonReceipt))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)

	var receiptRespose models.ReceiptResponse
	json.Unmarshal(w.Body.Bytes(), &receiptRespose)
	assert.NotEmpty(t, receiptRespose)
}

func TestGetPoints(t *testing.T) {
	router := SetUpRouter()
	router.POST("/receipts/process", services.ProcessReceipt)
	router.GET("/receipts/:id/points", services.GetPoints)

	jsonReceipt, _ := json.Marshal(getReceipt())
	request, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonReceipt))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)
	assert.Equal(t, http.StatusOK, w.Code)

	var receiptRespose models.ReceiptResponse
	json.Unmarshal(w.Body.Bytes(), &receiptRespose)
	assert.NotEmpty(t, receiptRespose)

	var mockresponse models.PointsResponse
	mockresponse.Points = 28

	url := "/receipts/" + receiptRespose.ID + "/points"
	req, _ := http.NewRequest("GET", url, nil)
	wget := httptest.NewRecorder()
	router.ServeHTTP(wget, req)

	var pointsresponse models.PointsResponse
	json.Unmarshal(wget.Body.Bytes(), &pointsresponse)

	assert.Equal(t, mockresponse.Points, pointsresponse.Points)
	assert.Equal(t, http.StatusOK, wget.Code)
}
