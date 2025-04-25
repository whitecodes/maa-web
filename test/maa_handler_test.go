package handler

import (
	"bytes"
	"encoding/json"
	"maa-web/internal/handler"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetMaaVersion(t *testing.T) {
	tests := []struct {
		name        string
		wantErr     bool
		wantVersion string
	}{
		{
			name:        "test getversion",
			wantErr:     false,
			wantVersion: "v5.13.0-beta.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/version", nil)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			if err := handler.GetVersion(c); (err != nil) != tt.wantErr {
				t.Errorf("GetMaaVersion() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, http.StatusOK, c.Response().Status)
			assert.Equal(t, "application/json", c.Response().Header().Get(echo.HeaderContentType))

			assert.Equal(t, "{\"version\":\""+tt.wantVersion+"\"}\n", rec.Body.String())
		})
	}
}

func TestGetMaaConnected(t *testing.T) {
	tests := []struct {
		name          string
		wantErr       bool
		wantConnected bool
	}{
		{
			name:          "test connected",
			wantErr:       false,
			wantConnected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/maa/connected", nil)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			if err := handler.GetMaaConnected(c); (err != nil) != tt.wantErr {
				t.Errorf("GetMaaConnected() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, http.StatusOK, c.Response().Status)
			assert.Equal(t, "application/json", c.Response().Header().Get(echo.HeaderContentType))

			assert.Equal(t, "{\"connected\":"+strconv.FormatBool(tt.wantConnected)+"}\n", rec.Body.String())
		})
	}
}

func TestGetMaaConnectConf(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		wantIp  string
		wantAdb string
	}{
		{
			name:    "test getConnectConf",
			wantErr: false,
			wantIp:  "192.168.123.234:5555",
			wantAdb: "adb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/maa/connectConf", nil)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			if err := handler.GetMaaConnectConf(c); (err != nil) != tt.wantErr {
				t.Errorf("GetMaaConnectConf() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, http.StatusOK, c.Response().Status)
			assert.Equal(t, "application/json", c.Response().Header().Get(echo.HeaderContentType))
			assert.Equal(t, "{\"adbPath\":\""+tt.wantAdb+"\",\"address\":\""+tt.wantIp+"\"}\n", rec.Body.String())
		})
	}
}

func TestSetMaaConnectConf(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
		wantIp  string
		wantAdb string
	}{
		{
			name:    "test setConnectConf",
			wantErr: false,
			wantIp:  "192.168.123.204:5555",
			wantAdb: "/usr/bin/adb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // 创建请求参数

			requestData := map[string]interface{}{
				"address": tt.wantIp,
				"adbPath": tt.wantAdb,
			}

			// 将请求参数转为 JSON 格式
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				t.Fatalf("Failed to marshal JSON: %v", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/maa/connectConf", bytes.NewReader(jsonData))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			if err := handler.SetMaaConnectConf(c); (err != nil) != tt.wantErr {
				t.Errorf("SetMaaConnectConf() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equal(t, http.StatusOK, c.Response().Status)
			assert.Equal(t, "application/json", c.Response().Header().Get(echo.HeaderContentType))
			assert.Equal(t, "{\"adbPath\":\""+tt.wantAdb+"\",\"address\":\""+tt.wantIp+"\"}\n", rec.Body.String())
		})
	}

}

// func TestConnectDevice(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		wantErr      bool
// 		inputAddress string
// 		inputAdbPath string
// 		inputConfig  string
// 	}{
// 		{
// 			name:         "test connect device",
// 			wantErr:      false,
// 			inputAddress: "192.168.123.234:5555",
// 			inputAdbPath: "/usr/bin/adb",
// 			inputConfig:  "CompatPOSIXShell",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 			requestData := map[string]interface{}{
// 				"address": tt.inputAddress,
// 				"adbPath": tt.inputAdbPath,
// 				"config":  tt.inputConfig,
// 			}

// 			// 将请求参数转为 JSON 格式
// 			jsonData, err := json.Marshal(requestData)
// 			if err != nil {
// 				t.Fatalf("Failed to marshal JSON: %v", err)
// 			}

// 			req := httptest.NewRequest(http.MethodPost, "/maa/testConnect", bytes.NewReader(jsonData))
// 			req.Header.Set("Content-Type", "application/json")

// 			rec := httptest.NewRecorder()
// 			c := echo.New().NewContext(req, rec)

// 			if err := handler.TestConnect(c); (err != nil) != tt.wantErr {
// 				t.Errorf("TestConnect() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 			assert.Equal(t, http.StatusOK, c.Response().Status)
// 			assert.Equal(t, "application/json", c.Response().Header().Get(echo.HeaderContentType))
// 			assert.Equal(t, "{\"status\":\"success\"}\n", rec.Body.String())
// 		})
// 	}
}
