package mysql

import (
	"github.com/Dramane-dev/todolist-api/api/models"
	"github.com/google/uuid"
)

func (db *MySQLDatabase) GetSubscriptionById(subscriptionId string) (*models.Subscription, error) {
	var subscription *models.Subscription

	errWhenGetSubscriptionById := db.connection.Model(&models.Subscription{}).Where("subscriptionId = ?", subscriptionId).Find(&subscription).Error

	if errWhenGetSubscriptionById != nil {
		return nil, errWhenGetSubscriptionById
	}

	return subscription, nil
}

func (db *MySQLDatabase) GetSubscriptionByUserId(userId string) (*models.Subscription, error) {
	var subscription *models.Subscription

	errWhenGetSubscriptionByUserId := db.connection.Model(&models.Subscription{}).Where("userId = ?", userId).Find(&subscription).Error

	if errWhenGetSubscriptionByUserId != nil {
		return nil, errWhenGetSubscriptionByUserId
	}

	return subscription, nil
}

func (db *MySQLDatabase) Subscribe(userId string, subscription *models.Subscription) (*models.Subscription, error) {
	subscription.SubscriptionId = "SUB" + uuid.NewString()
	subscription.UserId = userId

	errWhenSaveSubscription := db.connection.Model(&models.Subscription{}).Create(subscription).Error

	if errWhenSaveSubscription != nil {
		return nil, errWhenSaveSubscription
	}

	return db.GetSubscriptionById(subscription.SubscriptionId)
}

func (db *MySQLDatabase) UnSubscribe(subscriptionId string) (*string, error) {
	errWhenUnSubscription := db.connection.Model(&models.Subscription{}).Where("subscriptionId = ?", subscriptionId).Delete(subscriptionId).Error

	if errWhenUnSubscription != nil {
		return nil, errWhenUnSubscription
	}

	response := "Unsubscribe successfully ✅"
	return &response, nil
}
