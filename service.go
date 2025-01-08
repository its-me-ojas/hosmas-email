package hosmas_email

import (
	"crypto/tls"
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
	// Create TLS config
	tlsConfig := &tls.Config{
		ServerName: s.config.Server,
	}

	// Connect to the SMTP Server
	conn, err := smtp.Dial(fmt.Sprintf("%s:%d", s.config.Server, s.config.Port))
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer conn.Close()

	// Start TLS
	if err = conn.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %v", err)
	}

	// Auth
	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Server)
	if err = conn.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	// Set the sender and recipient
	if err = conn.Mail(s.config.Sender); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	for _, recipient := range to {
		if err = conn.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient %s: %v", recipient, err)
		}
	}

	// Send the email body
	wc, err := conn.Data()
	if err != nil {
		return fmt.Errorf("failed to create data writer: %v", err)
	}
	defer wc.Close()

	// Construct email headers
	message := fmt.Sprintf("From: %s\r\n", s.config.Sender)
	message += fmt.Sprintf("To: %s\r\n", to[0])
	message += fmt.Sprintf("Subject: %s\r\n", subject)
	message += "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
	message += "\r\n" + body

	if _, err = fmt.Fprint(wc, message); err != nil {
		return fmt.Errorf("failed to write email body: %v", err)
	}

	return nil
}
