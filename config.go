package hosmas_email

// Config holds the essential SMTP configuration
type Config struct {
	Server   string
	Port     int
	Username string
	Password string
	Sender   string
}

// NewConfig creates a new SMTP configuration
func NewConfig(server string, port int, username, password string) *Config {
	return &Config{
		Server:   server,
		Port:     port,
		Username: username,
		Password: password,
		Sender:   "noreply@ccstiet.com", // Default sender
	}
}
