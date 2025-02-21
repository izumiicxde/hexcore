package mail

import (
	"fmt"
	"hexcore/config"

	"github.com/resend/resend-go/v2"
)

func SendMail(to, username, code string) error {
	client := resend.NewClient(config.Envs.RESEND_API_KEY)

	params := &resend.SendEmailRequest{
		From:    "zen-axcd <zen.axcd@resend.dev>",
		To:      []string{to},
		Html:    GenerateVerificationEmail(username, code),
		Subject: "Verification email from Zen-Axcd",
	}

	_, err := client.Emails.Send(params)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
