package main

import (
	"fmt"

	"github.com/apipgdt/poc-apns/internal/config"
	apns "github.com/apipgdt/poc-apns/internal/pkg"
)

func main() {
	cfg := config.Get()

	fmt.Println(cfg.Apns.KeyID)

	client := apns.NewClient(cfg.Apns)

	payload := apns.Payload{
		Aps: apns.ApsPayload{
			MutableContent: true,
			Alert: struct {
				Title    string "json:\"title\""
				Subtitle string "json:\"subtitle\""
				Body     string "json:\"body\""
			}{
				Title:    "Title",
				Subtitle: "Subtitle",
				Body:     "Body",
			},
		},
		Gdt: apns.GdtPayload{
			Title:            "test",
			Message:          "test message",
			NotificationType: "ENUM_OF_TYPE",
			Source:           "INTERNAL",
			Deeplink:         "https://www.google.com",
			Media:            "https://s3.amazon.com/test/image.jpg",
			Data: struct {
				TrackerProperties struct {
					NotificationID string "json:\"notification_id\""
					AnotherKey     string "json:\"another_key\""
				} "json:\"tracker_properties\""
			}{
				TrackerProperties: struct {
					NotificationID string "json:\"notification_id\""
					AnotherKey     string "json:\"another_key\""
				}{
					NotificationID: "12345",
					AnotherKey:     "another_value",
				},
			},
		}}

	err := client.Send(payload)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Notification sent")
}
