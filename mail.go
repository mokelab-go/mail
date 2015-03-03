package mail

type MailClient interface {
	// Sends email
	Send(email, subject, body string) error
}
