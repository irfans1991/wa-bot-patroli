package handler

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

// var client *whatsmeow.Client

func SendMessages(ctx context.Context, client *whatsmeow.Client, message string) error {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	if client == nil {
		return fmt.Errorf("client is nil")
	}

	if ctx == nil {
		return fmt.Errorf("context is nil")
	}

	if message == "" {
		return fmt.Errorf("message is empty")
	}
	targetJID := types.JID{
		User:   os.Getenv("NUMBER_PHONE"),
		Server: "s.whatsapp.net",
	}
	_, err := client.SendMessage(ctx, targetJID, &proto.Message{
		Conversation: &message,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Message sent: %s\n", message)
	return nil
}
