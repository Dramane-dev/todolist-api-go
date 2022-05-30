package controllers

import (
	"net/http"
	"os"

	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

// type Charge struct {
// 	Formula      string `json:"formula"`
// 	Amount       int64  `json:"amount"`
// 	ReceiptEmail string `json:"receiptEmail"`
// }

// func (paymentService *PaymentController) Subscribe(ctx *gin.Context) {
// 	var subscription *Charge
// 	ctx.BindJSON(&subscription)

// 	stripeApiKey := os.Getenv("STRIPE_TEST_KEY")
// 	stripe.Key = stripeApiKey

// 	_, errWhenCreateCharge := charge.New(&stripe.ChargeParams{
// 		Amount:   stripe.Int64(subscription.Amount),
// 		Currency: stripe.String(string(stripe.CurrencyEUR)),
// 		Source: &stripe.SourceParams{
// 			Token: stripe.String("tok_visa"),
// 		},
// 		ReceiptEmail: stripe.String(subscription.ReceiptEmail),
// 	})

// 	if errWhenCreateCharge != nil {
// 		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errWhenCreateCharge})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"message":      "Payment has been sucessfully received ✅",
// 		"subscription": subscription,
// 	})
// }

func (paymentService *PaymentController) Subscribe(ctx *gin.Context) {
	var subscription *models.Subscription
	userId, ok := ctx.Params.Get("userId")

	if !ok {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "userId not provided or incorrect...❌"})
		return
	}

	ctx.BindJSON(&subscription)

	userAlreadySubscribe, userNotSubscribe := paymentService.database.GetSubscriptionByUserId(userId)

	if userNotSubscribe != nil {
		stripeApiKey := os.Getenv("STRIPE_TEST_KEY")
		stripe.Key = stripeApiKey

		_, errWhenCreateCharge := charge.New(&stripe.ChargeParams{
			Amount:   stripe.Int64(subscription.Amount),
			Currency: stripe.String(string(stripe.CurrencyEUR)),
			Source: &stripe.SourceParams{
				Token: stripe.String("tok_visa"),
			},
		})

		if errWhenCreateCharge != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errWhenCreateCharge})
			return
		}

		subscriptionSaved, errWhenSavedSubscriptionToDatabase := paymentService.database.Subscribe(userId, subscription)

		if errWhenSavedSubscriptionToDatabase != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errWhenSavedSubscriptionToDatabase)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message":      "Payment has been sucessfully received ✅",
			"subscription": subscriptionSaved,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":      "User already subscribe ✅",
		"subscription": userAlreadySubscribe,
	})
}
