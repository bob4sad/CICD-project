package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestHandleMain(t *testing.T) {
	testData := []string{"this test is made by bob4sad", "i dont know how to use golang", "help me"}

	testCases := []struct {
		name string
		data string
		want []byte
	}{
		{
			name: testData[0],
			data: testData[0],
			want: []byte(testData[0]),
		},
		{
			name: testData[1],
			// data: "hello i am error",
			data: testData[1],
			want: []byte(testData[1]),
		},
		{
			name: testData[2],
			data: testData[2],
			want: []byte(testData[2]),
		},
	}

	handler := http.HandlerFunc(handleMain)
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("/?data=%s", tc.data), nil)
			handler.ServeHTTP(rec, req)
			if (string(tc.want) != string(rec.Body.Bytes()) ) {
				t.FailNow()
			}
		})
	}
}
