package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewHandlerV1(t *testing.T) {

	v1Base, hdlr := NewHandlerV1("/testing/v1", mockHandler.repos, mockHandler.jwt, mockHandler.log)

	assert.NotNil(t, hdlr)

	assert.Equal(t, hdlr.log, mockHandler.log)
	assert.Equal(t, hdlr.jwt, mockHandler.jwt)
	assert.Equal(t, hdlr.repos, mockHandler.repos)

	assert.Equal(t, "/testing/v1/v1/abc123", v1Base("/abc123"))
	assert.Equal(t, "/testing/v1/v1/abc123", v1Base("/abc123"))
}
