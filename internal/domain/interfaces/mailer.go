package interfaces

import "context"

type Mailer interface {
	SendTasksFetchedEmail(context.Context, string) error
}
