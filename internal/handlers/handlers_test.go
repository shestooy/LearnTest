package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mainHandle(t *testing.T) {
	tests := []struct {
		name    string
		params  []string
		expCode int
		expBody string
	}{
		{name: "TestValidRequest", params: []string{"3", "moscow"}, expCode: http.StatusOK, expBody: "Мир кофе,Сладкоежка,Кофе и завтраки"},
		{name: "TestUnsupportedCity", params: []string{"4", "testCity"}, expCode: http.StatusBadRequest, expBody: "wrong city value"},
		{name: "TestReturnAllCafe", params: []string{"5", "moscow"}, expCode: http.StatusOK, expBody: "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/cafe?count=%s&city=%s", tt.params[0], tt.params[1]), nil)
			respRec := httptest.NewRecorder()
			h := http.HandlerFunc(MainHandle)
			h.ServeHTTP(respRec, req)

			assert.Equal(t, tt.expCode, respRec.Code, "unexpected response code")
			body := respRec.Body.String()

			assert.Equal(t, tt.expBody, body, "unexpected bytes response body")
		})
	}
}
