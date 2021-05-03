package handler_test

import (
	"testing"

	"github.com/arthurDonizetti/go-hexagonal/adapters/web/handler"
	"github.com/stretchr/testify/require"
)

func TestHandler_JsonError(t *testing.T) {
	msg := "invalid_message"
	result := handler.JsonError(msg)
	require.Equal(t, []byte(`{"message":"invalid_message"}`), result)
}
