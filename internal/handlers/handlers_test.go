package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, ts *httptest.Server, method string, path string) (*http.Response, Response) {
	defer ts.Close()
	req, err := http.NewRequest(method, ts.URL+path, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	respBody, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	defer resp.Body.Close()

	var response Response
	err = json.Unmarshal(respBody, &response)
	require.NoError(t, err)
	return resp, response
}

func TestHandlerAdd(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  float64
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "5.5",
			operand2: "7.8",
			wantRes:  13.3,
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "5",
			operand2: "7",
			wantRes:  12,
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-2",
			operand2: "",
			wantRes:  0,
		},
		{
			name:     "test #4",
			success:  true,
			operand1: "-0.0",
			operand2: "3.2",
			wantRes:  3.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(HandlerAdd())
			target := fmt.Sprintf("/api/add?a=%s&b=%s", tt.operand1, tt.operand2)
			resp, body := testRequest(t, ts, http.MethodGet, target)

			defer resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Value)
		})
	}
}

func TestHandlerSub(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  float64
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "4.6",
			operand2: "5.2",
			wantRes:  -0.6000000000000005,
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "0",
			operand2: "7",
			wantRes:  -7.0,
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-2",
			operand2: "",
			wantRes:  0,
		},
		{
			name:     "test #4",
			success:  true,
			operand1: "-0.0",
			operand2: "-3.2",
			wantRes:  3.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(HandlerSub())
			target := fmt.Sprintf("/api/sub?a=%s&b=%s", tt.operand1, tt.operand2)
			resp, body := testRequest(t, ts, http.MethodGet, target)

			defer resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Value)
		})
	}
}

func TestHandlerMul(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  float64
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "5",
			operand2: "7",
			wantRes:  35.0,
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "-9",
			operand2: "0",
			wantRes:  0.0,
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-1",
			operand2: "",
			wantRes:  0,
		},
		{
			name:     "test #4",
			success:  true,
			operand1: "-0.0",
			operand2: "3.2",
			wantRes:  0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(HandlerMul())
			target := fmt.Sprintf("/api/mul?a=%s&b=%s", tt.operand1, tt.operand2)
			resp, body := testRequest(t, ts, http.MethodGet, target)

			defer resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Value)
		})
	}
}

func TestHandlerDiv(t *testing.T) {
	tests := []struct {
		name     string
		success  bool
		operand1 string
		operand2 string
		wantRes  float64
	}{
		{
			name:     "test #1",
			success:  true,
			operand1: "44",
			operand2: "2",
			wantRes:  22.0,
		},
		{
			name:     "test #2",
			success:  true,
			operand1: "3",
			operand2: "8",
			wantRes:  0.375,
		},
		{
			name:     "test #3",
			success:  false,
			operand1: "-2",
			operand2: "0",
			wantRes:  0,
		},
		{
			name:     "test #4",
			success:  false,
			operand1: "-0.0",
			operand2: "",
			wantRes:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(HandlerDiv())
			target := fmt.Sprintf("/api/div?a=%s&b=%s", tt.operand1, tt.operand2)
			resp, body := testRequest(t, ts, http.MethodGet, target)

			defer resp.Body.Close()
			assert.Equal(t, http.StatusOK, resp.StatusCode)
			assert.Equal(t, tt.success, body.Success)
			assert.Equal(t, tt.wantRes, body.Value)
		})
	}
}
