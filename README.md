# hosmas-email

A lightweight Go package for sending emails via SMTP with TLS support.

## Installation

```bash
go get github.com/its-me-ojas/hosmas-email
```

## Usage with AWS SES

```go
package main

import (
    "github.com/its-me-ojas/hosmas-email"
    "log"
)

func main() {
    config := hosmas_email.NewConfig(
        "email-smtp.ap-south-1.amazonaws.com", // AWS SES SMTP endpoint
        587,                                   // Port
        "YOUR_AWS_SES_SMTP_USERNAME",          // SMTP username
        "YOUR_AWS_SES_SMTP_PASSWORD",          // SMTP password
    )

    emailService := hosmas_email.NewEmailService(config)

    err := emailService.SendEmail(
        []string{"recipient@example.com"},
        "Test Subject",
        "This is a test email sent using AWS SES SMTP interface.",
    )
    if err != nil {
        log.Printf("Error sending email: %v", err)
        return
    }
    log.Println("Email sent successfully!")
}
```

## Features

- TLS encryption support
- AWS SES compatible
- Multiple recipients support
- Proper email headers (From, To, Subject, MIME-Version, Content-Type)

## Configuration

- `Server`: SMTP server address (e.g., AWS SES endpoint)
- `Port`: SMTP port (usually 587 for TLS)
- `Username`: SMTP username (AWS SES SMTP credentials)
- `Password`: SMTP password (AWS SES SMTP credentials)
- `Sender`: Email address of the sender

## License

MIT License
