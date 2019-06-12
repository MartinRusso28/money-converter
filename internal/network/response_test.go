package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fail_Response_Without_StatusCode(t *testing.T) {
	response := Response{
		Body: "body",
	}

	valid := response.Valid()

	assert.False(t, valid)
}