package main

import (
	"fizzbuzz/internal/conf"
	"fizzbuzz/internal/server"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_fizzbuzz_handler(t *testing.T) {
	tests := []struct {
		name         string
		target       string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "nominal",
			target:       "/custom-fizzbuzz/3/5/2/fizz/buzz",
			expectedCode: http.StatusOK,
			expectedBody: "1,2,",
		},
		{
			name:         "fizz",
			target:       "/custom-fizzbuzz/4/9/4/fizz/buzz",
			expectedCode: http.StatusOK,
			expectedBody: "1,2,3,fizz,",
		},
		{
			name:         "no fizz",
			target:       "/custom-fizzbuzz/4/9/9//buzz",
			expectedCode: http.StatusOK,
			expectedBody: "1,2,3,,5,6,7,,buzz,",
		},
		{
			name:         "bad first int arg",
			target:       "/custom-fizzbuzz/bad/1/1/fizz/buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: "the first 3 parameters should be integers.",
		},
		{
			name:         "bad third int arg",
			target:       "/custom-fizzbuzz/1/1/wrong/fizz/buzz",
			expectedCode: http.StatusBadRequest,
			expectedBody: "the first 3 parameters should be integers.",
		},
		{
			name:         "no buzz",
			target:       "/custom-fizzbuzz/1/1/1/fizz/",
			expectedCode: http.StatusMovedPermanently,
			expectedBody: "<a href=\"/custom-fizzbuzz/1/1/1/fizz\">Moved Permanently</a>.\n\n",
		},
		{
			name:         "buzzfizz",
			target:       "/custom-fizzbuzz/2/4/8/buzz/fizz",
			expectedCode: http.StatusOK,
			expectedBody: "1,buzz,3,buzzfizz,5,buzz,7,buzzfizz,",
		},
	}

	gin.SetMode(gin.TestMode)
	router := server.Start(conf.NewConfig())

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRecorder()

			request := httptest.NewRequest(http.MethodGet, tt.target, nil)
			router.ServeHTTP(r, request)

			assert.Equal(t, tt.expectedCode, r.Code)
			assert.Equal(t, tt.expectedBody, r.Body.String())
		})
	}
}
