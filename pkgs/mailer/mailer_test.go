package mailer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Mailer(t *testing.T) {
	t.Parallel()
	t.Skip("TODO: implement mailer tests, issue with loading dynamic config")

	mb := NewMessageBuilder().
		SetBody("Hello World!").
		SetSubject("Hello").
		SetTo("John Doe", "john@doe.com").
		SetFrom("Jane Doe", "jane@doe.com")

	msg := mb.Build()

	mailer := Mailer{
		Host:     "",
		Port:     465,
		Username: "",
		Password: "",
		From:     "",
	}

	err := mailer.Send(msg)

	assert.Nil(t, err)
}
