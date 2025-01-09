package hosmas_email

// Config holds the SMTP and database configuration
type Config struct {
	// SMTP Configuration
	Server   string
	Port     int
	Username string
	Password string
	Sender   string

	// Database Configuration
	DatabaseURL string
}

// NewConfig creates a new configuration with SMTP and database settings
func NewConfig(smtpServer string, smtpPort int, smtpUsername, smtpPassword, dbURL string) *Config {
	return &Config{
		Server:      smtpServer,
		Port:        smtpPort,
		Username:    smtpUsername,
		Password:    smtpPassword,
		Sender:      "noreply@ccstiet.com", // Default sender
		DatabaseURL: dbURL,
	}
}
