package hosmas_email

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"
)

// EmailService handles email operations
type EmailService struct {
	config        *Config
	templateStore *TemplateStore
}

// NewEmailService creates a new email service instance
func NewEmailService(config *Config) (*EmailService, error) {
	templateStore, err := NewTemplateStore(config.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create template store: %v", err)
	}

	return &EmailService{
		config:        config,
		templateStore: templateStore,
	}, nil
}

// SendMail is the main function to send emails using templates
// messageType: type of message from database (e.g., "leave", "fee_reminder")
// to: recipient email addresses
// data: map of required fields (e.g., "student_name", "roll_number")
func (s *EmailService) SendMail(messageType string, to []string, data map[string]string) error {
	template, err := s.templateStore.GetTemplate(messageType)
	if err != nil {
		return fmt.Errorf("failed to get template: %v", err)
	}

	body := template.MessageTemplate
	for key, value := range data {
		placeholder := fmt.Sprintf("{%s}", key)
		body = strings.Replace(body, placeholder, value, -1)
	}

	// Send email using SMTP
	tlsConfig := &tls.Config{
		ServerName: s.config.Server,
	}

	conn, err := smtp.Dial(fmt.Sprintf("%s:%d", s.config.Server, s.config.Port))
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer conn.Close()

	if err = conn.StartTLS(tlsConfig); err != nil {
		return fmt.Errorf("failed to start TLS: %v", err)
	}

	auth := smtp.PlainAuth("", s.config.Username, s.config.Password, s.config.Server)
	if err = conn.Auth(auth); err != nil {
		return fmt.Errorf("failed to authenticate: %v", err)
	}

	if err = conn.Mail(s.config.Sender); err != nil {
		return fmt.Errorf("failed to set sender: %v", err)
	}

	for _, recipient := range to {
		if err = conn.Rcpt(recipient); err != nil {
			return fmt.Errorf("failed to set recipient %s: %v", recipient, err)
		}
	}

	wc, err := conn.Data()
	if err != nil {
		return fmt.Errorf("failed to create data writer: %v", err)
	}
	defer wc.Close()

	message := fmt.Sprintf("From: %s\r\n", s.config.Sender)
	message += fmt.Sprintf("To: %s\r\n", to[0])
	message += fmt.Sprintf("Subject: %s\r\n", template.Subject)
	message += "MIME-Version: 1.0\r\n"
	message += "Content-Type: text/plain; charset=\"utf-8\"\r\n"
	message += "\r\n" + body

	if _, err = fmt.Fprint(wc, message); err != nil {
		return fmt.Errorf("failed to write email body: %v", err)
	}

	return nil
}

// Close closes the database connection
func (s *EmailService) Close() error {
	return s.templateStore.Close()
}
