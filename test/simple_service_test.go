package test

import (
	"programmerzamannow/belajar-golang-restful-api/simple"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceWithError(t *testing.T) {
	simpleService, err := simple.InitializeSimpleService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializeSimpleService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}
