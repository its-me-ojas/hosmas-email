# hosmas-email

A lightweight Go package for sending templated emails using SMTP with database-driven templates.

## Installation

```bash
go get github.com/its-me-ojas/hosmas-email
```

## Database Setup

The package expects a PostgreSQL database with the following table:

```sql
CREATE TABLE email_subjects (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) UNIQUE NOT NULL,
    subject TEXT NOT NULL,
    message_template TEXT NOT NULL
);
```

## Usage

```go
package main

import (
    "log"
    "github.com/its-me-ojas/hosmas-email"
)

func main() {
    // Create configuration
    config := hosmas_email.NewConfig(
        "smtp.example.com",                     // SMTP server
        587,                                    // SMTP port
        "your-username",                        // SMTP username
        "your-password",                        // SMTP password
        "postgres://user:pass@localhost:5432/hosmas-email", // Database URL
    )

    // Initialize email service
    emailService, err := hosmas_email.NewEmailService(config)
    if err != nil {
        log.Fatal(err)
    }
    defer emailService.Close()

    // Send an email using a template
    err = emailService.SendMail(
        "leave",                              // message type
        []string{"recipient@example.com"},     // recipients
        map[string]string{                     // template data
            "student_name": "John Doe",
            "roll_number": "12345",
            "start_date": "2024-01-10",
            "end_date": "2024-01-15",
            "reason": "Family function",
        },
    )
    if err != nil {
        log.Printf("Error: %v", err)
    }
}
```

## Available Message Types

The package supports various message types, each requiring specific template data:

1. `leave`:

   - student_name
   - roll_number
   - start_date
   - end_date
   - reason

2. `fee_reminder`:

   - student_name
   - amount
   - semester
   - due_date

3. `maintenance_request`:

   - student_name
   - room_number
   - issue_type
   - ticket_number

4. `event_invitation`:

   - event_name
   - event_date
   - event_time
   - venue
   - event_description

5. `room_change`:
   - student_name
   - new_room
   - new_block
   - new_floor
   - deadline

## License

MIT License
