package apns

import (
	"encoding/json"
	"fmt"

	"github.com/apipgdt/poc-apns/internal/config"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

type Client struct {
	cfg  config.ApnsConfig
	Apns *apns2.Client
}

type Payload struct {
	Aps ApsPayload `json:"aps"`
	Gdt GdtPayload `json:"gdt"`
}

type ApsPayload struct {
	Alert struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Body     string `json:"body"`
	} `json:"alert"`
	MutableContent bool `json:"mutable-content"`
}

type GdtPayload struct {
	Title            string `json:"title"`
	Message          string `json:"message"`
	NotificationType string `json:"notification_ype"`
	Source           string `json:"source"`
	Deeplink         string `json:"deeplink"`
	Media            string `json:"media"`
	Data             struct {
		TrackerProperties struct {
			NotificationID string `json:"notification_id"`
			AnotherKey     string `json:"another_key"`
		} `json:"tracker_properties"`
	}
}

func NewClient(cfg config.ApnsConfig) *Client {
	authKey, err := token.AuthKeyFromFile(cfg.KeyFile)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	newToken := &token.Token{
		AuthKey: authKey,
		KeyID:   cfg.KeyID,
		TeamID:  cfg.TeamID,
	}

	client := Client{
		cfg:  cfg,
		Apns: apns2.NewTokenClient(newToken).Development(),
	}

	return &client
}

func (c *Client) Send(payload Payload) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	notification := &apns2.Notification{
		DeviceToken: c.cfg.DeviceToken,
		Topic:       c.cfg.Topic,
		Payload:     jsonPayload,
	}

	res, err := c.Apns.Push(notification)
	if err != nil {
		return err
	}

	fmt.Println("statusCode:", res.StatusCode)
	fmt.Println("reason:", res.Reason)
	fmt.Println("apnsID:", res.ApnsID)

	return nil
}
