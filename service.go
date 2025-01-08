package hosmas_email

import (
	"fmt"
	"net/smtp"
)

// EmailService handles email operations
type EmailService struct {
	config *Config
}

// NewEmailService creates a new email service instance
func NewEmailService(config *Config) *EmailService {
	return &EmailService{
		config: config,
	}
}

// SendEmail sends an email using the configured SMTP settings
func (s *EmailService) SendEmail(to []string, subject, body string) error {
	addr := fmt.Sprintf("%s:%d", s.config.Server, s.config.Port)
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Server)

	// Construct email headers
	headers := make(map[string]string)
	headers["From"] = s.config.Sender
	headers["To"] = to[0]
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/plain; charset=\"utf-8\""

	// Build message
	message := ""
	for key, value := range headers {
		message += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	message += "\r\n" + body

	return smtp.SendMail(addr, auth, s.config.Sender, to, []byte(message))
}
