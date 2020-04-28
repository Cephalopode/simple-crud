package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"simple-crud/conf"
	"simple-crud/models"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, data gin.H) *httptest.ResponseRecorder {
	reqBody, _ := json.Marshal(data)
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func readBodytoUserList(t *testing.T, w *httptest.ResponseRecorder) []models.Users {
	retData, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	var list []models.Users
	err = json.Unmarshal(retData, &list)
	if err != nil {
		t.Fatal(err)
	}
	return list
}
func TestServer(t *testing.T) {
	conf.Init()
	conf.C.Database.DB = conf.C.Database.DB + "_test"
	conf.C.Database.Debug = false
	conf.C.Reset = true
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	t.Run("ping", func(t *testing.T) {
		w := performRequest(router, "GET", "/ping", gin.H{})
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("list users", func(t *testing.T) {
		// Create user and check
		w := performRequest(router, "PUT", "/api/users", gin.H{"name": "test1", "age": 20})
		assert.Equal(t, http.StatusCreated, w.Code)

		w = performRequest(router, "GET", "/api/users", nil)
		assert.Equal(t, http.StatusOK, w.Code)
		users := readBodytoUserList(t, w)
		assert.Equal(t, "test1", users[0].Name)

		// Delete user and check
		w = performRequest(router, "DELETE", "/api/users", gin.H{"id": 1})
		assert.Equal(t, http.StatusOK, w.Code)

		w = performRequest(router, "GET", "/api/users", nil)
		assert.Equal(t, http.StatusOK, w.Code)
		users = readBodytoUserList(t, w)
		assert.Equal(t, 0, len(users))

	})
}
