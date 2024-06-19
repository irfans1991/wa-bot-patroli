package usecase

import (
	"context"
	"fmt"
	"time"
	"whatsapp-bot/domain"
	"whatsapp-bot/handler"

	"go.mau.fi/whatsmeow"
	"gorm.io/gorm"
)

var (
	client *whatsmeow.Client
)

func PollNewActivities(db *gorm.DB) {
	ctx = context.Background()
	var lastCheckedID uint = 0
	message := "Halo Cinta ! "

	for {
		var Users []domain.Users
		db.Where("id > ?", lastCheckedID).Order("id asc").Find(&Users)

		for _, activity := range Users {
			fmt.Printf("New activity detected: %v\n", activity)
			lastCheckedID = uint(activity.Id)
			message += activity.Name
		}

		// Sleep for a while before the next poll
		time.Sleep(10 * time.Second)
	}
	handler.SendMessages(ctx, client, message)

	// Send the message
}
