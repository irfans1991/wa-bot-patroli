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

// var (
// 	client *whatsmeow.Client
// )

func PollNewActivities(db *gorm.DB, client *whatsmeow.Client) {
	ctx = context.Background()

	var lastCheckedID uint = 0
	var message string
	// var Unit string
	for {
		var mutasi_masuks []domain.Mutasi_masuks

		db.Where("id > ?", lastCheckedID).Order("id asc").Find(&mutasi_masuks)
		timeNow := time.Now().Format("2006-01-02")

		for _, activity := range mutasi_masuks {
			activity.Supplier
			fmt.Printf("New activity detected: %v\n", activity)
			lastCheckedID = uint(activity.Id)
			dates := activity.Created_at.Format("2006-01-02")
			switch {
			case activity.Type_Mutasi == "Masuk" && activity.Supplier == "IKAN" && dates == timeNow:
				message = fmt.Sprintf("Halo kak, ada Ikan masuk di area perusahaan dari supplier %s, dengan total items %s %s, diterima oleh Pak %s,\n *_dengan informasi berikut : %s_*, \n silahkan hubungi security, Terimakasih !", activity.Supplier_Name, activity.TotalItems, activity.Unit, activity.Security, activity.Remark)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case activity.Type_Mutasi == "Masuk Barang" && dates == timeNow:
				message = fmt.Sprintf("Halo kak, ada barang %s %s masuk di area perusahaan dari %s, %s, diterima oleh Pak %s,\n *_dengan informasi berikut : %s_*, \n silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Unit, activity.Supplier_Name, activity.From, activity.Security, activity.Remark)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case activity.Type_Mutasi == "Keluar" && dates == timeNow:
				message = fmt.Sprintf("Halo kak, ada aktifitas barang Keluar, dengan informasi : *_%s_* , dari supplier %s ke %s, \n silahkan hubungi security untuk informasi lebih lanjut, Terimakasih !", activity.Remark, activity.Supplier_Name, activity.From)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			}
		}

		// Sleep for a while before the next poll
		time.Sleep(10 * time.Second)

	}

}
