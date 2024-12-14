package http

import (
	"bytes"
	"call-center/internal/domain"
	"call-center/internal/repository"
	"call-center/internal/service"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"

	"net/http/httptest"
	"testing"
)

var (
	callRepo    = repository.NewCallMockRepo()
	callService = service.NewCallService(callRepo)
	services    = service.NewServices(callService)
	router      = NewHandler(services).Init()
)

func TestCreateCall(t *testing.T) {
	t.Run("Test Create Call", func(t *testing.T) {
		createdCall := domain.CallCreateRequest{
			ClientName:  "John Doe",
			PhoneNumber: "+1234567890",
			Description: "Need support",
		}
		w := httptest.NewRecorder()
		body, _ := json.Marshal(createdCall)
		req, _ := http.NewRequest("POST", "/calls", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		w1 := httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/calls/1", nil)
		router.ServeHTTP(w1, req)

		data, _ := io.ReadAll(w1.Result().Body)
		var response domain.Call
		json.Unmarshal([]byte(string(data)), &response)

		responseCall := domain.CallCreateRequest{
			ClientName:  response.ClientName,
			Description: response.Description,
			PhoneNumber: response.PhoneNumber,
		}

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, createdCall, responseCall)
	})
}

func TestGetAllCalls(t *testing.T) {
	t.Run("Test Get All Calls", func(t *testing.T) {

		createdCall := domain.CallCreateRequest{
			ClientName:  "John Doe",
			PhoneNumber: "+1234567890",
			Description: "Need support",
		}
		w := httptest.NewRecorder()
		expectedCall := []domain.CallCreateRequest{createdCall, createdCall}
		body, _ := json.Marshal(createdCall)
		req, _ := http.NewRequest("POST", "/calls", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		body, _ = json.Marshal(createdCall)
		req, _ = http.NewRequest("POST", "/calls", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/calls", nil)
		router.ServeHTTP(w, req)

		data, _ := io.ReadAll(w.Result().Body)
		var response []domain.Call
		json.Unmarshal([]byte(string(data)), &response)
		responseCall := []domain.CallCreateRequest{
			domain.CallCreateRequest{
				ClientName:  response[0].ClientName,
				Description: response[0].Description,
				PhoneNumber: response[0].PhoneNumber,
			}, domain.CallCreateRequest{
				ClientName:  response[1].ClientName,
				Description: response[1].Description,
				PhoneNumber: response[1].PhoneNumber,
			}}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedCall, responseCall)
	})
}

func TestUpdateCallStatus(t *testing.T) {
	t.Run("Test Update Call Status", func(t *testing.T) {
		type StatusRequest struct {
			Status string `json:"status" binding:"required,oneof=открыта закрыта"`
		}

		createdCall := domain.CallCreateRequest{
			ClientName:  "John Doe",
			PhoneNumber: "+1234567890",
			Description: "Need support",
		}
		updatedCall := StatusRequest{
			Status: "закрыта",
		}

		w := httptest.NewRecorder()
		body, _ := json.Marshal(createdCall)
		req, _ := http.NewRequest("POST", "/calls", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		body, _ = json.Marshal(createdCall)
		req, _ = http.NewRequest("POST", "/calls", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		body, _ = json.Marshal(updatedCall)
		req, _ = http.NewRequest("PATCH", "/calls/1/status", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		req, _ = http.NewRequest("GET", "/calls/1", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		data, _ := io.ReadAll(w.Result().Body)
		var response domain.Call
		json.Unmarshal([]byte(string(data)), &response)
		assert.Equal(t, updatedCall.Status, response.Status)
	})
}

func TestDeleteCall(t *testing.T) {
	t.Run("Test Delete Call", func(t *testing.T) {
		createdCall := domain.CallCreateRequest{
			ClientName:  "John Doe",
			PhoneNumber: "+1234567890",
			Description: "Need support",
		}
		w := httptest.NewRecorder()
		body, _ := json.Marshal(createdCall)
		req, _ := http.NewRequest("POST", "/calls", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)

		w = httptest.NewRecorder()
		body, _ = json.Marshal(createdCall)
		req, _ = http.NewRequest("DELETE", "/calls/1", bytes.NewReader(body))
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
