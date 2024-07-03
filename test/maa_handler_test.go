package handler

import (
	"maa-web/internal/handler"
	"net/http"
	"net/http/httptest"
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
			wantVersion: "v5.4.0-beta.2",
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
