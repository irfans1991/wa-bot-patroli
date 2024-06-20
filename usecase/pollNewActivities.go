package usecase

import (
	"context"
	"fmt"
	"log"
	"time"
	"whatsapp-bot/domain"
	"whatsapp-bot/handler"

	"gorm.io/gorm"
)

func PollNewActivities(db *gorm.DB) {
	ctx = context.Background()
	var lastCheckedID uint = 0
	var message string
	var Unit string

	for {
		var mutasi_masuks []domain.Mutasi_masuks

		db.Where("id > ?", lastCheckedID).Order("id asc").Find(&mutasi_masuks)
		timeNow := time.Now().Format("2006-01-02")

		for _, activity := range mutasi_masuks {

			fmt.Printf("New activity detected: %v\n", activity)
			lastCheckedID = uint(activity.Id)
			dates := activity.Created_at.Format("2006-01-02")
			Unit = activity.Unit

			switch {
			case Unit == "Boks" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada barang %s Boks masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Colly" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada barang %s Colly masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Galon" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada Air %s Galon masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Botol" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada %s Botol masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Roll" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada Barang %s Roll masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Tabung" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada %s Tabung masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Pcs" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada Barang %s Pcs masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Pail" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada %s pail masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Liter" && dates == timeNow && activity.Type_Mutasi == "Masuk Barang":
				message = fmt.Sprintf("Halo kak, ada %s Liter masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			}

			// if message == "Jhins" && dates == timeNow {
			// 	fmt.Println(message)
			// 	if err := handler.SendMessages(ctx, client, message); err != nil {

			// 		log.Fatalf("Failed to send message: %v", err)
			// 	}

			// }
		}

		// Sleep for a while before the next poll
		time.Sleep(10 * time.Second)
	}

}
