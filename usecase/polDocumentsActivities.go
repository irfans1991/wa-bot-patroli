package usecase

import (
	"context"
	"fmt"
	"log"
	"time"
	"whatsapp-bot/domain"
	"whatsapp-bot/handler"

	"go.mau.fi/whatsmeow"
	"gorm.io/gorm"
)

func PollNewActivitiesDocuments(db *gorm.DB, client *whatsmeow.Client) {
	ctx = context.Background()
	var lastCheckedID uint = 0
	var message string
	// var Unit string

	for {
		var documents []domain.Documents

		db.Where("id > ?", lastCheckedID).Order("id asc").Find(&documents)
		timeNow := time.Now().Format("2006-01-02")

		for _, activity := range documents {

			fmt.Printf("New activity Documents detected	: %v\n", activity)
			lastCheckedID = uint(activity.Id)
			dates := activity.Created_at.Format("2006-01-02")
			// Unit = activity.Unit

			if dates == timeNow {
				message = fmt.Sprintf("Halo kak, ada %s, Dikirim oleh %s, dengan alamat %s. \n tujuan Penerima %s, \n silahkan hubungi security pak %s, Terimakasih !", activity.Document_type, activity.Sender, activity.Address, activity.Receiver, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			}

			// Sleep for a while before the next poll
			time.Sleep(10 * time.Second)
		}

	}
}
