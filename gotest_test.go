package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCurrencyConversionHandler_Success(t *testing.T) {
	fmt.Println("TestCurrencyConversionHandler_Success")
	c := gin.Default()
	c.GET("/", CurrencyConversionHandler)
	req, err := http.NewRequest("GET", "/?source=TWD&target=JPY&amount=$1,234,567", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}
	var responseMap map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseMap); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if msg, ok := responseMap["msg"].(string); !ok || msg != "success" {
		t.Errorf(`Expected "msg" field to be "success", but got %v`, msg)
	}
	fmt.Println("")
	fmt.Println("***************")
	fmt.Println("")
}

func TestCurrencyConversionHandler_SourceFormatError(t *testing.T) {
	fmt.Println("TestCurrencyConversionHandler_SourceFormatError")
	c := gin.Default()
	c.GET("/", CurrencyConversionHandler)
	req, err := http.NewRequest("GET", "/?source=CNY&target=JPY&amount=$1,234,567", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
	var responseMap map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseMap); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if msg, ok := responseMap["msg"].(string); !ok || !strings.Contains(msg, "Source Format Error") {
		t.Errorf(`Expected "msg" field to contain "Source Format Error", but got %v`, msg)
	}
	fmt.Println("")
	fmt.Println("***************")
	fmt.Println("")
}

func TestCurrencyConversionHandler_AmountFormatError(t *testing.T) {
	fmt.Println("TestCurrencyConversionHandler_AmountFormatError")
	c := gin.Default()
	c.GET("/", CurrencyConversionHandler)
	req, err := http.NewRequest("GET", "/?source=TWD&target=JPY&amount=aaa", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	w := httptest.NewRecorder()
	c.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, w.Code)
	}
	var responseMap map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &responseMap); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if msg, ok := responseMap["msg"].(string); !ok || !strings.Contains(msg, "Amount Format Error") {
		t.Errorf(`Expected "msg" field to contain "Amount Format Error", but got %v`, msg)
	}
	fmt.Println("")
	fmt.Println("***************")
	fmt.Println("")
}
