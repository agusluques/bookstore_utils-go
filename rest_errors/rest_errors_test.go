package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("msg", errors.New("error msg"))

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "msg")
	assert.EqualValues(t, err.Status, http.StatusInternalServerError)
	assert.EqualValues(t, err.Error, "internal_server_error")
	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, len(err.Causes), 1)
	assert.EqualValues(t, err.Causes[0], "error msg")
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("msg")

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "msg")
	assert.EqualValues(t, err.Status, http.StatusBadRequest)
	assert.EqualValues(t, err.Error, "bad_request")
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("msg")

	assert.NotNil(t, err)
	assert.EqualValues(t, err.Message, "msg")
	assert.EqualValues(t, err.Status, http.StatusNotFound)
	assert.EqualValues(t, err.Error, "not_found")
}
