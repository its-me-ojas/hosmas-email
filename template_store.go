package hosmas_email

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// EmailTemplate represents an email template from the database
type EmailTemplate struct {
	Type            string // Type of message (e.g., "leave", "fee_reminder")
	Subject         string // Email subject line
	MessageTemplate string // Email body template with placeholders
}

// TemplateStore handles database operations for email templates
type TemplateStore struct {
	db *sql.DB
}

// NewTemplateStore creates a new template store instance
func NewTemplateStore(dbURL string) (*TemplateStore, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &TemplateStore{
		db: db,
	}, nil
}

// GetTemplate fetches an email template by its type
func (s *TemplateStore) GetTemplate(messageType string) (*EmailTemplate, error) {
	template := &EmailTemplate{}
	err := s.db.QueryRow(
		"SELECT type, subject, message_template FROM email_subjects WHERE type = $1",
		messageType,
	).Scan(&template.Type, &template.Subject, &template.MessageTemplate)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("no template found for type: %s", messageType)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch template: %v", err)
	}

	return template, nil
}

// Close closes the database connection
func (s *TemplateStore) Close() error {
	return s.db.Close()
}
