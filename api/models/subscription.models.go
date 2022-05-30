package models

type Subscription struct {
	SubscriptionId string `gorm:"column:subscriptionId;primaryKey" json:"subscriptionId"`
	Name           string `gorm:"column:name" json:"name"`
	Description    string `gorm:"column:description" json:"description"`
	Amount         int64  `gorm:"column:amount" json:"amount"`
	UserId         string `gorm:"column:userId" json:"userId"`
}
