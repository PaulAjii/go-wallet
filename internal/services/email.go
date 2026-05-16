package services

import (
	"fmt"

	"github.com/PaulAjii/go-wallet/pkg/config"
	"github.com/resend/resend-go/v3"
)

func SendVerificationEmail(to, verificationLink string) error {
	client := resend.NewClient(config.ApplicationConfig.Resend.APIKey)

	params := &resend.SendEmailRequest{
		From:    config.ApplicationConfig.Resend.FromEmail,
		To:      []string{to},
		Subject: "Verify your email",
		Html:    fmt.Sprintf(`<p>Verify your email by clicking on this <a href="%s">link</a>.</p>`, verificationLink),
	}

	_, err := client.Emails.Send(params)
	return err
}
