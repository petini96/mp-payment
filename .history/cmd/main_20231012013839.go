package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/petini96/payment/mp"
)

type WebhookNotificationRequest struct {
	ID            int         `json:"id"`
	LiveMode      bool        `json:"live_mode"`
	Type          string      `json:"type"`
	DateCreated   time.Time   `json:"date_created"`
	ApplicationID int         `json:"application_id"`
	UserID        string      `json:"user_id"`
	Version       int         `json:"version"`
	ApiVersion    string      `json:"api_version"`
	Action        string      `json:"action"`
	Data          interface{} `json:"data"`
}

func main() {
	// envFile, _ := godotenv.Read("../.env")

	// mpToken := envFile["MP_TOKEN"]
	r := gin.Default()
	r.GET("/payment/notification", func(ctx *gin.Context) {
		collectionStatus := ctx.Query("collection_status")
		switch collectionStatus {
		case "approved":
			fmt.Println("approved")
		case "rejected":
			fmt.Println("rejected")
		case "in_process":
			fmt.Println("in_process")
		default:
			fmt.Println("nenhuma opção")
		}
		ctx.JSON(http.StatusOK, "Nosso webhook é um sucesso!")
	})
	r.POST("/payment/notification", func(ctx *gin.Context) {
		var req WebhookNotificationRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		switch req.Action {
		case "payment.created":
			log.Fatal("criou")
			idMPPayment := req.Data.(map[string]interface{})["id"].(string)

			apiUrl := "https://api.mercadopago.com"
			resourceApi := "/v1/payments/" + idMPPayment
			queryParams := make(mp.KeyValueString)

			headers := make(mp.KeyValueString)
			headers["Authorization"] = "Bearer " + mp.TokenSandbox

			resultRequest, _ := mp.MakeGetRequest(apiUrl, resourceApi, queryParams, headers)
			var dat map[string]interface{}
			if err := json.Unmarshal(resultRequest, &dat); err != nil {
				panic(err)
			}

		case "payment.pending":
			log.Fatal("pending")
		case "payment.failure":
			log.Fatal("falhou")
		default:
			fmt.Println("nenhuma opção")
		}

		ctx.JSON(http.StatusOK, "Recebido webhook")
		return
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
