package main

import (
	"net/http"
	"os"
	"testing"

	"github.com/azbshiri/common/test"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var testApp *App

func TestMain(m *testing.M) {
	testApp = newApp(
		connect(),
		mux.NewRouter(),
	)
	os.Exit(m.Run())
}

func TestGetBudgets(t *testing.T) {
	res, err := test.DoRequest(testApp, "GET", "/address/40115", nil)
	assert.NoError(t, err)
	assert.Equal(t, res.Code, http.StatusOK)
}
