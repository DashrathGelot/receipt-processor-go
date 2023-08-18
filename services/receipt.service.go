package services

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils/rules"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetPoints(ctx *gin.Context) {
	id := ctx.Param("id")

	if storage.IsContains(id) {
		var pointsRespo models.PointsResponse
		pointsRespo.Points = storage.GetPoints(id)

		ctx.IndentedJSON(http.StatusOK, pointsRespo)
	} else {
		ctx.JSON(http.StatusBadRequest, "Invalid ID")
	}

}

func ProcessReceipt(ctx *gin.Context) {
	var receipt models.Receipt

	if err := ctx.BindJSON(&receipt); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id := uuid.New()
	var receiptRespose models.ReceiptResponse
	receiptRespose.ID = id.String()

	var errors []string
	var points int

	//Rule 1
	if len(strings.TrimSpace(receipt.Retailer)) > 0 {
		points += rules.Alphanumeric(receipt.Retailer)
	} else {
		errors = append(errors, "Invalid Retailer Name")
	}

	//Rule 2 & Rule 3
	if total, err := rules.ConverToFloat(receipt.Total); err == nil {
		if rules.IsTotalRound(total) {
			points += 50
		}
		if rules.IsTotalMultiple(total) {
			points += 25
		}
	} else {
		errors = append(errors, "Invalid Total amount value it should be number")
	}

	//Rule 4
	if len(receipt.Items) > 0 {
		points += rules.Rule4Items(receipt.Items)
	} else {
		errors = append(errors, "Invalid Items")
	}

	//Rule 5
	if reward, err := rules.Rule5TrimDesc(receipt.Items); err == nil {
		points += reward
	} else {
		errors = append(errors, "Invalid Item Price")
	}

	//Rule 6
	if odd, err := rules.IsPurchaseDateOdd(receipt.PurchaseDate); err == nil {
		if odd {
			points += 6
		}
	} else {
		errors = append(errors, "Invalid Purchase Date")
	}

	//Rule 7
	if time2to4, err := rules.IsTime2to4(receipt.PurchaseTime); err == nil {
		if time2to4 {
			points += 10
		}
	} else {
		errors = append(errors, "Invalid Purchase Time")
	}

	if len(errors) > 0 {
		ctx.JSON(http.StatusBadRequest, errors)
	} else {
		storage.Save(id.String(), points)
		ctx.IndentedJSON(http.StatusOK, receiptRespose)
	}
}
