package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"whatsapp-bot/database"
	"whatsapp-bot/domain"
	"whatsapp-bot/handler"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"gorm.io/gorm"
	// "google.golang.org/protobuf/proto"
	// "google.golang.org/protobuf/proto"
)

var (
	// spreadsheetId string     = "19fcIZe-836Qczm7SO_vKD-j2AWEM8xhtVN6bP5D2lUE"
	// sheetKey      domain.Key = domain.Key{
	// 	Type:                       "service_account",
	// 	Project_id:                 "sturdy-rampart-414507",
	// 	PrivateKeyId:               "f93d36b069b800113fffd57438e2e8157b864f6e",
	// 	PrivateKey:                 "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDWs3maOaxWFrBX\nP80ySsEg1XI3mIzl2+FFtpLY6Jo9SwhoCXCqeEySm/7DZvKrQvMOf2lmlrRT/vSz\n/Alvjl7YmEdQRR35MfCwMqoDGpGzQzpa2htENu/CNoBFT0KuRI7RoR6wYD0r2Yiv\nYdx1PJFgxzl7Fw9jZGpqTM0zEg8wDFWc4PUFhrQ/pjrQeNpAKEmAHpc1CpgXf5T4\nxeDfIept0vDzvgOayLLaiebEHgqX2A76ZyMZMKQ0tojVn0ewxXFOIxur96uC7PrE\nwI4nmB+ZPUTWFjWcckFt7nPBkgpZnqZxWc12rzD8O+We2ssdJ5Lu6e0x573sXBM8\nzyZYn2oNAgMBAAECggEAFfxE7IURX7NtFVIkOg0rn6gbzROjAfq1I8VhXjLbcIFZ\neAXXcq1k8fJlcS7+lIkDiK2FZoryJlPKTQ/3C2Rh0/4r901Ml3MEuZZGkJDM8LVm\nqV0CQ8b4YEPieLZ4Fo5Qrai5EsiYGwBoncUApUSxgl6ERVWELuJef0okDSHBSmMY\n6gpJ2Bs00RCA/RttWMYCGhX7Ap7le8HeBhdo4haedK4g7AEzOE7r5wx3n29GVXYP\nWIdkPM5HGXTr0hoXqkuVzYE7sz/SR8bArCkSSaV3Elg5NhuWVo5bbKHGfDCR1pgG\nh4OLorxG5HjnUvRJi6wgbl/irMdHHZbbBkIJgNbw5QKBgQDyu2qwQ9qG16Le9Fak\neZsQveTSBArUqcakiP10J51OEDKFILgV5QmtSuVgDmvt3GtS2BvHWAQnSATmJbYE\nzCH5SA9yAj3jMULvUadqSqpUOylZWdsstzFBF4JfSexdujSstBOwkR+4bjqko+pZ\nALJLGGk8En9mz4wMA8/n4HjyawKBgQDib9D8JiXnvaOvkVnZJg53Wm0mU8ECOwz7\n4dLqd1/4Ne6i5wuakdhXxH/Xi6mTh6UpIjpkaAGTtUQ4o7UpjJZXv56B7yuXRzX8\niacG/ztbTQPgkBMKVvEMIA5ZlwBY2N4WmHW2uuqKVyiDoJXeB1ZqkfERrxGF9hhe\nTm7UxUPjZwKBgBGKeatDPAeCjcKO5bP9FbegAWkr4Zx8qF47iBVcx4Fcqkn0k1kB\ntDFRuEG4yRPnWdoiw98j/SjZi4EWOJEgL1RUYMlcYByK8stYm0CTJLvxiotqcuxI\nUMbjbAh2zx321tekK8gWtLQfQx2iusyzwC6tFO3CaHKFjRlRLTUUg6kBAoGAZIsZ\n1JxrwFaoU9DKgG0xhzWTuaz+yqrFvSIssQeIiYIJEDRfJqcT1QTfbyoIhgV5BKa6\nHtM/wAKA5FVsn6JZL2VaUH+Ob4GjKxtEGwtBo+yyiAfxLGomZiflt0ohUVIaaxYP\nTq/4bKz+xFONuCSx7mgXDq8ZrM31TAqvk7JzzokCgYAf38RFSi4K7zchae6jTF+R\npDcP+A6xuPLkdC7tLXdUMzpHGi9GPu8K+C4S4HySaDEWPuaucYq7GmnMIAna/Y53\ny2Qezpd4sZ3qmxeyZVtLiQeSZ7LmZNkR+RvExfUwg+tLqM0GIPctG+yK03uYsEpg\neay1pU4Iy0vGrfgMswGRzg==\n-----END PRIVATE KEY-----\n",
	// 	ClientEmail:                "sheet-service-account@sturdy-rampart-414507.iam.gserviceaccount.com",
	// 	ClientId:                   "113724890331456851600",
	// 	AuthUri:                    "https://accounts.google.com/o/oauth2/auth",
	// 	TokenUri:                   "https://oauth2.googleapis.com/token",
	// 	AuthProvider_x509_cert_url: "https://www.googleapis.com/oauth2/v1/certs",
	// 	Client_x509_cert_url:       "https://www.googleapis.com/robot/v1/metadata/x509/sheet-service-account%40sturdy-rampart-414507.iam.gserviceaccount.com",
	// }
	client *whatsmeow.Client
	ctx    context.Context
)

func main() {
	ctx := context.Background()
	var mutasi_masuks []domain.Mutasi_masuks
	// koneksi database mysql
	db := database.MariadbConnect(ctx)

	if db == nil {
		panic("error connect database")
	}

	db.AutoMigrate(&mutasi_masuks)

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New("sqlite3", "file:myfirstbot.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	// If you want multiple sessions, remember their JIDs and use .GetDevice(jid) or .GetAllDevices() instead.
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	client = whatsmeow.NewClient(deviceStore, clientLog)
	// client.AddEventHandler(eventHandler)

	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				// Render the QR code here
				// e.g. qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
				// or just manually `echo 2@... | qrencode -t ansiutf8` in a terminal
				fmt.Println("QR code:", evt.Code)
				qrterminal.Generate(evt.Code, qrterminal.M, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		// Already logged in, just connect
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	PollNewActivities(db)

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
	// select {}
}

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
			case Unit == "Boks" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada barang %s Boks masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Colly" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada barang %s Colly masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Galon" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada Air %s Galon masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Botol" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada %s Botol masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Roll" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada Barang %s Roll masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Tabung" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada %s Tabung masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Pcs" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada Barang %s Pcs masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Pail" && dates == timeNow && activity.Type_Mutasi == "in":
				message = fmt.Sprintf("Halo kak, ada %s pail masuk di area perusahaan dari %s, diterima oleh Pak %s, silahkan hubungi security, Terimakasih !", activity.TotalItems, activity.Supplier_Name, activity.Security)
				if err := handler.SendMessages(ctx, client, message); err != nil {
					log.Fatalf("Failed to send Unit: %v", err)
				}
			case Unit == "Liter" && dates == timeNow && activity.Type_Mutasi == "in":
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
