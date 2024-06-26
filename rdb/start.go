package rdb

import (
	"encoding/json"
	"flag"
	"fmt"
)

type ExecCommand struct {
	Action      string
	Credentials string
	Webhook     string
	Secret      string
}

type Connection struct {
	Username string
	Password string
	Host     string
	Port     int
}

func Run() (*ExecCommand, error) {
	var action string
	var credentials string
	var webhookURL string
	var secret string

	flag.StringVar(&action, "action", "", "API action to perform")
	flag.StringVar(&credentials, "credentials", "", "Connection credentials in JSON format")
	flag.StringVar(&webhookURL, "webhook", "", "Webhook URL to post results")
	flag.StringVar(&secret, "secret", "", "Secret key for webhook authentication")
	flag.Parse()

	// 解析连接信息
	var conn Connection
	if credentials != "" {
		if err := json.Unmarshal(
			[]byte(credentials),
			&conn,
		); err != nil {
			return nil, fmt.Errorf("error parsing credentials json: %s", err)
		}
	}

	return &ExecCommand{
		Action:      action,
		Credentials: credentials,
		Webhook:     webhookURL,
		Secret:      secret,
	}, nil

}
