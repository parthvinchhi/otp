package otp

import (
	"fmt"
	"net/smtp"
	"time"

	"math/rand"
)

func VerifyOtp(otp, newOtp int) bool {
	if otp != newOtp {
		return false
	}

	return true
}

func GenerateOtp(n int) int {
	if n < 1 {
		return 0
	}

	min := 1
	for i := 1; i < n; i++ {
		min *= 10
	}
	max := min*10 - 1

	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

type EmailConfig struct {
	From       string
	Password   string
	SMTPServer string
	SMTPPort   string
}

func NewEmailConfig(from, password, smtpServer, smtpPort string) *EmailConfig {
	return &EmailConfig{
		From:       from,
		Password:   password,
		SMTPServer: smtpServer,
		SMTPPort:   smtpPort,
	}
}

func (cfg *EmailConfig) SendEmail(to []string, otp int) error {
	msg := fmt.Sprintf("Subject : OTP Verification\n\n Your OTP is : %d", otp)

	auth := smtp.PlainAuth("", cfg.From, cfg.Password, cfg.SMTPServer)

	if err := smtp.SendMail(
		fmt.Sprintf("%s:%s", cfg.SMTPServer, cfg.SMTPPort),
		auth,
		cfg.From,
		to,
		[]byte(msg),
	); err != nil {
		return err
	}

	return nil
}
