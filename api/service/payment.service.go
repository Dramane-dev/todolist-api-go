package service

import "github.com/Dramane-dev/todolist-api/api/models"

type PaymentService interface {
	GetSubscriptionByUserId(userId string) (*models.Subscription, error)
	Subscribe(userId string, subscription *models.Subscription) (*models.Subscription, error)
	UnSubscribe(subscriptionId string) (*string, error)
}
