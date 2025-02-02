package handler

import (
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
		name     string
		wantErr  bool
		wantIp   string
		wantPort int
	}{
		{
			name:    "test getConnectConf",
			wantErr: false,
			wantIp:  "192.168.123.234:5555",
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
			assert.Equal(t, "{\"address\":\""+tt.wantIp+"\"}\n", rec.Body.String())
		})
	}

}

// func TestConnectDevice(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		wantErr bool
// 		inputAddress string
// 		inputAdbPath string
// 	}{
// 		{
// 			name:    "test connect device",
// 			wantErr: false,
// 			inputAddress: "192.168.123.234:5555",
// 			inputAdbPath: "/usr/bin/adb",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

// 	})
// }
