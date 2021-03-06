package models

import "errors"

var ErrInvalidEmailCode = errors.New("Invalid or expired email code")

type SendEmailCommand struct {
	To           []string
	Template     string
	Data         map[string]interface{}
	Info         string
	EmbededFiles []string
}

type SendEmailCommandSync struct {
	SendEmailCommand
}

type SendWebhook struct {
	Url        string
	User       string
	Password   string
	Body       string
	HttpMethod string
}

type SendWebhookSync struct {
	Url        string
	User       string
	Password   string
	Body       string
	HttpMethod string
}

type SendResetPasswordEmailCommand struct {
	User *User
}

type ValidateResetPasswordCodeQuery struct {
	Code   string
	Result *User
}
